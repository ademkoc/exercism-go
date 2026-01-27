package bob

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func Hey(remark string) string {

	trimedRemark := strings.TrimSpace(remark)

	if trimedRemark == "" {
		return "Fine. Be that way!"
	}

	lastCharacter := trimedRemark[utf8.RuneCountInString(trimedRemark)-1]
	onlyLetters := strings.TrimFunc(remark, func(r rune) bool {
		if unicode.IsLetter(r) {
			return false
		}
		return true
	})

	if onlyLetters != "" && strings.ToUpper(onlyLetters) == onlyLetters {
		if lastCharacter == '?' {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}

	if lastCharacter == '?' {
		return "Sure."
	}

	return "Whatever."
}
