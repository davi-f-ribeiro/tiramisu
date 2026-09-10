package subprovider

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// SubDBProvider implements Provider for thesubdb.com.
// SubDB is hash-based (MD5 of first+last 64KB of the video file).
// No authentication required. Rate limit: 15 downloads / 15 min.
type SubDBProvider struct {
	userAgent string
	client    *http.Client
}

// NewSubDBProvider creates a SubDB provider with sensible defaults.
func NewSubDBProvider() *SubDBProvider {
	return &SubDBProvider{
		userAgent: "Tiramisu/" + getAppVersion() + " (https://github.com/MrRobotoGit/tiramisu)",
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (p *SubDBProvider) Name() string { return "subdb" }

// CanMatch: true only when we have a valid video hash.
func (p *SubDBProvider) CanMatch(videoHash, imdbID, torrentName string) bool {
	return videoHash != ""
}

// Search queries SubDB for available languages for the given hash.
func (p *SubDBProvider) Search(videoHash, imdbID, torrentName string, lang LanguageTag, limit int) ([]Subtitle, error) {
	url := fmt.Sprintf("https://api.subdb.com/?action=search&hash=%s", videoHash)

	resp, err := p.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("subdb search GET: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("subdb search: HTTP %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("subdb search read: %w", err)
	}

	// Parse XML response: each available language is a <p language="xx"/> element
	langs := parseSubDBXML(string(body))
	if len(langs) == 0 {
		return nil, nil
	}

	// Filter for preferred language, rank results
	var results []Subtitle
	for _, l := range langs {
		score := 0
		if l == string(lang) {
			score = 100
		} else if l == "por" || l == "pt" {
			score = 90
		} else if l == "multi" || l == "multilang" {
			score = 80
		} else if l == "en" || l == "eng" {
			score = 60
		} else {
			score = 20
		}
		if score == 0 {
			continue
		}
		results = append(results, Subtitle{
			ID:           videoHash,
			ProviderName: "subdb",
			Language:     LanguageTag(l),
			Format:       "srt",
			VideoHash:    videoHash,
			Score:        score,
		})
	}

	return results, nil
}

// Download fetches the subtitle for the given hash and language.
func (p *SubDBProvider) Download(sub Subtitle) ([]byte, error) {
	url := fmt.Sprintf(
		"https://api.subdb.com/?action=download&hash=%s&language=%s",
		sub.VideoHash, sub.Language,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("subdb download request: %w", err)
	}
	req.Header.Set("User-Agent", p.userAgent)

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("subdb download GET: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusPreconditionFailed || resp.StatusCode == http.StatusTooManyRequests {
		return nil, fmt.Errorf("subdb rate limited (429/412)")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("subdb download: HTTP %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("subdb download read: %w", err)
	}

	if len(body) == 0 {
		return nil, fmt.Errorf("subdb download: empty response (not found)")
	}

	return body, nil
}

// parseSubDBXML parses the XML response from SubDB search and returns
// a list of language codes (e.g. ["en", "pt", "es"]).
func parseSubDBXML(xmlBody string) []string {
	var langs []string
	const prefix = `language="`
	for {
		idx := strings.Index(xmlBody, prefix)
		if idx == -1 {
			break
		}
		start := idx + len(prefix)
		end := strings.IndexByte(xmlBody[start:], '"')
		if end == -1 {
			break
		}
		code := xmlBody[start : start+end]
		if len(code) > 0 {
			langs = append(langs, code)
		}
		xmlBody = xmlBody[start+end+1:]
	}
	return langs
}

// ComputeSubDBHash computes the MD5 hash used by SubDB.
// The hash is MD5(bytes(first_64KB) + bytes(last_64KB)).
func ComputeSubDBHash(first64KB, last64KB []byte, fileSize int64) string {
	var data []byte
	if fileSize < 131072 { // 128 KB
		data = first64KB
	} else {
		data = make([]byte, 0, len(first64KB)+len(last64KB))
		data = append(data, first64KB...)
		data = append(data, last64KB...)
	}
	sum := md5.Sum(data)
	return fmt.Sprintf("%x", sum)
}
