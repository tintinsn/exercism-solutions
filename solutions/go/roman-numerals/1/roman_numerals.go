package romannumerals

import "errors"

type romanPair struct {
	symbol string
	value  int
}

func ToRomanNumeral(input int) (string, error) {
	if input > 3999 || input < 1 {
		return "", errors.New("Input can not more than 3999 and less than 1")
	}
	roman := []romanPair{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}

	result := ""

	for input > 0 {
		for _, pair := range roman {
			if pair.value <= input {
				result += pair.symbol
				input -= pair.value
				break
			}
		}
	}

	return result, nil
}
