package wordcount

import (
	"regexp"
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {

	var str strings.Builder
	for _, letter := range phrase {
		if unicode.IsLetter(letter) || unicode.IsDigit(letter) {
			str.WriteRune(unicode.ToLower(letter))
		} else if letter == '\'' {
			str.WriteRune('\'')
		} else {
			str.WriteRune(' ')
		}
	}

	re := regexp.MustCompile(`\s+`)
	trimed := re.ReplaceAllString(str.String(), ` `)

	freq := make(Frequency)

	for _, word := range strings.Split(trimed, " ") {
		word = strings.Trim(word, "'")
		if word == "" {
			continue
		}
		freq[word]++
	}

	return freq
}
