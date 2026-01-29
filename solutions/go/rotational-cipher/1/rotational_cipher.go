package rotationalcipher

import (
	"slices"
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	alphabet := []rune("abcdefghijklmnopqrstuvwxyz")

	var result strings.Builder
	for _, char := range plain {
		if !unicode.IsLetter(char) {
			result.WriteRune(char)
			continue
		}

		upper := unicode.IsUpper(char)
		baseIndex := slices.Index(alphabet, char)
		if upper {
			baseIndex = slices.Index(alphabet, unicode.ToLower(char))
		}

		shiftedIndex := (baseIndex + shiftKey) % len(alphabet)
		shiftedChar := alphabet[shiftedIndex]
		if upper {
			shiftedChar = unicode.ToUpper(alphabet[shiftedIndex])
		}
		result.WriteRune(shiftedChar)
	}
	return result.String()
}
