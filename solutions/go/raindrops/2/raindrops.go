// Package raindrops is imprements a function that convert a number
// into raindrop sounds.
package raindrops

import "strconv"

// Convert take an integer and return a string base on this rule:
// If the number divisible by 3, add "Pling"
// If the number divisible by 5, add "Plang"
// If the number divisible by 5, add "Plong"
// If the number is not divisible by 3,5,7 return a number as a string
func Convert(number int) string {
	result := ""

	if number%3 == 0 {
		result = "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		result = strconv.Itoa(number)
	}

	return result

}
