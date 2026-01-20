package strand

import "strings"

var NucleotideMap = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	rna := strings.Builder{}
	for _, nucleotide := range dna {
		nucleotide, exists := NucleotideMap[nucleotide]
		if !exists {
			continue
		}
		rna.WriteRune(nucleotide)
	}
	return rna.String()
}
