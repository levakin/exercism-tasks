// Package isogram implements function to determine if a word or phrase is an isogram.
package isogram

import "unicode"

// IsIsogram determines if a word or phrase is an isogram
func IsIsogram(word string) bool {
	letters := make(map[rune]bool)
	for _, l := range word {
		if !unicode.IsLetter(l) {
			continue
		}
		l = unicode.ToLower(l)
		if letters[l] {
			return false
		}
		letters[l] = true
	}
	return true
}
