package grains

import (
	"fmt"
)

func Square(number int) (uint64, error) {
	if number > 64 || number < 1 {
		return 0, fmt.Errorf("number must be more than 1 and less than 65")
	}

	return 1 << (number - 1), nil
}

func Total() uint64 {
	return 1<<64 - 1
}
