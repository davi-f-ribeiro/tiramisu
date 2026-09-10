	// --- Rota C: Subtitle Provider Engine Hook (D4: manual resync supported) ---
	cfg := gc()
	if cfg.Subtitle.Enabled {
		go func() {
			videoPath := exactMatch
			videoHash := exactState.Hash
			imdbID := exactState.ImdbID
			if imdbID == "" {
				imdbID = webhookImdbID // fallback
			}
			torrentName := filepath.Base(videoPath)

			// Extract preferred language (first in list, default: por)
			preferredLang := subprovider.LangPortuguese
			if len(cfg.Subtitle.Preferred) > 0 {
				switch cfg.Subtitle.Preferred[0] {
				case "eng", "en":
					preferredLang = subprovider.LangEnglish
				case "multi", "multilang":
					preferredLang = subprovider.LangMulti
				}
			}

			// Build engine config
			engineCfg := subprovider.EngineConfig{
				FUSEMountPath:         cfg.FuseMountPath,
				PreferredLanguages:    []subprovider.LanguageTag{preferredLang},
				OpenSubtitlesKey:      cfg.Subtitle.APIKey,
				OpenSubtitlesBaseURL:  cfg.Subtitle.BaseURL,
				OpenSubtitlesUser:     cfg.Subtitle.User,
				OpenSubtitlesPass:     cfg.Subtitle.Password,
				MaxResultsPerProvider: cfg.Subtitle.MaxResults,
			}

			// Create engine
			engine, err := subprovider.NewSubtitleEngine(engineCfg)
			if err != nil {
				logger.Printf("[SUB] Engine init failed: %v", err)
				return
			}

			// Run subtitle download (fire-and-forget)
			resultCh := make(chan subprovider.Result, 1)
			engine.Run(subprovider.SubtitleJob{
				VideoPath:     videoPath,
				TorrentName:   torrentName,
				VideoHash:     videoHash,
				ImdbID:        imdbID,
				PreferredLang: preferredLang,
				ResultCh:      resultCh,
			})

			// Read result (non-blocking with timeout)
			select {
			case res := <-resultCh:
				if res.Error != nil {
					logger.Printf("[SUB] Download failed: %v", res.Error)
				} else if res.Content == nil && res.Filename != "" {
					logger.Printf("[SUB] Using cached subtitle: %s", res.Filename)
				} else if res.Content != nil {
					logger.Printf("[SUB] Downloaded via %s, wrote to: %s", res.Provider, res.Filename)
				}
			case <-time.After(35 * time.Second):
				logger.Printf("[SUB] Download timed out")
			}
		}()
	}

		if exactState.Hash != "" {