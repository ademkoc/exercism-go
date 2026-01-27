package grains

import (
	"errors"
	"math"
)

const TotalSquareCount = 64

func Square(number int) (uint64, error) {
	if number <= 0 || number > TotalSquareCount {
		return 0, errors.New("invalid number")
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	n := math.Pow(2, TotalSquareCount-1)
	return uint64((n * (n + 1) * (2*n + 1)) / 6)
}
