// Package isogram is Exercism.io exercise
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram — Determine if a word or phrase is an isogram.
func IsIsogram(input string) bool {
	lowInput := strings.ToLower(input)
	readed := make(map[rune]bool)
	for _, v := range lowInput {
		if !unicode.IsLetter(v) {
			continue
		}
		if readed[v] {
			return false
		}
		readed[v] = true
	}
	return true
}
