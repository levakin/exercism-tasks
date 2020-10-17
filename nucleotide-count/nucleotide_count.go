// Package dna implements histogram of DNA generation
package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

func newHistogram() *Histogram {
	return &Histogram{
		Adenine:  0,
		Cytosine: 0,
		Guanine:  0,
		Thymine:  0,
	}
}

// DNA is a list of nucleotides.
type DNA string

var errInvNuc = errors.New("dna contains an invalid nucleotide")

const (
	Adenine  = 'A'
	Cytosine = 'C'
	Guanine  = 'G'
	Thymine  = 'T'
)

func isValidNucleotide(n rune) bool {
	return n == Adenine || n == Cytosine || n == Guanine || n == Thymine
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := *newHistogram()
	for _, n := range d {
		if !isValidNucleotide(n) {
			return nil, errInvNuc
		}
		h[n] += 1
	}
	return h, nil
}
