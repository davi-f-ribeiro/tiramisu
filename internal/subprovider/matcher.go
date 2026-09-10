package subprovider

import (
	"strings"
	"unicode"
)

// FuzzyMatch computes a simple fuzzy match score between 0-100.
// Compares release strings (from provider metadata) with the torrent name.
func FuzzyMatch(a, b string) int {
	if a == "" || b == "" {
		return 0
	}

	// Normalize both strings
	na := normalizeRelease(a)
	nb := normalizeRelease(b)

	if na == nb {
		return 100
	}

	// Length difference penalty
	diff := len(na) - len(nb)
	if diff < 0 {
		diff = -diff
	}
	if diff > 10 {
		return 20 // too different in length
	}

	// Character overlap
	overlap := 0
	for i, r := range na {
		if i < len(nb) && r == []rune(nb)[i] {
			overlap++
		}
	}

	total := len(na)
	if len(nb) > total {
		total = len(nb)
	}

	score := (overlap * 100) / total
	if score < 20 {
		return 20
	}
	return score
}

// normalizeRelease strips common noise from release strings
// (YIFY, Netflix, WEB-DL, etc.) to improve matching.
func normalizeRelease(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)

	var out []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			out = append(out, r)
		}
	}

	return string(out)
}

// SelectBestSubtitle picks the best subtitle based on:
// 1. Language preference (preferred lang first)
// 2. Fuzzy match against torrent/release name
func SelectBestSubtitle(subs []Subtitle, torrentName string, preferredLangs []LanguageTag) int {
	if len(subs) == 0 {
		return -1
	}

	// Create a lookup for preferred languages
	preferredSet := make(map[LanguageTag]int)
	for i, lang := range preferredLangs {
		preferredSet[lang] = i
	}

	var bestIdx int
	bestScore := -1

	for i := range subs {
		s := &subs[i]
		score := 0

		// Priority 1: exact preferred language match
		if idx, ok := preferredSet[s.Language]; ok {
			score += 1000 - idx*10 // 1000 base, minus penalty for lower preference
		} else if s.Language == "por" || s.Language == "pt" || s.Language == "pt-br" {
			score += 800 // fallback portuguese
		} else if s.Language == "multi" || s.Language == "multilang" {
			score += 600
		} else if s.Language == "eng" || s.Language == "en" {
			score += 400
		} else {
			score += 100
		}

		// Priority 2: fuzzy match against release info
		releaseScore := FuzzyMatch(s.ReleaseInfo, torrentName)
		score += releaseScore

		// Priority 3: provider's own score (as tiebreaker)
		score += s.Score / 10

		if score > bestScore {
			bestScore = score
			bestIdx = i
		}
	}

	return bestIdx
}

// BestSubtitleByLanguage picks the best subtitle for the requested language.
func BestSubtitleByLanguage(subs []Subtitle, lang LanguageTag) *Subtitle {
	var best *Subtitle
	bestScore := -1

	for i := range subs {
		s := &subs[i]
		var score int

		if s.Language == lang {
			score = 100
		} else if s.Language == "por" || s.Language == "pt" || s.Language == "pt-br" {
			score = 90
		} else if s.Language == "multi" || s.Language == "multilang" {
			score = 80
		} else if s.Language == "eng" || s.Language == "en" {
			score = 60
		} else {
			score = 20
		}

		score += s.Score / 2

		if score > bestScore {
			bestScore = score
			best = s
		}
	}

	return best
}

// BestSubtitleByTorrentName picks the subtitle with highest fuzzy match
// against the torrent/release name.
func BestSubtitleByTorrentName(subs []Subtitle, torrentName string) *Subtitle {
	var best *Subtitle
	bestScore := -1

	for i := range subs {
		s := &subs[i]
		score := FuzzyMatch(s.ReleaseInfo, torrentName)
		if score > bestScore {
			bestScore = score
			best = s
		}
	}

	return best
}
