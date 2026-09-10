package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"tiramisu/internal/prowlarr"

	"github.com/google/uuid"
)

// NatPMPConfig holds the configuration for NAT-PMP port forwarding.
type NatPMPConfig struct {
	Enabled      bool   `json:"enabled"`
	Gateway      string `json:"gateway"`
	LocalPort    int    `json:"local_port"`
	VPNInterface string `json:"vpn_interface"`
	Lifetime     int    `json:"lifetime"`
	Refresh      int    `json:"refresh"`
}

// DailyJobConfig: task that can run on specific days of the week.
// DaysOfWeek uses JS convention: 0=Sunday … 6=Saturday.
type DailyJobConfig struct {
	Enabled    bool  `json:"enabled"`
	DaysOfWeek []int `json:"days_of_week"` // 0=Sun, 1=Mon, …, 6=Sat
	Hour       int   `json:"hour"`
	Minute     int   `json:"minute"`
}

type WatchlistSyncConfig struct {
	Enabled       bool `json:"enabled"`
	IntervalHours int  `json:"interval_hours"` // 1,2,3,4,6,8,12,24
}

type SchedulerConfig struct {
	Enabled       bool                `json:"enabled"`
	MoviesSync    DailyJobConfig      `json:"movies_sync"`
	TVSync        DailyJobConfig      `json:"tv_sync"`
	WatchlistSync WatchlistSyncConfig `json:"watchlist_sync"`
}

// EngineConfig holds per-engine paths for subprocess sync.
type EngineConfig struct {
	ScriptPath string
	LogsDir    string
}

// MovieWeights holds the scoring weights and size gates used by the movie engine
// (and by the watchlist, which is movies-only). Defaults are the values that were
// compiled into internal/syncer/engines/movie_go.go before they became configurable.
type MovieWeights struct {
	Res4K                int `json:"res_4k"`
	Res1080p             int `json:"res_1080p"`
	HDR                  int `json:"hdr"`
	DolbyVision          int `json:"dolby_vision"`
	Atmos                int `json:"atmos"`
	Audio51              int `json:"audio_5_1"`
	StereoPenalty        int `json:"stereo_penalty"`
	Remux                int `json:"remux"`
	PreferredLanguage    int `json:"preferred_language"`
	UnknownSize4KPenalty int `json:"unknown_size_4k_penalty"`
	SeederCap            int `json:"seeder_cap"`
	MinSeeders           int `json:"min_seeders"`
	Min4KGB              int `json:"min_4k_gb"`
	Max4KGB              int `json:"max_4k_gb"`
	Min1080pGB           int `json:"min_1080p_gb"`
	Max1080pGB           int `json:"max_1080p_gb"`
}

// TVWeights holds the scoring weights and gates used by the TV engine. The seeder
// bonus is tiered rather than proportional, which is why it does not share a struct
// with MovieWeights.
type TVWeights struct {
	Res4K             int `json:"res_4k"`
	Res1080p          int `json:"res_1080p"`
	HDR               int `json:"hdr"`
	DolbyVision       int `json:"dolby_vision"`
	Atmos             int `json:"atmos"`
	Audio51           int `json:"audio_5_1"`
	PreferredLanguage int `json:"preferred_language"`
	Fullpack          int `json:"fullpack"`
	SeederTier100     int `json:"seeder_tier_100"`
	SeederTier50      int `json:"seeder_tier_50"`
	SeederTier20      int `json:"seeder_tier_20"`
	MinSeeders        int `json:"min_seeders"`
	MinSeeders4K      int `json:"min_seeders_4k"`
	// SeasonSkipScore is measured against averages of QualityScore, so it must move
	// with the weights that build that score - a fixed number would drift off the scale.
	SeasonSkipScore int `json:"season_skip_score"`
}

// QualityScoringConfig carries optional per-profile overrides. A nil profile means
// "use the defaults" - that is what keeps existing installs on their current picks
// without writing anything into their config.json.
type QualityScoringConfig struct {
	Movies *MovieWeights `json:"movies,omitempty"`
	TV     *TVWeights    `json:"tv,omitempty"`
}

// DefaultMovieWeights returns the shipped movie scoring profile.
func DefaultMovieWeights() MovieWeights {
	return MovieWeights{
		Res4K: 1000, Res1080p: 200, HDR: 60, DolbyVision: 100,
		Atmos: 50, Audio51: 25, StereoPenalty: -50, Remux: 30,
		PreferredLanguage: 60, UnknownSize4KPenalty: -5,
		SeederCap: 50, MinSeeders: 15,
		Min4KGB: 10, Max4KGB: 40, Min1080pGB: 4, Max1080pGB: 20,
	}
}

