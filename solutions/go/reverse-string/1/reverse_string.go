package reverse

import "fmt"

func Reverse(input string) string {

	runes := []rune(input)
	fmt.Println(input, len(input), len(runes))

	left := 0
	right := len(runes) - 1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--

	}
	return string(runes)

}
