package common

import (
	"html"
	"strings"
)

func SanitizeInput(value string) string {
	trimmed := strings.TrimSpace(value)
	return html.EscapeString(trimmed)
}

func SanitizeStringSlice(values []string) []string {
	result := make([]string, 0, len(values))
	for _, value := range values {
		sanitized := SanitizeInput(value)
		if sanitized != "" {
			result = append(result, sanitized)
		}
	}
	return result
}
