// Package diffsquares provides a functions to calculate the difference
// Between the square of the sum and the sum of the squares 
// Of the first N natural numbers.
package diffsquares

import (
	"math"
)

// SquareOfSum returns the square of the sum of the first n natural numbers.
func SquareOfSum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))

}

// SumOfSquares returns the sum of the squares of the first n natural numbers.
func SumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += int(math.Pow(float64(i), 2))
	}
	return sum
}

// Difference returns the difference between the square of the sum
// and the sum of the squares of the first n natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
