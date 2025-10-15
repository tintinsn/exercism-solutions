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
			result += 10
		} else {
			result += int(num-'0') * (len(isbn) - i)
		}
	}
	return result%11 == 0
}
