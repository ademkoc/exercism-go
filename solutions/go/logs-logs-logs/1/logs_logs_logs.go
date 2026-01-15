package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	var service string
loop:
	for i, char := range log {
		fmt.Printf("i: %d", i)
		switch char {
		case '❗':
			service = "recommendation"
			break loop
		case '🔍':
			service = "search"
			break loop
		case '☀':
			service = "weather"
			break loop
		default:
			service = "default"
		}

	}
	return service
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newLog strings.Builder
	for _, char := range log {
		if char == oldRune {
			newLog.WriteRune(newRune)
		} else {
			newLog.WriteRune(char)
		}
	}
	return newLog.String()
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
