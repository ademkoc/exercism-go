package isbn

import (
	"strconv"
	"unicode"
)

func IsValidISBN(isbn string) bool {

	if isbn == "" || len(isbn) > 10 || len(isbn) < 10 {
		return false
	}

	var result int
	multiplier := 10
	for _, value := range isbn {
		if value != 'X' && !unicode.IsDigit(value) {
			continue
		}
		digit := 0
		if value == 'X' {
			digit = 10
		} else {
			digit, _ = strconv.Atoi(string(value))
		}
		result += digit * multiplier
		multiplier = multiplier - 1
	}

	return result%11 == 0
}
