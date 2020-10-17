// Package strand implements DNA strand to RNA convert function
package strand

import (
	"strings"
)

// ToRNA given a DNA strand, returns its RNA complement (per RNA transcription).
//
//Both DNA and RNA strands are a sequence of nucleotides.
//
//The four nucleotides found in DNA are adenine (**A**), cytosine (**C**),
//guanine (**G**) and thymine (**T**).
//
//The four nucleotides found in RNA are adenine (**A**), cytosine (**C**),
//guanine (**G**) and uracil (**U**).
//
//Given a DNA strand, its transcribed RNA strand is formed by replacing
//each nucleotide with its complement:
//
//* `G` -> `C`
//* `C` -> `G`
//* `T` -> `A`
//* `A` -> `U`
func ToRNA(dna string) string {
	strings.NewReplacer()
	var rna strings.Builder
	for _, d := range dna {
		r := DNAToRNA[d]
		rna.WriteRune(r)
	}
	return rna.String()
}

var DNAToRNA = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}
