package isbn

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	var sum int
	for i, char := range isbn {
		var value int

		if i == 9 && (char == 'X' || char == 'x') {
			value = 10
		} else if unicode.IsDigit(char) {
			value = int(char - '0')
		} else {
			return false
		}

		weight := 10 - i
		sum += value * weight
	}

	return sum%11 == 0
}
