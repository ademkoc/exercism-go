package resistorcolorduo

import (
	"strconv"
	"strings"
)

var colorMap = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	str := strings.Builder{}
	for i, color := range colors {
		if i >= 2 {
			break
		}
		value := colorMap[color]
		str.WriteString(strconv.Itoa(value))
	}

	result, _ := strconv.Atoi(str.String())

	return result
}
