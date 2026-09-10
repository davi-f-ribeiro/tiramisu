// handleSubtitleResync handles the POST /api/subtitle/resync endpoint.
// Called by Hermes/Aegir voice command: "resync legenda".
// Deletes the cached subtitle and forces a new download.
func handleSubtitleResync(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		VideoPath string `json:"video_path"`
		Language  string `json:"language"` // optional, default: por
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.VideoPath == "" {
		http.Error(w, "video_path is required", http.StatusBadRequest)
		return
	}

	cfg := gc()
	if !cfg.Subtitle.Enabled {
		http.Error(w, "Subtitle provider disabled", http.StatusBadRequest)
		return
	}

	// Determine language
	lang := subprovider.LangPortuguese
	if req.Language != "" {
		switch req.Language {
		case "eng", "en":
			lang = subprovider.LangEnglish
		case "multi", "multilang":
			lang = subprovider.LangMulti
		}
	}

	// Build engine config
	engineCfg := subprovider.EngineConfig{
		FUSEMountPath:         cfg.FuseMountPath,
		PreferredLanguages:    []subprovider.LanguageTag{lang},
		OpenSubtitlesKey:      cfg.Subtitle.APIKey,
		OpenSubtitlesBaseURL:  cfg.Subtitle.BaseURL,
		OpenSubtitlesUser:     cfg.Subtitle.User,
		OpenSubtitlesPass:     cfg.Subtitle.Password,
		MaxResultsPerProvider: cfg.Subtitle.MaxResults,
	}

	// Create engine
	engine, err := subprovider.NewSubtitleEngine(engineCfg)
	if err != nil {
		logger.Printf("[SUB] Resync engine init failed: %v", err)
		http.Error(w, fmt.Sprintf("Engine init failed: %v", err), http.StatusInternalServerError)
		return
	}

	// Delete cached subtitle
	if deleted := engine.DeleteCachedSubtitle(req.VideoPath, lang); deleted {
		logger.Printf("[SUB] Resync: deleted cached subtitle for %s (%s)", req.VideoPath, lang)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"deleted","message":"Cached subtitle removed"}`))
	} else {
		logger.Printf("[SUB] Resync: no cached subtitle found for %s (%s)", req.VideoPath, lang)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"status":"not_found","message":"No cached subtitle found"}`))
	}
}
