// Package collatzconjecture is implements a function  that
// Find number of steps it takes to reach 1
package collatzconjecture

import "errors"

// CollatzConjecture take a positive interger and return a number of steps
// Its takes ti reach 1 base on rules:
// If it's even, divide by 2
// If it's odd, multiplt by 3 and add 1
func CollatzConjecture(n int) (int, error) {
	num := n
	i := 0
	if num <= 0 {
		return 0, errors.New("input should be positive number")
	}
	for num > 1 {
		if num%2 == 0 {
			num = num / 2
			i++
			continue
		}
		if num%2 == 1 {
			num = (num * 3) + 1
			i++
			continue
		}

	}

	return i, nil
}
