package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if len(id) <= 1 {
		return false
	}

	var digits []int

	for _, digit := range id {
		if !unicode.IsDigit(digit) {
			return false
		}
		digits = append(digits, int(digit-'0'))
	}

	sum := 0
	double := false
	for i := len(digits) - 1; i >= 0; i-- {
		if double {
			if digits[i]*2 > 9 {
				sum += digits[i]*2 - 9
			} else {
				sum += digits[i] * 2
			}
		} else {
			sum += digits[i]
		}
		double = !double
	}
	return sum%10 == 0
}