// DefaultTVWeights returns the shipped TV scoring profile.
func DefaultTVWeights() TVWeights {
	return TVWeights{
		Res4K: 1000, Res1080p: 200, HDR: 100, DolbyVision: 150, Atmos: 50, Audio51: 25,
		PreferredLanguage: 40, Fullpack: 500,
		SeederTier100: 100, SeederTier50: 50, SeederTier20: 10, MinSeeders: 5,
		MinSeeders4K: 5, SeasonSkipScore: 1000,
	}
}

// MovieWeights resolves the profile to use: the configured one when present, the
// defaults otherwise. A present profile is taken verbatim, so an explicit 0 stays 0.
func (q QualityScoringConfig) MovieWeights() MovieWeights {
	if q.Movies == nil {
		return DefaultMovieWeights()
	}
	return *q.Movies
}

// TVWeights resolves the TV profile the same way.
func (q QualityScoringConfig) TVWeights() TVWeights {
	if q.TV == nil {
		return DefaultTVWeights()
	}
	return *q.TV
}

// LanguageConfig controls preferred/excluded audio-language matching used
// by the Movie and TV sync engines when scoring and filtering torrents.
type LanguageConfig struct {
	// PreferredTerms are case-insensitive, word-boundary-matched release-name terms (e.g. "ita", "multi", "dual").
	PreferredTerms []string `json:"preferred_terms"`
	// PreferredFlags/ExcludedFlags are ISO 3166-1 alpha-2 codes matched against flag emoji in indexer result lines.
	PreferredFlags []string `json:"preferred_flags"`
	ExcludedFlags  []string `json:"excluded_flags"`
}

