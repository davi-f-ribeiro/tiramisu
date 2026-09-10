package subprovider

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// ---------- Engine: orchestrator ----------

// SubtitleEngine is the main orchestrator. It runs two provider pipelines
// (SubDB via video hash, then OpenSubtitles via IMDB ID) in sequence,
// each in its own goroutine, and sends the best result through the channel.
type SubtitleEngine struct {
	cfg      EngineConfig
	subdb    *SubDBProvider
	os       *OSProvider
	writer   *Writer
	rateMu   sync.Mutex
	rateTok  int
	rateRef  time.Time
}

// NewSubtitleEngine creates a new engine. Validates mandatory config.
func NewSubtitleEngine(cfg EngineConfig) (*SubtitleEngine, error) {
	// Apply defaults
	if cfg.PreferredLanguages == nil {
		cfg.PreferredLanguages = getDefaultConfig().PreferredLanguages
	}
	if cfg.MaxResultsPerProvider == 0 {
		cfg.MaxResultsPerProvider = getDefaultConfig().MaxResultsPerProvider
	}
	if cfg.SubDBTimeout == 0 {
		cfg.SubDBTimeout = getDefaultConfig().SubDBTimeout
	}
	if cfg.OSDownloadTimeout == 0 {
		cfg.OSDownloadTimeout = getDefaultConfig().OSDownloadTimeout
	}
	if cfg.OpenSubtitlesBaseURL == "" {
		cfg.OpenSubtitlesBaseURL = getDefaultConfig().OpenSubtitlesBaseURL
	}

	// Validate mandatory config
	if cfg.OpenSubtitlesKey == "" {
		return nil, fmt.Errorf("subprovider: OpenSubtitles API key is required (set subtitle.api_key in config.json)")
	}
	if cfg.FUSEMountPath == "" {
		return nil, fmt.Errorf("subprovider: FUSE mount path is required (set TIRAMISU_FUSE_MOUNT_PATH env var)")
	}

	return &SubtitleEngine{
		cfg:      cfg,
		subdb:    NewSubDBProvider(),
		os:       NewOSProvider(cfg.OpenSubtitlesKey, cfg.OpenSubtitlesBaseURL, cfg.OpenSubtitlesUser, cfg.OpenSubtitlesPass),
		writer:   NewWriter(cfg.FUSEMountPath),
		rateTok:  15,
		rateRef:  time.Now(),
	}, nil
}

// SubtitleJob represents a request to download a subtitle.
type SubtitleJob struct {
	// Input (provided by caller)
	VideoPath     string      // full path to .mkv in FUSE mount
	TorrentName   string      // torrent/filename (e.g. "Interstellar.2014.1080p.BluRay.x264.YTS")
	VideoHash     string      // SubDB MD5 hash of first+last 64KB (may be empty)
	ImdbID        string      // IMDB ID (e.g. "tt0816692")
	PreferredLang LanguageTag // user's preferred language

	// Output (sent through resultCh)
	ResultCh chan<- Result
}

// Run launches the subtitle download pipeline for a single job.
// Returns immediately (fire-and-forget goroutine).
// The result is sent through job.ResultCh when done (or after timeout).
func (e *SubtitleEngine) Run(job SubtitleJob) {
	ch := make(chan Result, 1)
	job.ResultCh = ch

	go func() {
		torrentName := job.TorrentName
		if torrentName == "" {
			torrentName = CleanFilename(job.VideoPath)
		}

		// Step 1: Check if .srt already exists on disk
		srtPath := e.writer.BuildSRTPath(job.VideoPath, job.PreferredLang)
		if e.writer.FileExists(srtPath) {
			logf("subtitle cache hit: %s", srtPath)
			ch <- Result{
				Content:  nil, // nil = using existing file on disk
				Filename: srtPath,
				Provider: "cache",
				Language: string(job.PreferredLang),
				Error:    nil,
			}
			return
		}

		// Phase 1: SubDB (hash-based, fast)
		if job.VideoHash != "" {
			if content, filename, err := e.trySubDBDownload(job.VideoHash, job.PreferredLang, torrentName, job.VideoPath); err == nil {
				logf("phase 1: subdb success → %s", filename)
				ch <- Result{
					Content:  content,
					Filename: filename,
					Provider: "subdb",
					Language: string(job.PreferredLang),
					Error:    nil,
				}
				return
			} else {
				logf("phase 1: subdb failed: %v", err)
			}
		}

		// Phase 2: OpenSubtitles (IMDB ID, fallback)
		if job.ImdbID != "" {
			if content, filename, err := e.tryOSDownload(job.ImdbID, job.PreferredLang, torrentName, job.VideoPath); err == nil {
				logf("phase 2: opensubtitles success → %s", filename)
				ch <- Result{
					Content:  content,
					Filename: filename,
					Provider: "opensubtitles",
					Language: string(job.PreferredLang),
					Error:    nil,
				}
				return
			} else {
				logf("phase 2: opensubtitles failed: %v", err)
			}
		}

		// Neither provider found a subtitle
		logf("no subtitle available for %s", torrentName)
		ch <- Result{
			Content:  nil,
			Filename: "",
			Provider: "",
			Language: "",
			Error:    fmt.Errorf("no subtitle available for %s", torrentName),
		}
	}()

	// Set a deadline: total timeout = max(SubDBTimeout, OSDownloadTimeout), capped at 30s
	go func() {
		totalTimeout := e.cfg.SubDBTimeout + e.cfg.OSDownloadTimeout
		if totalTimeout > 30*time.Second {
			totalTimeout = 30 * time.Second
		}
		time.Sleep(totalTimeout)
		select {
		case ch <- Result{
			Content:  nil,
			Filename: "",
			Provider: "",
			Language: "",
			Error:    fmt.Errorf("subtitle download timed out after %v", totalTimeout),
		}:
		default:
		}
	}()
}

