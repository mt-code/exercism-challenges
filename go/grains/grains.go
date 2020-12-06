package grains

import "errors"

func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, errors.New("the square number provided is not valid")
	}

	return uint64(1) << uint64(input-1), nil
}

func Total() uint64 {
	// uint64(0) sets all the register bits to 0
	// ^ = XOR operator which flips the bits to 1, which equals uint64 max value
	return ^uint64(0)
}
