// Package isogram is Exercism.io exercise
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram — Determine if a word or phrase is an isogram.
func IsIsogram(input string) bool {
	lowInput := strings.ToLower(input)
	for _, v := range lowInput {
		if !unicode.IsLetter(v) {
			continue
		}
		if strings.Count(lowInput, string(v)) > 1 {
			return false
		}
	}
	return true
}
