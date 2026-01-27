package luhn

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Valid(id string) bool {

	input := normalizeInput(id)

	if len(input) < 1 {
		return false
	}

	var sum int

	var j = 0
	for i := len(input) - 1; i >= 0; i-- {
		digit := input[i]
		if j%2 == 0 {
			sum += digit
		} else {
			sum += doubleDigit(digit)
		}
		j++
	}

	return sum%10 == 0
}

func doubleDigit(a int) int {
	b := a * 2

	if b > 9 {
		return b - 9
	}
	return b
}

func normalizeInput(id string) []int {
	trimmedId := strings.ReplaceAll(id, " ", "")

	var normalizedArray []int

	if utf8.RuneCountInString(trimmedId) <= 1 {
		return normalizedArray
	}

	if trimmedId == "0" {
		return normalizedArray
	}

	for _, r := range trimmedId {
		isDigit := unicode.IsDigit(r)

		if !isDigit {
			normalizedArray = nil
			break
		}

		val, err := strconv.Atoi(string(r))
		if err != nil {
			panic("Failed to parse int")
		}
		normalizedArray = append(normalizedArray, val)
	}

	return normalizedArray
}
