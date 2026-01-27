package isogram

import "unicode"

type blank struct{}

func IsIsogram(word string) bool {
	lookup := make(map[rune]blank)

	for _, chr := range word {
		if !unicode.IsLetter(chr) {
			continue
		}
		ltr := unicode.ToLower(chr)
		if _, exists := lookup[ltr]; exists {
			return false
		}
		lookup[ltr] = blank{}
	}
	return true
}
