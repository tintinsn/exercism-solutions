package isbn

import (
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}

	result := 0

	for i, num := range isbn {
		if num == 'X' {
			if i != 9 {
				return false
			}
			result += 10
		} else if num >= '0' && num <= '9' {
			result += int(num-'0') * (10 - i)
		} else {
			return false
		}

	}
	return result%11 == 0
}