// trySubDBDownload searches and downloads a subtitle via SubDB.
func (e *SubtitleEngine) trySubDBDownload(videoHash string, lang LanguageTag, torrentName, videoPath string) ([]byte, string, error) {
	// Rate limit check (15 tokens / 15 minutes)
	if !e.tryAcquireSubDBToken() {
		return nil, "", fmt.Errorf("subdb rate limited (15/15min, retry later)")
	}
	defer e.releaseSubDBToken()

	// Search for available languages
	results, err := e.subdb.Search(videoHash, "", torrentName, lang, e.cfg.MaxResultsPerProvider)
	if err != nil {
		return nil, "", fmt.Errorf("subdb search: %w", err)
	}
	if results == nil || len(results) == 0 {
		return nil, "", fmt.Errorf("subdb: no languages available for hash %s", videoHash[:min(len(videoHash), 8)])
	}

	// Select best match using language preference + release profile
	bestIdx := SelectBestSubtitle(results, torrentName, e.cfg.PreferredLanguages)
	if bestIdx < 0 || bestIdx >= len(results) {
		bestIdx = 0
	}

	sub := results[bestIdx]

	// Download the subtitle
	content, err := e.subdb.Download(sub)
	if err != nil {
		return nil, "", fmt.Errorf("subdb download: %w", err)
	}

	// Build the .srt filename for this video+language
	srtPath := e.writer.BuildSRTPath(videoPath, sub.Language)

	return content, srtPath, nil
}

// tryOSDownload searches and downloads a subtitle via OpenSubtitles.
func (e *SubtitleEngine) tryOSDownload(imdbID string, lang LanguageTag, torrentName, videoPath string) ([]byte, string, error) {
	// Search for subtitles
	results, err := e.os.Search("", imdbID, torrentName, lang, e.cfg.MaxResultsPerProvider)
	if err != nil {
		return nil, "", fmt.Errorf("os search: %w", err)
	}
	if results == nil || len(results) == 0 {
		return nil, "", fmt.Errorf("os: no subtitles found for %s", imdbID)
	}

	// Select best match
	bestIdx := SelectBestSubtitle(results, torrentName, e.cfg.PreferredLanguages)
	if bestIdx < 0 || bestIdx >= len(results) {
		bestIdx = 0
	}

	sub := results[bestIdx]

	// Download
	content, err := e.os.Download(sub)
	if err != nil {
		return nil, "", fmt.Errorf("os download: %w", err)
	}

	srtPath := e.writer.BuildSRTPath(videoPath, sub.Language)

	return content, srtPath, nil
}

// tryAcquireSubDBToken implements the 15-token/15-minute sliding window.
func (e *SubtitleEngine) tryAcquireSubDBToken() bool {
	e.rateMu.Lock()
	defer e.rateMu.Unlock()

	now := time.Now()

	// Refill tokens if enough time has passed
	elapsed := now.Sub(e.rateRef)
	if elapsed >= 15*time.Minute {
		e.rateTok = 15
		e.rateRef = now
	} else if elapsed >= time.Minute {
		// Gradual refill: 1 token per minute
		tokensToAdd := int(elapsed.Minutes())
		if tokensToAdd > 0 && e.rateTok < 15 {
			e.rateTok += tokensToAdd
			if e.rateTok > 15 {
				e.rateTok = 15
			}
			e.rateRef = now
		}
	}

	if e.rateTok <= 0 {
		return false
	}

	e.rateTok--
	return true
}

// releaseSubDBToken releases a token (for error cases where we didn't actually download).
func (e *SubtitleEngine) releaseSubDBToken() {
	e.rateMu.Lock()
	defer e.rateMu.Unlock()
	if e.rateTok < 15 {
		e.rateTok++
	}
}

// Min helper
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// CleanFilename returns the base name of the file without extension.
func CleanFilename(path string) string {
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	return base[:len(base)-len(ext)]
}

// FileExists checks if a file exists at the given path.
func (e *SubtitleEngine) FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// DeleteCachedSubtitle removes a cached .srt file and allows re-fetching.
// Called by the manual "resync subtitle" command.
func (e *SubtitleEngine) DeleteCachedSubtitle(videoPath string, lang LanguageTag) bool {
	srtPath := e.writer.BuildSRTPath(videoPath, lang)
	if e.FileExists(srtPath) {
		if err := os.Remove(srtPath); err != nil {
			logf("failed to delete cached subtitle %s: %v", srtPath, err)
			return false
		}
		logf("deleted cached subtitle: %s", srtPath)
		return true
	}
	return false
}
