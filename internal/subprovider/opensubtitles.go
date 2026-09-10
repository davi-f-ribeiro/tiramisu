package subprovider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// OSProvider implements Provider for the OpenSubtitles REST API.
type OSProvider struct {
	apiKey   string
	baseURL  string
	user     string
	password string
	token    string
	client   *http.Client
}

// NewOSProvider creates an OpenSubtitles REST API provider.
func NewOSProvider(apiKey, baseURL, user, password string) *OSProvider {
	if baseURL == "" {
		baseURL = "https://api.opensubtitles.com"
	}
	return &OSProvider{
		apiKey:   apiKey,
		baseURL:  strings.TrimRight(baseURL, "/"),
		user:     user,
		password: password,
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (p *OSProvider) Name() string { return "opensubtitles" }

// CanMatch: true when we have an IMDB ID.
func (p *OSProvider) CanMatch(videoHash, imdbID, torrentName string) bool {
	return imdbID != ""
}

// Search finds subtitles for the given IMDB ID.
func (p *OSProvider) Search(videoHash, imdbID, torrentName string, lang LanguageTag, limit int) ([]Subtitle, error) {
	url := fmt.Sprintf(
		"%s/api/v1/subtitles?imdb_id=%s&languages=%s",
		p.baseURL, imdbID, lang,
	)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Api-Key", p.apiKey)

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("os search GET: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		return nil, fmt.Errorf("os search: unauthorized (check Api-Key)")
	}
	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, fmt.Errorf("os search: rate limited")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("os search: HTTP %d", resp.StatusCode)
	}

	var searchResp struct {
		Data []struct {
			Type  string `json:"type"`
			ID    int64  `json:"id"`
			Attrs struct {
				Language string   `json:"language"`
				Release  string   `json:"release"`
				Files    []struct {
					ID       int64  `json:"id"`
					FileName string   `json:"file_name"`
				} `json:"files"`
			} `json:"attributes"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		return nil, fmt.Errorf("os search parse: %w", err)
	}

	var results []Subtitle
	for _, item := range searchResp.Data {
		if len(item.Attrs.Files) == 0 {
			continue
		}

		langCode := item.Attrs.Language
		if langCode == "" {
			continue
		}

		score := 100
		if langCode == string(lang) {
			score = 100
		} else if langCode == "por" || langCode == "pt" {
			score = 90
		} else if langCode == "multi" {
			score = 80
		}

		fileName := item.Attrs.Files[0].FileName

		results = append(results, Subtitle{
			ID:           fmt.Sprintf("%d", item.Attrs.Files[0].ID),
			ProviderName: "opensubtitles",
			Language:     LanguageTag(langCode),
			Format:       "srt",
			Filename:     fileName,
			ReleaseInfo:  item.Attrs.Release,
			Score:        score,
		})
	}

	return results, nil
}

// Download obtains a subtitle from OpenSubtitles.
func (p *OSProvider) Download(sub Subtitle) ([]byte, error) {
	// Extract file_id from sub.ID
	var fileID int64
	fmt.Sscanf(sub.ID, "%d", &fileID)
	if fileID == 0 {
		return nil, fmt.Errorf("os download: invalid file_id %q", sub.ID)
	}

	downloadURL := p.baseURL + "/api/v1/download"

	downloadPayload, _ := json.Marshal(map[string]int64{
		"file_id": fileID,
	})

	req, _ := http.NewRequest("POST", downloadURL, bytes.NewReader(downloadPayload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", p.apiKey)
	if p.token != "" {
		req.Header.Set("Authorization", "Bearer "+p.token)
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("os download POST: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		return nil, fmt.Errorf("os download: unauthorized")
	}
	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, fmt.Errorf("os download: rate limited")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("os download: HTTP %d", resp.StatusCode)
	}

	var dlResp struct {
		Link string `json:"link"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&dlResp); err != nil {
		return nil, fmt.Errorf("os download parse: %w", err)
	}
	if dlResp.Link == "" {
		return nil, fmt.Errorf("os download: empty link")
	}

	// GET the signed link
	resp2, err := p.client.Get(dlResp.Link)
	if err != nil {
		return nil, fmt.Errorf("os download GET: %w", err)
	}
	defer resp2.Body.Close()

	if resp2.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("os download: GET HTTP %d", resp2.StatusCode)
	}

	return io.ReadAll(resp2.Body)
}
