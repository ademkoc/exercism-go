package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fd FodderCalculator, cows int) (float64, error) {
	amount, err := fd.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	factor, err := fd.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return amount * factor / float64(cows), nil
}

func ValidateInputAndDivideFood(fd FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(fd, cows)
	}
	return 0, errors.New("invalid number of cows")
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "there are no negative cows",
		}
	} else if cows == 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "no cows don't need food",
		}
	}
	return nil
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}
