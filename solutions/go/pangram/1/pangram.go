package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	normalized := strings.ToLower(input)
	stats := make(map[rune]struct{})

	for _, letter := range normalized {

		isLetter := unicode.IsLetter(letter)
		if !isLetter {
			continue
		}

		stats[letter] = struct{}{}
	}

	return len(stats) == 26

}
