// Package scrabble implements score calculation in Scrabble
package scrabble

import "strings"

var letterScore = map[rune]int{
	'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
	'D': 2, 'G': 2,
	'B': 3, 'C': 3, 'M': 3, 'P': 3,
	'F': 4, 'H': 4, 'Y': 4, 'W': 4, 'V': 4,
	'K': 5,
	'J': 8, 'X': 8,
	'Q': 10, 'Z': 10,
}

// Score given a word, computes the Scrabble score for that word.
func Score(w string) int {
	var s int
	for _, l := range strings.ToUpper(w) {
		s += letterScore[l]
	}
	return s
}
