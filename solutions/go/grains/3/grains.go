package grains

import (
	"errors"
)

const TotalSquareCount = 64

func Square(number int) (uint64, error) {
	if number <= 0 || number > TotalSquareCount {
		return 0, errors.New("invalid number")
	}
	return 1 << (number - 1), nil
}

func Total() uint64 {
	return 1<<64 - 1
}
