package dndcharacter

import (
	"math"
	"math/rand"
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
	modifierPoint := (float64(score) - 10) / 2
	return int(math.Floor(modifierPoint))

}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	min := rand.Intn(6) + 1
	sum := 0

	for i := 1; i <= 3; i++ {
		randNum := rand.Intn(6) + 1

		if randNum > min {
			sum += randNum
		} else {
			sum += min
			min = randNum
		}
	}

	return sum
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constitutionPoint := Ability()
	hitpoints := Modifier(constitutionPoint) + 10
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitutionPoint,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    hitpoints,
	}
}
