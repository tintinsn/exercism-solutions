// Package isogram check if words are isograms.
package isogram

import (
	"strings"
)

// IsIsogram reports whether word has no repeating letters.
// Spaces and hyphens are allowed to repeat.
func IsIsogram(word string) bool {

	word = strings.ToLower(word)
	letters := make(map[rune]bool)

	for _, letter := range word {
		if letter == '-' || letter == '_' || letter == ' ' {
			continue
		}

		if letters[letter] {
			return false
		}

		letters[letter] = true

	}
	return true
}
