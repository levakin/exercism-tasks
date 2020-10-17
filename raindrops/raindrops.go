// Package raindrops implements converting a number into a string that contains raindrop sounds
package raindrops

import (
	"strconv"
)

// Convert converts a number into a string that contains raindrop sounds corresponding to certain
//potential factors. A factor is a number that evenly divides into another number, leaving no remainder.
//
//The rules of `raindrops` are that if a given number:
//
//- has 3 as a factor, add 'Pling' to the result.
//- has 5 as a factor, add 'Plang' to the result.
//- has 7 as a factor, add 'Plong' to the result.
//- does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number.
func Convert(n int) string {
	var s string
	if n%3 == 0 {
		s += "Pling"
	}
	if n%5 == 0 {
		s += "Plang"
	}
	if n%7 == 0 {
		s += "Plong"
	}
	if s == "" {
		s = strconv.Itoa(n)
	}
	return s
}
