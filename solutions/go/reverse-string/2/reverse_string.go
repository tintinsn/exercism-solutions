package reverse

func Reverse(input string) string {

	runes := []rune(input)

	left := 0
	right := len(runes) - 1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--

	}
	return string(runes)

}
