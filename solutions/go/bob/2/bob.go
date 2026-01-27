package bob

import (
	"strings"
	"unicode"
)

func isShout(phrase string) bool {
	if strings.IndexFunc(phrase, unicode.IsLetter) == -1 {
		return false
	}
	return strings.ToUpper(phrase) == phrase
}

func isQuestion(phrase string) bool {
	return strings.HasSuffix(phrase, "?")
}

func isEmpty(s string) bool {
	return s == ""
}

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	question := isQuestion(remark)
	shout := isShout(remark)

	if shout {
		if question {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}

	if question {
		return "Sure."
	}

	return "Whatever."
}
