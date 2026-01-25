package dndcharacter

import (
	"math"
	"math/rand/v2"
	"slices"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	diceValues := make([]int, 4)
	for i := 0; i < 4; i++ {
		diceValues[i] = rand.IntN(6) + 1
	}

	slices.Sort(diceValues)
	slices.Reverse(diceValues)

	var abilityValue int

	for i := 0; i < 3; i++ {
		abilityValue += diceValues[i]
	}
	return abilityValue
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	c := Ability()
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: c,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(c),
	}
}
