package pangram

import "strings"

func IsPangram(input string) bool {
	charMap := map[rune]bool{}
	for _, char := range strings.ToLower(input) {
		if 'a' <= char && char <= 'z' {
			_, exists := charMap[char]
			if !exists {
				charMap[char] = true
			}
		}
	}
	return len(charMap) == 26
}
