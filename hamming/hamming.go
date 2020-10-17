// Package hamming implements Hamming Distance calculation
package hamming

import "errors"

var errNotEqual = errors.New("DNA strands not equal length")

// Distance calculates the Hamming Distance between two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errNotEqual
	}
	var n int
	for i := range a {
		if a[i] != b[i] {
			n++
		}
	}
	return n, nil
}
