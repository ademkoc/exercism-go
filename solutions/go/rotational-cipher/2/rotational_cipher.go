package rotationalcipher

import (
	"strings"
)

func RotationalCipher(plain string, shiftKey int) string {

	var result strings.Builder

	for _, r := range plain {
		base := 0
		switch {
		case r >= 'A' && r <= 'Z':
			base = int('A')
		case r >= 'a' && r <= 'z':
			base = int('a')
		}
		if base != 0 {
			r = rune(((int(r) - base + shiftKey) % 26) + base)
		}
		result.WriteRune(r)
	}

	return result.String()
}