// Config holds all configurable parameters for the FUSE proxy
type Config struct {
	// --- Internal / Derived Fields ---
	ConfigPath string `json:"-"`
	// LogDir holds the log files the dashboard tails. Docker points TIRAMISU_LOG_DIR at
	// a mounted volume; elsewhere the logs sit next to config.json.
	LogDir   string `json:"-"`
	RootPath string `json:"-"` // V138: Root path for state/config (default: /home/pi)

	// --- Core Tuning (JSON Mapped) ---
	MasterConcurrencyLimit int    `json:"master_concurrency_limit"` // Global limit for concurrent HTTP requests to GoStorm
	ReadAheadBudgetMB      int64  `json:"read_ahead_budget_mb"`     // Global budget for read-ahead in MB
	MetadataCacheSizeMB    int64  `json:"metadata_cache_size_mb"`   // Size of metadata LRU cache in MB (V178)
	FuseBlockSize          int    `json:"fuse_block_size_bytes"`
	StreamingThresholdKB   int64  `json:"streaming_threshold_kb"`
	LogLevel               string `json:"log_level"`

	// --- FUSE Timing ---
	AttrTimeoutSeconds     float64 `json:"attr_timeout_seconds"`
	EntryTimeoutSeconds    float64 `json:"entry_timeout_seconds"`
	NegativeTimeoutSeconds float64 `json:"negative_timeout_seconds"`

	// --- HTTP Resilience ---
	MaxRetryAttempts         int `json:"max_retry_attempts"`
	RetryDelayMS             int `json:"retry_delay_ms"`
	RescueGracePeriodSeconds int `json:"rescue_grace_period_seconds"`
	RescueCooldownHours      int `json:"rescue_cooldown_hours"`

	// --- Preload Engine ---
	PreloadWorkersCount   int `json:"preload_workers_count"`
	PreloadInitialDelayMS int `json:"preload_initial_delay_ms"`
	WarmStartIdleSeconds  int `json:"warm_start_idle_seconds"`
	MaxConcurrentPrefetch int `json:"max_concurrent_prefetch"`

	// --- Cache Management ---
	CacheCleanupIntervalMin int `json:"cache_cleanup_interval_min"`
	MaxCacheEntries         int `json:"max_cache_entries"`

	// --- Connectivity ---
	GoStormBaseURL   string `json:"gostorm_url"`
	ProxyListenPort  int    `json:"proxy_listen_port"`
	MetricsPort      int    `json:"metrics_port"`
	BlockListEnabled bool   `json:"blocklist_enabled"`
	BlockListURL     string `json:"blocklist_url"`
	// BlockListFilter keeps only the ranges whose description matches this regexp.
	// Empty keeps the whole list. See blockedIP.go for why published lists need it.
	BlockListFilter string `json:"blocklist_filter"`

	AIURL           string `json:"ai_url"`      // V1.4.5: AI Optimizer sidecar URL
	AIProvider      string `json:"ai_provider"` // V1.7.1: Provider type (local, openrouter, openai)
	AIModel         string `json:"ai_model"`    // V1.7.1: Model ID for cloud providers
	AI_API_KEY      string `json:"ai_api_key"`  // V1.7.1: API key for cloud providers

	// --- FUSE Paths ---
	// Fallback when CLI args are omitted. CLI args always take precedence.
	PhysicalSourcePath string `json:"physical_source_path"` // Real MKV dir (e.g. /mnt/torrserver)
	FuseMountPath      string `json:"fuse_mount_path"`      // FUSE virtual mount (e.g. /mnt/torrserver-go)

	// --- Legacy Compatibility Fields (populated from above) ---
	DefaultFileSize         int64         `json:"-"`
	ReadAheadBudget         int64         `json:"-"`
	MetadataCacheSize       int64         `json:"-"` // V178
	ReadAheadBase           int64         `json:"-"`
	ReadAheadInitial        int64         `json:"-"`
	StreamingThreshold      int64         `json:"-"`
	SequentialTolerance     int64         `json:"-"`
	MaxConcurrentHTTP       int           `json:"-"`
	RateLimitRequestsPerSec int           `json:"-"`
	PreloadWorkers          int           `json:"-"`
	MaxConnsPerHost         int           `json:"-"`
	ConcurrencyLimit        int           `json:"-"`
	KeepaliveInterval       time.Duration `json:"-"`
	KeepaliveIdleStart      time.Duration `json:"-"`
	KeepaliveMaxIdle        time.Duration `json:"-"`
	CacheTTL                time.Duration `json:"-"`
	UID                     uint32        `json:"-"`
	GID                     uint32        `json:"-"`

	// --- Disk Warmup ---
	DiskWarmupQuotaGB int64 `json:"disk_warmup_quota_gb"` // Total SSD quota for warmup cache (default: 32)
	// Deprecated: warmupFileSize is now hardcoded at 64MB. Field kept for
	// backward-compatible JSON unmarshal of existing config.json files.
	WarmupHeadSizeMB int64 `json:"warmup_head_size_mb"`

	// --- NAT-PMP (V228) ---
	NatPMP NatPMPConfig `json:"natpmp"`

	// --- External Services (V1.4.6) ---
	Plex struct {
		URL         string `json:"url"`
		Token       string `json:"token"`
		LibraryID   int    `json:"library_id"`
		TVLibraryID int    `json:"tv_library_id"`
	} `json:"plex"`
	TMDBAPIKey   string `json:"tmdb_api_key"`
	TorrentioURL string `json:"torrentio_url"` // Torrentio base URL (used when Prowlarr is disabled)

	// --- Prowlarr Indexer ---
	Prowlarr prowlarr.ConfigProwlarr `json:"prowlarr"`

	// --- Built-in Sync Scheduler ---
	Scheduler SchedulerConfig `json:"scheduler"`

	// --- Media Server ---
	MediaServerType string `json:"media_server_type"` // "plex" | "jellyfin"

	// --- Quality Scoring ---
	QualityScoringConfig QualityScoringConfig `json:"quality_scoring"`

	// --- Language Matching ---
	Language LanguageConfig `json:"language"`

	// --- Subtitle Provider (D1-D5: Rota C) ---
	Subtitle struct {
		Enabled      bool   `json:"enabled"`
		APIKey       string `json:"api_key"`        // OpenSubtitles REST API key
		BaseURL      string `json:"base_url"`       // Optional custom base URL
		User         string `json:"user"`           // Optional username for JWT
		Password     string `json:"password"`       // Optional password for JWT
		Preferred    []string `json:"preferred_languages"` // e.g. ["por", "multi", "eng"]
		MaxResults   int    `json:"max_results"`    // default: 5
	} `json:"subtitle"`

	// --- Engine Scripts (populated in LoadConfig, not from JSON) ---
	EngineScripts map[string]EngineConfig `json:"-"`

	// --- Telemetry (V1.4.7) ---
	TelemetryID     string `json:"telemetry_id"`
	EnableTelemetry bool   `json:"telemetry"`
	TelemetryURL    string `json:"telemetry_url"`

	// --- State DB (V1.7.1) ---
	EnableStateDB bool   `json:"enable_state_db"` // default: true
	StateDBPath   string `json:"state_db_path"`   // default: <STATE>/tiramisu.db
}

// Save persists the current configuration to config.json
func (c *Config) Save() error {
	// 1. Marshal config to JSON
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	// 2. Write to file
	return os.WriteFile(c.ConfigPath, data, 0644)
}

// DefaultBlockListURL is the list the filter below is written for. It is a real default,
// not just an example in config.json.example: a configuration without the key showed an
// empty field in the Control Panel, and an empty field is skipped on save, so enabling the
// blocklist downloaded nothing and there was no way to fix it from the panel.
const DefaultBlockListURL = "https://list.iblocklist.com/?list=ydxerpxkpcfqjaybcssw&fileformat=p2p&archiveformat=gz"

// DefaultBlockListFilter keeps only the anti-P2P section of a published blocklist. It is a
// real default, not a hint: without it Level 1 loads whole - 17% of IPv4, nearly all of it
// 1990s whois records - and rejects ordinary peers on netblocks that changed hands years
// ago. Set the key to an empty string in config.json to load a list unfiltered.
const DefaultBlockListFilter = `(?i)\bap2p\b|anti-?p2p`
