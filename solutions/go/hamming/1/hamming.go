// Package hamming provides a function to calculate the hamming distance
// Between two DNA strands.
package hamming

import "errors"

// Distance return the number of difference positions (Hamming distance)
// Between two DNA strands a and b. 
// If a and b not the same length, it returns an error.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("a and b must be the same length")
	}

	hamming := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hamming++
		}
	}

	return hamming, nil

}
