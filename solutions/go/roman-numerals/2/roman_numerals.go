package romannumerals

import (
	"errors"
	"strings"
)

// Values in descending order, including subtraction cases
var values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var numerals = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", errors.New("invalid")
	}

	var str strings.Builder
	for i, value := range values {
		for input >= value {
			str.WriteString(numerals[i])
			input -= value
		}
	}

	return str.String(), nil
}
