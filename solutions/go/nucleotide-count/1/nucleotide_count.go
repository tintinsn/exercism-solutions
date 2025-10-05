package dna

import "fmt"

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for _, dna := range d {
		_, exists := h[dna]
		if !exists {
			return h, fmt.Errorf("invalid nucleotide")
		}

		h[dna]++
	}
	return h, nil
}
