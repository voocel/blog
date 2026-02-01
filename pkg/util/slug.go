package util

import (
	"regexp"
	"strings"
	"unicode"
)

var slugRegex = regexp.MustCompile(`^[a-z0-9\p{Han}-]+$`)

// GenerateSlug generates a URL-friendly slug from a title.
// It preserves Chinese characters and converts to lowercase.
func GenerateSlug(title string) string {
	s := strings.TrimSpace(title)
	s = strings.ToLower(s)

	var result []rune
	prevDash := false

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			result = append(result, r)
			prevDash = false
		} else if r == '-' || r == ' ' || r == '_' || r == '/' || r == '\\' {
			if !prevDash && len(result) > 0 {
				result = append(result, '-')
				prevDash = true
			}
		}
		// Skip other special characters
	}

	return strings.Trim(string(result), "-")
}

// IsValidSlug validates slug format.
// Valid slugs contain only lowercase letters, numbers, Chinese characters, and hyphens.
func IsValidSlug(slug string) bool {
	if slug == "" || len(slug) > 255 {
		return false
	}
	return slugRegex.MatchString(slug)
}
