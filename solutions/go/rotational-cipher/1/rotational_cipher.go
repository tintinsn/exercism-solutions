package rotationalcipher

func shiftChar(char rune, shift int) rune {
	if char >= 'a' && char <= 'z' {
		position := char - 'a'
		newPosition := (position + rune(shift)) % 26
		char = 'a' + newPosition
	} else if char >= 'A' && char <= 'Z' {
		position := char - 'A'
		newPosition := (position + rune(shift)) % 26
		char = 'A' + newPosition
	}
	return char
}

func RotationalCipher(plain string, shiftKey int) string {
	var result string
	for _, char := range plain {
		result += string(shiftChar(char, shiftKey))
	}
	return result
}
