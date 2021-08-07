// Package isogram provides utilities for handling isogram words
package isogram

import "strings"

// IsIsogram function determines whether a word is isogram or not (case insensitive)
func IsIsogram(word string) bool {
	letters := make(map[rune]bool)
	word = strings.ToLower(word)
	for _, c := range word {
		if c == ' ' || c == '-' {
			continue
		}
		if letters[c] {
			return false
		} else {
			letters[c] = true
		}
	}
	return true
}
