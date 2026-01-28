package isbn

import (
	"strconv"
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	var normalized strings.Builder
	for i, r := range isbn {
		if unicode.IsDigit(r) {
			normalized.WriteRune(r)
		} else if r == 'X' && i == len(isbn)-1 {
			normalized.WriteRune(r)
		} else if unicode.IsLetter(r) {
			return false
		}
	}
	isbn = normalized.String()

	if len(isbn) != 10 {
		return false
	}

	var result int
	multiplier := 10
	for _, value := range isbn {
		digit := 0
		if value == 'X' {
			digit = 10
		} else {
			digit, _ = strconv.Atoi(string(value))
		}
		result += (digit * multiplier)
		multiplier = multiplier - 1
	}

	return result%11 == 0
}
