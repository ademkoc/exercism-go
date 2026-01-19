package isogram

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func IsIsogram(word string) bool {
	var re = regexp.MustCompile(`[-|\s]`)
	var normalized = re.ReplaceAllString(strings.ToLower(word), "")
	length := utf8.RuneCountInString(normalized)
	letters := strings.Split(normalized, "")
	temp := make(map[string]bool)

	for i := 0; i < length-1; i++ {
		if temp[letters[i]] {
			return false
		}

		if i == 0 && letters[i] == letters[length-1] {
			return false
		}

		if letters[i] == letters[i+1] {
			return false
		}

		temp[letters[i]] = true
	}

	return true
}
