package scrabble

import "strings"

func Score(word string) int {
	var sum int
	for _, letter := range word {
		sum += getLetterValue(strings.ToUpper(string(letter)))
	}
	return sum
}

func getLetterValue(letter string) int {
	var value int
	switch letter {
	case "Q", "Z":
		value = 10
	case "J", "X":
		value = 8
	case "K":
		value = 5
	case "F", "H", "V", "W", "Y":
		value = 4
	case "B", "C", "M", "P":
		value = 3
	case "D", "G":
		value = 2
	case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
		value = 1
	default:
		value = 0
	}
	println("value ", value)
	return value
}
