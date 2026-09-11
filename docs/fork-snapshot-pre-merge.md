### Pre-Merge Snapshot

Date: 2026-09-11T20:44:27-03:00
Branch: main

## Git Log (main)

218be5e fix(syncer): add IT (Italy) to excludedFlagLanguages map
f87d4e5 fix(subprovider): adjust 'tu' marker to match sentence-start position
446e510 test(subprovider): add 'Tu' at start-of-sentence test case
003ea17 test(subprovider): add unit tests for DetectPortugueseVariant
96aa0ea chore(config): set preferred_languages priority to pob (PT-BR) first
c4f5bd9 feat(subtitle): integrate subtitle engine with FUSE (Readdir/Lookup passthrough, Open() trigger, global engine lifecycle)
38a6a0a feat(subprovider): fix sidecar writing, add PT-BR/PT-PT detection, fix Run() channel bug
5770f9e fix(settings): subtitle toggle now uses <label class='switch'> pattern like other working toggles
c0be8a2 fix(settings): add width/height to .toggle-switch so .slider doesn't overflow
f0c92aa fix(subprovider): register global logger so all logf() calls are emitted
a94f4b3 fix: add missing CSS for btn-secondary, btn-sm, toggle-row, toggle-switch and handle 'main' branch in version compare
d14df92 fix: copy go.mod/go.sum for anacrolix-torrent before go mod download
1921c76 fix: update docker workflow - remove QEMU, add GHCR login, optimize Dockerfile cache
3d8658e fix: add Subtitle struct to Config (missing from Rota C)
9d22eee fix: add Subtitle struct to Config (missing from Rota C)
a9d8879 fix: add getDefaultConfig to config.go (moved from engine.go)
de729a9 fix: remove EngineConfig redeclaration and unused imports
9cc46b9 fix: add SubDBTimeout and OSDownloadTimeout fields to EngineConfig
1c97a0b fix: remove docker/login-action@v3 and set push: false
d61b89b fix: restore complete main.go and internal/subprovider package
808215d docs: link Hermes and OpenClaw, fix two stale anchors in the table of contents
a29d8a2 docs: refresh the API reference and name OpenClaw and Claude Code as skill hosts
4ebb50e docs: add AI Agent Skill to the table of contents
8135baf docs: 'Use when' description, surface the release tradeoff instead of silently gating it (3.0.3)
1024b2f docs: deletion goes through FUSE so it blacklists, plus a bulk removal section (3.0.2)
e298ec2 docs: IMDB id resolution, Prowlarr-only endpoint, library scan, exclusive DV/HDR tags (3.0.1)
bc8b5a6 docs: duplicate check uses the filesystem plus the engine list, not the dashboard endpoint
039610a docs: search goes through Prowlarr via Tiramisu or Torrentio, not a generic indexer
6f8df19 revert: restore the em-dashes in the skill file, they were only unwanted in the README
09d304c docs: AI agent skill for adding a release to the library by hand
8e8523e chore: build the Docker image on Go 1.26, the 1.24 series no longer gets security patches
f02172b fix: with a single session the /metrics/ttff percentiles reported 1ms instead of the sample
346c6b4 refactor: to8mb and window(MB) named something other than what they measure, now deepread and protected/reader
c1be39c fix: a panic in cleanPieces killed a cache's eviction for good, the recover sat outside the loop
c73c656 fix: data races on the cache fill counters and PlaybackState, and a panic when the piece bitmap is not yet sized
babfc3d fix: no pump slot was granted to any playback with master_concurrency_limit at 5 or below
928cc43 fix: a corrupt settings file is set aside instead of blocking every save
cb54148 fix: do not repeat the unmounted-volume warning for STATE_DIR under ROOT_PATH
359ef1c fix: container state defaults to /state instead of the image layer
a008862 fix: a corrupt settings.json no longer wipes the stored settings
3b73eb1 perf: MarshalTo serialises without binary.Write, which allocated per field
4a2a11b fix: a failed Prowlarr search is no longer cached as "no such movie", and Torrentio is always considered
aa94421 fix: trim whitespace and CR when reading trackers.txt
91106b9 fix: the blocklist URL is a default too, not only an example
b499f8b feat: the movies and TV syncs notify Jellyfin too, not only Plex
171b986 feat: expose media_server_type in the Control Panel and name the fields for both products
04d2a4b fix: read the blocklist from the state directory, where it is downloaded
38835a8 docs: document the state directory contents and the container update sequence
066c932 fix: GoStorm state lands in the --path directory, not the working directory
28685d9 fix: the blocklist filter is a default, not a placeholder
e2abb2f feat: blocklist_filter keeps only the maintained part of a published list
490312a fix: default to the badpeers blocklist, and refresh immediately when the URL changes
684e71d feat: show how many IPs the blocklist keeps out, and fix Docker logs in the Control Panel
f47cf84 fix: disabling the blocklist now really disables it, and gzip is detected from content
7dcde11 fix: keep Docker settings across docker rm, and stop reporting a phantom Plex
85b4807 fix: restart from the Control Panel no longer stops the Docker container
2f933ee docs: trim the UDP rationale from the port forwarding steps
58d689b docs: document forwarding the peer port when the router has UPnP disabled
a10e386 feat: discover Plex and Jellyfin on the LAN instead of assuming Plex
dfa2676 docs: mount config.json read-write, the Control Panel saves to it
5dfc627 fix: give Prowlarr hash resolution 20s and no retry, so 1337x results survive
6cde1ab Add tiramisu icon for directory listings
6de17b4 revert: drop the episode-name check, ShowMatch already rejects other shows
4abb23f perf: recycle piece buffers per cache instead of a global pool
011846b fix: find new seasons in TV sync and reject wrong-show releases
a5c484b chore: gofmt and drop the orphan bolt piece completion
fcfcf1d fix: closing the cache frees the piece buffers without recycling them
3e5e986 fix: read Piece.Size and the piece map without racing
60cca31 fix: cap the DB read cache by bytes
d16c9f9 fix: read the log tail without loading the whole file
73bf8de fix: split the readahead per reader like the protected window
7951139 fix: cap the prefetch by bytes instead of piece count
5b8bfd8 fix: a blocking read waits for the block it asked for, not the whole window
93359d3 fix: the pump uses its captured reader, not the field Release nils
5bab6d2 fix: complete short reads instead of serving them as success
8203451 chore: drop the orphaned hls-proxy README
8fa69f3 feat: count short reads reported as success
29b7981 fix: keep every path when saving the inode map
9ea8dc2 fix: stop a distant probe from taking the pump away from the player
a2f1fa7 fix: keep the pump anchored on the player instead of losing track of it
1f55e31 fix: stop the eviction diagnostic from flooding the log
f2f7111 fix: retry the tracker list instead of giving up silently
a44ec2c feat: the TV contents check now rejects, not just reports
3d1b9ca feat: check a torrent's contents belong to the show
d4a42c0 fix: corrupted AI-Pilot prompt and panic on a malformed hash
2aded6b feat: build the preview sample names from the weights
119fc45 fix: show the season pack bonus in the TV preview
0f06f79 fix: reach GoStorm settings same-origin from the Control Panel
8e56f7c fix: don't decrement a pump this handle never counted
26c1c93 fix: enforce CacheSize again in the torrstor eviction
464a618 chore: neutral sample title in the Release Selection preview
56d9df0 fix: drop the misleading InodeMap line when StateDB is active
13515f5 fix: let the score decide between 4K and 1080p, not the grouping
ea29587 feat: live scoring preview in the Release Selection card
8b6a5aa feat: score Dolby Vision separately from HDR on TV
f41e761 feat: configurable release scoring weights
4520357 fix: let a streaming handle earn primary so the pump follows the player
2ff16ef fix: warn about missing AVX2 instead of refusing to start
110b569 feat: fill the MKV tail window sequentially while the file is idle
4424502 fix: negative cache never expired, hiding complete warmup files
e16e024 fix: head cache could serve zeros from write holes
8f4a667 fix: seek/resume precision — tail cache holes and pump sync thrashing
2940d75 feat: TTFF instrumentation and first-read hedge coverage fix
27fc8f6 docs: update docker-windows README to reflect the shipped installer
77beb4c chore: credit contributor for Windows Docker jellyfin fix (#20)
129de54 fix: Windows Docker jellyfin template never updated after Python->Go port
c50772d feat: pump-loop starvation watchdog with 10s rate window
340c988 fix: serve tail resume from SSD instead of network, unreachable tail-probe state
a396d11 feat: refuse to start on amd64 CPUs without AVX2
427bb46 fix: raise pool ceiling 32->64 so the 2x storm margin holds at high budgets
02d3664 fix: anchor raCache pool cap to the configured budget
2d25cfa fix: AdaptiveShield strictCycleCount never reset, raise escalation cap
8334501 fix: playerOff flip-flops between real position and 0, thrashes pump
ca50c74 docs: add GPLv3 license badge
c51720b docs: update license reference to GPL-3.0
e084317 docs: relicense to GPL-3.0
dbc9b6e docs: document FUSE mount-propagation fix for Docker (virtual dir empty)
66d3540 fix: re-anchor ping-pong kills the pump; lower post-confirmation sync gate
cbf0193 fix: LastSeekOff poisoned by tail-probe reads, breaks ResumeAnchor
2ae5867 fix: pump never follows player on attached (secondary) handles
595ecf2 fix: sample hedge latency during playback pressure, not just warmup
4992b32 fix: PEXChurn probe window, second-chance cooldown, host fallback key
562714d fix: raCache adaptive chunk drift and budget/pieceLens hygiene
2cecc29 fix: Docker entrypoint gostream->tiramisu rename, scheduler API 404 on enable
ea7b6f9 debug: remove warmupActive tracing instrumentation, investigation closed
bf9475a fix: persist sole-dirtier V304 bans immediately; doc fixes for hedge ceiling and chunk alignment
d21e7cc docs: add Awesome Selfhosted badge
1fa192b feat: stateless no-baseline ceiling fallback for tail-hedging (4s, pre-calibration)
bcfdacc fix: pump throttle and stream timeout ignored inferred/STRICT playback
4f0c622 feat: add Rotten Tomatoes "Movies at Home" as a movie discovery source
7c753fa refactor: improve search accuracy by adding optional release year parameter to Prowlarr queries and synchronizing TV stream classification logic
3463a6e feat: switch default IP blocklist to iblocklist Level 1, document V304 peer auto-ban
78dda62 refactor: increase churn cooldown duration from 10s to 30s to reduce redundant re-probes
638e047 feat: enable AggressivePeerManagement by default for improved streaming cold-start performance
3a1deff fix: improve blocklist update robustness with timeouts, atomic writes, and dynamic DHT updates
c52a33f feat: make IP blocklist download opt-in via dashboard toggle
5710e2e fix: add nil checks for bts to prevent panic in apihelper methods
1f06949 feat: implement live IP blocklist hot-swapping and fix range sorting during ingestion
9600b91 fix: improve regex quality boundaries and integrate torrent removal into stub cleanup across sync engines
c920209 feat: implement hard language exclusion via TMDB code filtering and expanded title regex matching
724fe9a fix: PeerEject log showed duration overflow for peers that never delivered a chunk
ed0cb94 feat: persist V304 peer bans across restarts, count corruption per IP
3420ca5 feat: proactive EWMA outlier peer ejection
e7da32c feat: singleflight for concurrent FetchBlock misses in Read()
1d63317 chore: remove dead HTTP client config fields orphaned by retry consolidation
0f4284d feat: make HTTP retry configurable, add retry to Prowlarr/GoStorm/Collector
d9a1a2c feat: make preferred/excluded audio language configurable
190f52f fix: invalidate FUSE state when sync removes stub files
b0f9833 fix: recycle read-ahead buffers for adaptive chunk sizes
717bad2 docs: add demo GIF showing instant playback start
7b5cdf8 chore: remove dead code confirmed unreachable by deadcode analysis
3814dea feat: generalize tail-hedging to playback pressure, expose banned peer count
03b97cb docs: tweak tagline phrasing
6e1573b docs: fix H1 bold using strong tag instead of markdown asterisks
76fb951 docs: merge H1 into tagline sentence instead of separate centered heading
e6af6d8 docs: use nested sub tags to shrink H1 visually since GitHub strips style attrs
e602a53 docs: add semantic H1 for SEO, add FAQ section for GEO
155c889 docs: update Adaptive Shield description to reflect escalating clean-streak (30/60/90/120s)
35dc061 docs: add Adaptive Chunk Size, AdaptiveShield, TailHedge, PEXChurn to features; remove all em-dashes
a923fe6 feat: expose and visualize V304 banned peer count in monitoring dashboard
b04d506 docs: cache-bust install.png reference
275329b docs: update install.png screenshot with Tiramisu ASCII banner
e85f9fa docs: update control and health monitor dashboard screenshots
9f8023e docs: restyle README header — ASCII banner, badges up top, GitHub alert callout
8978be3 rename: macOS toolbar app GoStream.app -> Tiramisu.app
5d1b805 fix: complete gostream->tiramisu rename in hls-proxy README
68f6049 fix: run fuse.conf user_allow_other setup unconditionally in install.sh
b2e0db2 rename: Docker Hub/GHCR image gostream -> tiramisu
f644e87 fix: proper heading markup for Tiramisu title in README
d5ccc2c rename: fix leftover gostream references excluded from private->public sync
790223e rename: update GitHub repo URLs after MrRobotoGit/gostream -> tiramisu
2392390 rename: gostream -> tiramisu
f97e644 docs: update health monitor screenshot to reflect current UI
19f17e2 tune: lower PEXChurn peer cooldown 30s -> 10s
0917992 style: remove 'Last update' timestamp from dashboard header
5cd7d23 style: shorten 'Cache Hit Rate' label to 'Hit Rate'
d7b96d0 feat: torrent hash in TailHedge/PEXChurn logs + churn cooldown; calibrate breaker threshold
aaf67df fix: dev-build version badge shows plain text, not a false-alarm pill
794c7d6 fix: show — for Cache Hit Rate when no torrent active, not stale value
de8fdfd feat: dashboard mobile/desktop font sizing + cache hit rate; lower 4K max size cap
cde46d4 fix: WARMUP state never transitioned to STREAMING during linear playback
41e0cb7 fix: PEXChurn/TailHedge log visibility and circuit breaker idempotency
58a2d2c feat: aggressive peer management for warmup cold-start latency
d690453 perf: non-blocking semaphore check before spawning prefetch goroutines
97900c7 perf: reuse piece-hasher goroutines instead of spawning one per piece
9119fc1 perf: backport upstream fix removing extra per-tracker goroutine
61e8e2c refactor: make SetPieceLen explicit when piece length exceeds ReadAheadBase
a7b64c2 fix: wire up dead CloseHash() call, fixing critical piece-buffer memory leak
c0354f9 docs: fix stale README values (CacheSize, movie sync thresholds)
212bba4 fix: preserve imdb field on torrent rehydration
9732f9a fix: bound stream reads by context timeout, fix third adaptive-chunk-size gap
3e21bf8 fix: WarmLanding prefetch uses adaptive raCache.ChunkSize instead of fixed ReadAheadBase
39648fc fix: strict mode segments drawn as oblongs with rounded corners, bar height 8px
41d861b feat: add AdaptiveShield timeline to health monitor dashboard
9d8160c fix: make raCache chunk size adaptive and consistent with pump
f2b67f7 feat: adaptive chunk size, piece len logging, AdaptiveShield cycle persistence
556c0e7 fix: tighten StrategicReserve scan limit to 5 slots during active healthy playback
2c3cbfb fix: filter listActiveTorrents to streaming-only torrents to prevent 4K serialization
4e96ff9 fix: replace ListTorrent with ListActiveTorrent in CleanupManager to eliminate OOM leak
a0f5067 fix: lower tvMinSeeders4K from 10 to 5 to match non-4K threshold
e9329b8 fix: reset processedThisRun+stats per run and rehydrate before cleanup in TV sync
fa2cb44 fix: move interruptPending reset inside pump loop to enable ResetShield on repeated seeks
7f39695 fix: restore audio sync on resume and eliminate smbd D-state glitches
ed953b0 fix: prevent pump interrupt cascade when multiple handles share the same file
282790d refactor: improve concurrency safety with atomic status flags and explicit synchronization in torrent server and cache management
c48667f refactor: improve thread-safety and prevent race conditions by replacing boolean fields with atomics and adding missing mutex protection.
3b0508c fix: invalidate read-ahead cache on seek to prevent stale chunk usage
5363b9b refactor: rename internal/anacrolix-torrent-v1.55 → internal/anacrolix-torrent
b3b4300 chore: remove orphan CHANGELOG.md (replaced by GitHub Releases)
5080be2 refactor: extract root files into internal/ packages
4db2023 refactor: integrate stream filtering into getMovieStreams and refine cache logic based on raw stream availability
dd0cf0e fix: retain priority during unexpected pump exit if playback is confirmed healthy
ce67cb1 refactor: update playback state persistence logic and improve cleanup using COALESCE timestamps
636a7bb feat: implement playback state persistence, inference-based detection, and optimize Prowlarr search results ranking
3db5c4d feat: enhance TV discovery with seasonal keyword queries and improve movie stream filtering with classification reason logging.
af02fa0 fix: implement episode validation against TMDB metadata and optimize Prowlarr/movie recheck logic
85abe2a feat: add native macOS toolbar monitor application and documentation
8840115 Revert "fix: prioritize successful read returns over interrupt checks in native_bridge to prevent premature stream termination"
d1ab4fe fix: update Plex session matching to check both 8-character hash prefixes and suffixes
853b8ea feat: add logging for skipped episodes and fullpack quality checks in tv_go engine
50776b1 fix: adjust seeder validation logic, defer file removal until successful creation, and implement automatic cleanup of empty directories
b3dd4df refactor: migrate size parsing to Prowlarr API and update garbage pattern regex
cfec945 refactor: implement stream filtering for movies and add recency checks and quality optimization for TV shows
f4d8893 fis: Torrentio - Cloudflare check fixed
4f4525b refactor: remove byte-to-kilobyte conversion logic for rate limits in settings form
5cb7d68 docker: remove Python health-monitor from entrypoint
b736953 chore: update default disk warmup quota and fix nullish coalescing for engine settings input values
0761555 fix: resolve shadowing of globalJsonDB variable in GetInstance
a473a49 fix: better default settings
65d435a fix: resolve Prowlarr no-hash results and remove TV ghost registry entries
9bc41fc refactor: implement granular database migration, fix registry persistence logic, and resolve LRU cache data race condition.
ab73b6e fix: prioritize successful read returns over interrupt checks in native_bridge to prevent premature stream termination
9adfd15 fix: integrate SQLite backend for TV episode registry
79ce426 docs: update state persistence references from JSON to SQLite
179b624 remove deprecated swarm_evaluator
ad33a02 fix: remove unused imports and undefined OnBadSwarm in swarm_evaluator.go
d570854 fix: sync go.mod/go.sum for SQLite State DB dependencies
5ac1e7d SQLite State DB for unified persistence
b71bba3 refactor: remove swarm evaluation logic from AI tuner and clean up install script comments
82c8d4e feat: add OpenTracker for O(1) handle tracking per hash and path
91a5551 docs: update installation instructions to use the install script instead of git clone
cc1a43b docs: update installation screenshot to reflect current UI
1b3263b docs: update ai-pilot documentation to reflect dual-provider support and refined tuning logic
37144ae feat: Initial implementation AI-driven swarm health evaluation
a7656ad feat: overhaul install script UI with 256-color support, progress bars, and branded ASCII logo
53008bb refactor: simplify installer by making Samba optional and removing interactive Plex/API configuration steps
39f9e35 chore: downgrade go version to 1.24 in go.mod
d7bcad6 fix: add stop functionality and dynamic button states to library sync tasks
6a53af0 refactor: remove obsolete pump lifecycle and streaming flow unit tests
a285d4d fix: add stream termination UI, Plex thumbnail proxy, and improved sync logging while removing update check badge.
df3c45d fix: prevent deadlock in NativeReader.Close by closing pipeWriter before acquiring lock
eaf36c1 feat: implement automated update checking and UI notification badges for dashboard and settings
6ae80d4 docs: update README to reflect broader hardware support beyond Raspberry Pi 4
14e332f feat: implement cached file counting and public IP resolution in monitor collector with dashboard updates
69f13f5 feat: enhance Plex session matching with authoritative media info and improve torrent title cleaning
d36557c refactor: replace Python sync scripts and health monitor with pure Go
d46d5ab feat: virtual .mkv files — line-based → JSON format
e0110d6 feat: migrate Prowlarr client to Go and expose as a local API endpoint for sync scripts
3823c2e fix: improve request strategy unverified accounting and soft-interrupt watchdog
19bf152 feat: batch SetPriority + bitmap eviction optimization for torrstor cache
db187cc chore: reset AppVersion to dev for local development
cdae472 fix: conditionally apply ldflags to allow fallback to hardcoded version when no git tag is present
40536e6 feat: embed git tag version into binary during build process
3ba5d84 feat: implement dynamic versioning via build-time ldflags and application variable
fa4cd3c fix: improve SSD warmup logic by checking coverage range and increase threshold to right size
a27c3d5 chore: update gostream logo asset
9e5a2b7 chore: update gostream logo asset
64c27c3 chore: update gostream logo asset
db4f7b4 docs: remove redundant line break from README header
aa2e076 chore: update gostream logo asset
9e09757 chore: update gostream logo asset
cb57d20 logo
bf07a65 gostream logo
417df31 chore: update telemetry version to 1.6.3
73f805d feat: implement anonymous telemetry heartbeat with configurable ID generation
87f07e4 refactor: update warmup logic to trigger async wake based on head availability instead of full warmup status
4dbe5b2 refactor: replace boolean warmup flags with atomic state machine and add debounced timer management for pump cleanup
7d46987 feat: implement idle state tracking to trigger AI reset on fresh torrent starts
07c2940 fix: implement atomic reset locking, post-reset cycle skipping, and automatic default restoration in AI tuner
676ecdb refactor: remove legacy version comments and cleanup code documentation in main.go
4bea44c fix: reset stale resumeOffset when it exceeds warmupFileSize to prevent cache dead zones
15c6dd4 fix: remove redundant offset reset during tail probe disk warmup
5cc614d fix: prevent premature stale offset reset and refine warmup eligibility threshold during initial reads
c79bad5 feat: implement disk warmup bypass to accelerate initial playback by reading from SSD instead of torrent
8bf9207 refactor: implement atomic warmup eligibility to disable SSD reads during seeks and resumes
6b8b996 fix: INFUSE fix for seek/resume when WARMUP is enabled:
7923914 docs: add documentation for Windows Docker installer and setup instructions
3c40ef3 Merge pull request #15 from JaredJomar/docker
b73b147 Merge branch 'MrRobotoGit:main' into docker
87778d2 feat: add Windows Docker installer workflow and harden container config loading
2aa2590 chore: remove main branch trigger from docker-publish workflow
f9c055a chore: update Docker Hub short description for better clarity
76926a1 docs: Add Docker pull badges, instructions for pre-built Docker images, and update the Docker run command.
8176d64 feat: Display total peers alongside active peers in AI tuner's metrics snapshot.
9ce1a9a typo
dc47ef0 docs: update Docker Hub short description in publish workflow.
b125c11 feat: add GitHub Actions workflow to build and publish Docker images to Docker Hub and GitHub Container Registry.
2319123 fix: CIFS false pump interrupts and AdaptiveShield escalation
91d6d8b feat: Introduce explicit rejection for non-European, non-EN/IT languages and refine language filtering logging.
65217d6 feat: Introduce a circuit breaker for AI communication, reduce the AI client timeout, and add a dedicated short-timeout client for KV cache resets.
5d7eb91 feat: Add title parameter to Prowlarr TV series search for improved 4K detection.
5bde11f config: Increase `MIN_EPISODE_SIZE` from 200MB to 1GB.
1894573 refactor: Extract Prowlarr API query into `_query` and merge HD/UHD search results by `infoHash`.
144bf61 feat: filter Prowlarr searches by HD categories for movies and series.
d26be96 fix: default unknown resolution tag to an empty string instead of 1080p
58a27d6 release: v1.5.8 — warmup boundary fix, async Wake, Jellyfin support
5793f58 refactor: remove Head Warmup UI field, hardcode warmupFileSize at 64MB
5249ffb fix: eliminate FetchBlock stall at warmupFileSize boundary
af04d1f Restored AI tuner default connections/timeout and reset current limit on playback end, and deferred pump start until first real read while refining logging for pump attachment, seeks, and handle releases.
b931ba8 refactor: adjust native bridge `Wake` behavior to be asynchronous with warmup and synchronous without, and proactively start the native pump during `Open()`.
ed45d64 refactor: update `getHandle` to return `cachedHandle` directly for improved access to its fields.
9d0a620 fix: Prevent writing to closed disk warmup handles by introducing a `closed` atomic flag and adding checks in write operations.
206f501 fix: Prevent panic by adding a nil check for `fileStat` during torrent status processing.
6696807 feat: Add Jellyfin webhook compatibility with JSON body support and event/type normalization
539ef59 Update README for Plex and Jellyfin compatibility
d7ae0db feat: Adjust read-ahead initial size, refine disk warmup coverage and pump activity tracking, and improve pump throttling during the warmup phase.
77311d6 Fixes cross-boundary cache reads, improves LRU promotion and eviction robustness, and adds Llama cache reset on communication delay.
f6663f3 feat: Introduce configurable Torrentio URL and refactor NativeReader pipe handling with atomic pointer.
06b01eb Add DeepWiki badge to README
cfc07e8 Regression from Direct Reader because stability.
0e62dd9 refactor: simplify disk warmup logic and enhance AI tuner robustness by adjusting peer timeout limits and adding response pre-processing.
716ce1f - feat: enhance disk warmup performance.
8c21b8e feat: Increase disk warmup safety margin and cold start look-ahead buffer to 32MB for smoother 4K playback.
9dbc351 feat: introduce `pumpOff` to accurately steer the pump by tracking official player position and prevent hash verification race conditions by disabling buffer pooling.
a0c48de feat: Implement cold start look-ahead for pump, increase initial read-ahead
c692f85 fix: use io.ReadFull instead of Read to ensure full buffer reads.
b3abd68 refactor: replace pipe with direct torrstor.Reader in NativeReader and FetchBlock
4822cf6 Refactor: Simplify AI tuner's metric history format, update AI prompt with examples, and remove time-based cooldown for discovery boost.
68ed1db feat: Introduce AI-driven discovery boost for re-announcing to trackers and refactor `Torrent.Announce` to avoid deadlocks.
d4e1eda refactor: Improve disk warmup tail range tracking with a high watermark, reduce initial read-ahead default, and remove pump V-lookahead logic.
253caf6 feat: Add ExecStartPre command to create the /tmp/llama-slots directory before server startup.
f5f3d97 Remove unused `lastKnownTotalSpeed` and `totalSpeedRaw` variables and their associated speed accumulation logic.
6575c2c refactor: conditionally reset Llama cache only when `lastActiveHash` is not empty
acdeedf feat: Implement Llama KV cache reset and configure the AI server for memory locking and slot management.
75ea629 Fix multiple race conditions, memory leaks, and context handling issues across native pump, disk warmup, and caching.
ab7b3e3 refactor: clarify file size display in AI prompt and remove unused contextStr variable.
57cc6ce feat: Enhance torrent completion detection by considering active and queued piece hashes and calling `updateComplete` when active hashes are zero.
3d23394 feat: Limit concurrent piece hashing at the client level to `runtime.NumCPU()` and improve `requestingPeer` safety.
6e925c2 feat: Include file size in AI tuning prompt and optimizer log message.
34c27e4 feat: upgrade AI model to Qwen3-0.6B, add idle guard, remove dead config field
406a02d feat: remove the 60-second upper limit for PeerTimeout in the AI tuner
d818d46 docs: add dynamic optimization example; refactor: simplify AI tuner by removing player state logic from the LLM prompt.
39f9da8 feat: Implement adaptive AI tuning with player state and swarm peer context, refine multi-stream handling, and add connected seeders to torrent statistics.
2903dec feat: Upgrade AI setup script to download and test the Qwen3-0.6B-Instruct model instead of Qwen2.5-0.5B.
d5c4cd0 Merge pull request #10 from drewwells/dockerLazyFuse
06c7143 Merge pull request #11 from drewwells/monitorVPN
ae3f448 use fromkeys
88fe4c7 health monitor vpn widget
318ae29 detect stale fuse layers and recover
e130829 perf: Decrease minimum ConnectionsLimit and PeerTimeout values from 15 to 10 in the AI tuner.
5d8a51b Merge pull request #9 from drewwells/dockerizedHealthMonitor
c073b10 feat: Switch AI Pilot LLM to Qwen3-0.6B, update model details, and add prompt design documentation.
67351e8 Containerize health monitor and harden FUSE shutdown
2ff1612 docs: Update AI Pilot to use Llama-3.2-1B model, refine operational logic, and enhance fail-safe mechanisms.
6f23aed refactor: use configurable defaults for AI-Pilot multi-stream protection and extend metrics history retention.
4caaaff feat: update AI server to Llama-3.2-1B with increased context and memory, and refactor AI tuner with multi-stream safety defaults, auto-disabling on LLM un
55b7570 fix: Prevent negative refCount and spurious grace period timers by only decrementing for handles that acquired a pump slot.
fd9192f style: increase margin-bottom for section titles in settings.html
d7a7ac9 fix: Use `staticCorruptionCount` for watchdog activation to prevent indefinite dangling corruption states.
4d0400f fix: Start watchdog immediately on first corruption (`count >= 1`) to prevent indefinitely dangling pending states.
7a1641e style: Add styling for time input fields and update selected day pill background and border colors to use accent variable.
1849eb4 docs: document MKV Creation Scheduler, replace cron references
fa14254 feat: Implement built-in scheduler for daily and interval sync tasks, including watchlist synchronization.
e2e3d10 docs: update README and prowlarr-adapter for v1.5.5 — config.json-based Prowlarr settings, plex.tv_library_id, Control Panel sections
df9afac feat: Load Prowlarr client configuration from `config.json` instead of using hardcoded values.
7138ce3 feat:  Prowlarr indexer integration in UI and add a dedicated Plex TV library ID, updating configuration and UI accordingly.
c649421 Merge pull request #6 from drewwells/dockerRestart
5adbaa3 fixes for docker restart breaking the fuse mount
7b7dd2a feat: Implement delayed adaptive shield activation based on consecutive piece corruptions to prevent micro-stutters.
30b309d Update AdaptiveShield clean streak detection threshold from 60 seconds to 30 seconds.
936579f fix: Prevent native bridge read hangs on context timeout and enhance smbd watchdog with emergency unblock.
c3fa0b5 Fix duplicate caution message in README
cbe471f docs: Add a caution to the README explaining potential slow playback during initial Plex library scans due to BitTorrent engine congestion.
17d2bc3 fix: Lowercase the info hash before returning it.
4f368b3 chore: Disable Prowlarr client by default.
2d612c1 typo fixed
8d5bb22 feat: Implement Go-based TV episode registry management with real-time removal and a self-healing watchdog, eliminating the Python rehydration limit.
5e5a74a feat: increase `MAX_REHYDRATE_PER_RUN` limit to 100 in `gostorm-tv-sync.py`.
644671e feat: Improve service portability and logging by using relative paths and centralizing logs, and enhance the install script with auto-detection for user, group, and install directory.
8dd251e fix: standardize path resolution for `RootPath`, `_state_dir`, and `_log_dir` to use the configuration directory directly.
45f829a refactor: remove environment variable loading for Prowlarr client configuration.
60deee3 feat: Implement environment variable configuration and an `ENABLED` flag for the Prowlarr client.
46ed136 Change timeout setting in Prowlarr adapter documentation
17da091 feat: configure Prowlarr client with specific API key, base URL, and use a `requests.Session` with a 30-second timeout.
65d1aeb fix: map 'warn' to 'warning' log level and increase Prowlarr search timeout to 10 seconds.
55978bc fix: Reduce Prowlarr search request timeout from 45 seconds to 2 seconds.
fc26be7 docs: Add Prowlarr Integration Resilience section to README.
9fa0f23 feat: Implement resolution extraction and fake Torrentio name in stream objects to trigger GoStream quality filters.
cbf973b refactor:  direct TorrServer API interaction and virtual .mkv file creation logic.
58e8145 refactor: Refine logging verbosity and messages for movie processing and existing library checks, and update the startup version string.
a650ab0 feat: Improve logging clarity by differentiating between no streams found and all streams filtered out from Prowlarr results.
f842634 fix: Update Prowlarr client to use 'query' for IMDB ID searches and 'tvsearch' for series, including all indexers.
629c9e2 Fix typo in Prowlarr adapter configuration
3c596ba Updated Prowlarr adapter configuration details
f22b2f3 feat: document Prowlarr integration with strict fallback logic and generalize client API configuration.
5d51de6 feat(sync): Integrate Prowlarr Adapter as primary torrent source`
8537744 Performance: Replace hash/fnv with xxhash/v2 for Zero-Alloc Hot Paths
77c457e Gemini GitHub Actions
5f17090 Update README
b65ff7f ...
974ccb3 Update README
41ad621 README update
16e8e3c Moved Health Monitor screenshot
25f3368 Center align health monitor image in README
bd1d6de Update image syntax in README for Health Monitor
a8a2a1f image update
f9a9820 README update
8fb32e5 README update
6a2f163 README update
7485415 Remove badges from README
6f02813 Remove MIT License badge from README
c94248a Change license from MIT to GPL v2.0
6a96695 Replaced MIT license with GPL-2.0
7589108 fix: Add Plex and TMDB API key configuration with corresponding UI fields in control panel.
57a4d99 refactoring
8891332 Update AI GoStream Pilot description in documentation
ce2a8bc Enhance AI GoStream Pilot overview and author notes
c58b6a3 Adjust timeout values in AI optimization logs
fc7a203 Update AI real-time adjustments section
9c01d2f Add link to AI GoStream Pilot documentation
c31b913 Update AI Pilot to AI GoStream Pilot in documentation
d2dcaf2 feat: enhance AI Pilot with torrent context change detection, stricter prompt sanitization, and optimized inference frequency. Documentation update.
a2da633 fix: Enhance AI response parsing by adding newline stop sequence, detecting malformed JSON, and improving unit sanitization.
c6aa7e4 feat: Implement AI tuner rolling averages, high-resolution stats, and hysteresis for improved stability and responsiveness.
40ca51d chore: Improve AI-Pilot optimizer logs by including the full context string and clarifying the download speed metric.
768480a feat: Implement dynamic torrent connection limits, prioritizing AI tuner's output over hardcoded values and allowing retrieval of current limits.
0d577b6 refactor: remove outdated and redundant comments from disk warmup cache logic.
bc952df refactor: Clarify total download speed variable name in AI tuner prompt context.
259b189 feat: Update AI tuner prompt for better reactivity and refine AI response parsing logic.
721b5c2 feat: enhance AI tuner with active torrent filtering, detailed context for prompts, and updated logging.
6ae02e1 refactor: improve `if` statement readability and update AI tuning prompt objective.
99f8cd5 AI Pilot for GoStream (Unified Engine) Doc Update
0786b2f Change AI Server RAM limit from 800MB to 500MB
78e41b7 feat: Introduce AI Stream Pilot for dynamic BitTorrent parameter tuning on Raspberry Pi, including tuner logic, setup scripts, and service integration. LLM qwen2.5-0.5b installed locally. setup_pi.sh for auto setup on Pi 4.
c3dafbe update buttons
d64932b update buttons
ed3ec13 README update
7f4d8f5 README update
0530eec README update
5045d49 README update
e5fecf2 README update
9171ca4 Update GitHub Sponsors username in FUNDING.yml
8230e93 IMG update
0cc1930 IMG update
efa2589 feat: Implement lazy pump start anchored to player position on first read to prevent resume stutter.
591ab3a README update
4e92946 README update
270df34 feat: add readahead coupling metric box and rename FUSE Cache Budget
c436a50 fix: remove preloadCache JS references to fix settings load
efce8ed fix: V309 shield reset on seek, remove Preload Cache from control panel
7c69b7d feat: Add `torrstor.ResetShield()` and update log message when interrupting pump for seek.
c76bfa3 feat: increase default cache size and read-ahead percentage, and improve AdaptiveShield reliability by preventing false positives from evicted pieces and extending its clean streak detection window.
ebdceec feat: implement V304 explicit session-based IP banning for corrupt peers, optimize native pump with V-lookahead, and refine Adaptive Shield logic.
885ecca fix: add existence check for torrent cache directory before removal
3e1f9e1 chore: Simplify button labels for saving configuration and applying settings in the settings page.
bfd367f feat: add GitHub star badge to settings page header
cdd5f14 README update
2ce5219 README update
6927b17 README update
1a1000a refactor: adjust state and log directory paths to align with GoStream's new layout and update script description.
14ae48d fix: Reduce default `MOVIE_RECHECK_CACHE_TTL_HOURS` from 720 to 48.
d45204f Merge pull request #3 from drewwells/syncCaching
fce1b39 Fix: Uncheck shield input when its controlling responsive setting is disabled.
9ce128c feat: Conditionally disable the Adaptive Shield setting based on the Smart Responsive Mode state, providing visual feedback.
e94b258 docs: Clarify Smart Responsive Mode description to explain FAST mode and Adaptive Shield integration.
75167ee feat: Add Adaptive Shield  to automatically activate STRICT mode on piece corruption, with corresponding backend setting and UI toggle.
59a87be fix(sync): restore premium-lang bypass and avoid premature recheck mark
5579472 fix: remove xyz from GoStorm Sync integration log message
baca733 feat: Remove user-configurable disk warmup quota from installer and set a default value. Warmup disabled by default
74d6e7f README update
3988aec feat: enable multi-architecture Docker builds and profile-guided optimization for the GoStream binary.
e71f1c3 Merge pull request #2 from drewwells/dockerfile
741f330 fix: Conditionally extend Shared Pump grace period to 90s for confirmed playback to prevent micro-stutter during Plex CIFS reconnects.
4609e09 perf(sync): add long recheck and add-fail cooldown caches
73c1a39 perf(sync): cut repeated API work in movie sync
24f439d dockerfile for fuse mounts
7d8d198 fix: reduce AdaptiveShield recovery time and eliminate delayed torrent save to prevent playback contention.
76587ba feat: implement IMDB ID bootstrapping and caching from webhooks into playback state for improved matching.
13a5991 fix: Make `nativeBridge.Wake` always synchronous to prevent `FetchBlock` timeouts during resume.
91b59d6 fix: Prevent `raCache` ring buffer exhaustion by removing emergency overdrive throttling and eliminate playback stutters by skipping torrent reads for disk-warmed data.
68163f8 README update
f92cdea README update
a1827ca README update
b3b306c Merge pull request #1 from drewwells/bugfixes
cd56a3a fix(fuse): guard tail warmup when disk cache is disabled
8fdc3be README update
c9f4b32 README update
ec28b3b README update
b6f2474 install: warn about direct BitTorrent traffic when NAT-PMP is disabled
4d88411 feat: Add `playerOff` to `NativePumpState` to persist and inherit player read position, preventing false V286 backward-seek interrupts.
9bfadf3 fix: V701 — prevent false V286 seek interrupt on FUSE handle reopen
2283c6a perf: reduce BoltDB I/O — NoSync + save debounce (30s)
5a6fe17 feat: V304 Adaptive Shield — raCache responsiveOnly fix
65f28f6 docs: mention auto-upgrade of existing MKVs to higher quality
0c0a25c fix: add wg-quick After= only when NAT-PMP is enabled
4b73625 fix: start gostream after wg-quick@wg0.service
3722fac fix: set PreloadCache default to 0%
1c497e5 fix: Plex section auto-discovery (env var), add tv_library_id support
b8ecfb6 docs: fix warmup cache path — configured via Control Panel, not STATE/
5489f97 docs: remove binary size from Single binary label
a24bc29 docs: remove stats table from intro
4911193 docs: update peak throughput to 400+ Mbps
fb75261 docs: remove emoji from What's included list
d01bcff docs: remove emoji from all section headings
02f8732 docs: add installer screenshot to Quick Install section
af4bc80 docs: add Plex library screenshot to Sync Scripts section
48fca70 docs: clarify Key File Locations — runtime vs build dirs
73c6439 docs: update control_1.png screenshot
2ce9213 config: set default ReaderReadAHead to 50% and CacheSize to 64MB
1ecc2bd docs: update health_monitor_1.png screenshot
3924f60 docs: credit TorrServer Matrix 1.37 + anacrolix/torrent v1.55 forks with streaming patches
59b78c3 docs: add full feature list — auto-discovery movies/TV, NAT-PMP, blocklist, watchlist, control panel
91c9de1 docs: update tagline — Forget Real-Debrid
6d17f05 docs: rewrite intro to lead with FUSE concept; demote Watchlist to optional automation
62c1d1e fix: replace nslookup with getent hosts in service DNS check (nslookup not in Debian 13)
2ed594d fix: enable FUSE user_allow_other in /etc/fuse.conf during install
3a0cc9e ux: auto-recommend disk warmup quota based on available space (20% of avail, 3-50 GB cap) with description
bd48b37 config: reduce default disk warmup quota from 32GB to 3GB
cfedd6a fix: set GOTMPDIR to home dir to avoid /tmp tmpfs OOM during linking on Pi 4
6ef1b37 fix: add swap creation and -p 2 to prevent OOM during Go compilation on Pi 4
d7b747d fix: add missing state packages; fix STATE/ gitignore pattern (case-insensitive macOS matched state/)
219aab9 fix: restore correct module import paths (gostream/internal/... not github.com/MrRobotoGit/gostream/...)
52eb201 fix: restore module name to 'gostream' (imports use gostream/internal/...)
72e88f1 fix: add libfuse3-dev and gcc to auto-install deps (required for CGO compilation)
dc62fad fix: rename module to github.com/MrRobotoGit/gostream to resolve standard library lookup errors
2495bfd fix: make install.sh executable
c8398bd Fix build: lower go.mod requirement to 1.24.0
9acdd03 Fix build: add GOTOOLCHAIN=local to prevent toolchain download path collision
88b6bbf Fix installer: GoStream + GoStorm naming, auto-detect target platform
e0a86ec Remove control_2 and control_3 screenshots
da9de00 Remove health_monitor_2 screenshot from README
f79c519 Restore screenshots to be visible directly (no collapsible sections)
7c024f6 Revamp README.md — premium visual identity, badges, collapsible sections, emoji navigation
3a2e779 docs: remove all version references (Vxxx) from README
74b9983 docs: rewrite opening description — emphasize custom FUSE, Native Bridge, seek architecture
12d3de7 docs: replace GoStream_Src with gostream throughout README
b49fb37 fix: detect architecture and OS dynamically for Go install and compilation
d678773 feat: auto-compile GoStream binary during installation
a8bc21b fix: suppress pip script-location warning (uvicorn/fastapi used as modules)
b02c0e6 fix: add --break-system-packages for pip3 on Debian 12 / Raspberry Pi OS Bookworm
085f3de fix: deploy files to INSTALL_DIR and resolve config.json.example path
ce563e5 fix: auto-install missing system dependencies via apt in install.sh
3dd5eb3 docs: add screenshots for Control Panel and Health Monitor
4eb09c0 docs: add GoStream Control Panel and Health Monitor sections to README
0182f9d docs: comprehensive README, install.sh with cron/sudoers/webhook guide
32934dd update .gitignore
5eecd20 Initial public release: GoStream v1.4.4 (Gillian)

## Upstream main vs Fork HEAD — Stat

 .github/workflows/docker-publish.yml              |   58 +-
 README.md                                         |  124 +--
 config.json.example                               |   12 +-
 docker-windows/templates/Dockerfile.jellyfin.tmpl |    2 +-
 docker-windows/templates/Dockerfile.plex.tmpl     |    2 +-
 docker/Dockerfile                                 |   11 +
 go.mod                                            |    6 +-
 go.sum                                            |    4 +-
 hermes/SKILL.md                                   |  962 -------------------
 hermes/skill.md                                   | 1022 +++++++++++++++++++++
 install.sh                                        |    2 +-
 internal/catalog/mediaserver/client.go            |    3 +-
 internal/config/config.go                         |   30 +-
 internal/gostorm/torr/utils/torrent.go            |  104 +--
 internal/library/gostorm.go                       |   29 -
 internal/library/handler.go                       |   99 --
 internal/library/locks.go                         |   73 --
 internal/library/magnet.go                        |   42 -
 internal/library/manager.go                       |  932 -------------------
 internal/library/naming.go                        |  170 ----
 internal/library/refresh.go                       |   71 --
 internal/prowlarr/client.go                       |   58 +-
 internal/subprovider/config.go                    |   98 ++
 internal/subprovider/engine.go                    |  355 +++++++
 internal/subprovider/matcher.go                   |  220 +++++
 internal/subprovider/matcher_test.go              |  105 +++
 internal/subprovider/opensubtitles.go             |  187 ++++
 internal/subprovider/provider.go                  |  102 ++
 internal/subprovider/subdb.go                     |  169 ++++
 internal/subprovider/writer.go                    |  102 ++
 internal/syncer/engines/gostorm.go                |   44 +-
 internal/syncer/engines/language.go               |    1 +
 internal/syncer/engines/movie_go.go               |   99 +-
 internal/syncer/engines/tv_go.go                  |   61 +-
 internal/updater/updater.go                       |    7 +-
 main.go                                           |  305 +++++-
 settings.html                                     |  204 +++-
 37 files changed, 3095 insertions(+), 2780 deletions(-)

## Rota C: Arquivos únicos (subprovider/ e subtitle-related)

### internal/subprovider/ (Rota C core)
internal/subprovider/config.go
internal/subprovider/engine.go
internal/subprovider/matcher.go
internal/subprovider/matcher_test.go
internal/subprovider/opensubtitles.go
internal/subprovider/provider.go
internal/subprovider/subdb.go
internal/subprovider/writer.go

### Arquivos com Rota C/ subtitle em diff
internal/subprovider/config.go
internal/subprovider/engine.go
internal/subprovider/matcher.go
internal/subprovider/matcher_test.go
internal/subprovider/opensubtitles.go
internal/subprovider/provider.go
internal/subprovider/subdb.go
internal/subprovider/writer.go

### Diff stat separado por área

Rota C (internal/subprovider/):
 internal/subprovider/config.go        |  98 ++++++++++
 internal/subprovider/engine.go        | 355 ++++++++++++++++++++++++++++++++++
 internal/subprovider/matcher.go       | 220 +++++++++++++++++++++
 internal/subprovider/matcher_test.go  | 105 ++++++++++
 internal/subprovider/opensubtitles.go | 187 ++++++++++++++++++
 internal/subprovider/provider.go      | 102 ++++++++++
 internal/subprovider/subdb.go         | 169 ++++++++++++++++
 internal/subprovider/writer.go        | 102 ++++++++++
 8 files changed, 1338 insertions(+)

Parser EBML (internal/anacrolix-torrent):
