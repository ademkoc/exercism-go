package anagram

import (
	"slices"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var result []string
	for _, candidate := range candidates {
		if strings.EqualFold(subject, candidate) {
			continue
		}

		normalize := func(input string) string {
			chars := strings.Split(strings.ToLower(input), "")
			slices.Sort(chars)
			return strings.Join(chars, "")
		}

		if normalize(subject) == normalize(candidate) {
			result = append(result, candidate)
		}
	}
	return result
}
