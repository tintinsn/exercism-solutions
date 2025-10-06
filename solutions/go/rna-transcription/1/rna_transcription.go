package strand

func ToRNA(dna string) string {
	dnaMap := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	var rna string
	for _, nucleotides := range dna {
		rna += string(dnaMap[nucleotides])
	}
	return rna
}
