package romannumerals

import (
	"errors"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", errors.New("invalid")
	}

	// Values in descending order, including subtraction cases
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	numerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result strings.Builder

	for i := 0; i < len(values); i++ {
		// Repeat the numeral while the input is >= value
		for input >= values[i] {
			result.WriteString(numerals[i])
			input -= values[i]
		}
	}

	return result.String(), nil
}
