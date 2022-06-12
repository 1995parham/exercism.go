package strand

import "strings"

func ToRNA(dna string) string {
	rna := &strings.Builder{}
	t := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	for _, n := range dna {
		rna.WriteRune(t[n])
	}

	return rna.String()
}
