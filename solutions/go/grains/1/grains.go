package grains

import (
	"fmt"
	"math"
)

func Square(number int) (uint64, error) {
	if number > 64 || number < 1 {
		return 0, fmt.Errorf("number must be more than 1 and less than 65")
	}

	grain := uint64(math.Pow(2, float64(number)-1))
	return grain, nil
}

func Total() uint64 {
	var total uint64
	for i := 1; i < 65; i++ {
		gain, _ := Square(i)
		total += gain
	}
	return total
}
