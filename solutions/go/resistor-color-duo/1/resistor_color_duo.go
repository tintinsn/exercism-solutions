package resistorcolorduo

import "strconv"

func Value(colors []string) int {
	colorsCode := map[string]string{
		"black":  "0",
		"brown":  "1",
		"red":    "2",
		"orange": "3",
		"yellow": "4",
		"green":  "5",
		"blue":   "6",
		"violet": "7",
		"grey":   "8",
		"white":  "9",
	}
	result := ""

	for i, color := range colors {
		if i < 2 {
			val, ok := colorsCode[color]
			if ok {
				result += val
			}
		}
	}
	num, err := strconv.Atoi(result)
	if err != nil {
		return 0
	}
	return num
}
