package grains

import (
	"errors"
	"math"
)

func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("the square number provided is not valid")
	}
	if input == 1 {
		return 1, nil
	}

	return uint64(math.Pow(2, float64(input)-1)), nil
}

func Total() (result uint64) {
	for i := 1; i <= 64; i++ {
		value, _ := Square(i)
		result += value
	}

	return result
}
