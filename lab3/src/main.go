package main

import (
	"fmt"
	"lab3/levenshtein"
)

func main() {
	var specialReplacementCost, specialInsertionCost int
	var specialReplacementCharacter, specialInsertionCharacter rune

	fmt.Print("Enter special replacement character and it's cost: ")
	fmt.Scanf("%c %d", &specialReplacementCharacter, &specialReplacementCost)

	fmt.Print("Enter special insertion character and it's cost: ")
	fmt.Scanf("%c %d", &specialInsertionCharacter, &specialInsertionCost)

	cfg := levenshtein.Config{
		SpecialReplacementCharacter: specialReplacementCharacter,
		SpecialReplacementCost:      specialReplacementCost,
		SpecialInsertionCharacter:   specialInsertionCharacter,
		SpecialInsertionCost:        specialInsertionCost,
	}

	var s1, s2 string

	fmt.Print("Enter first string: ")
	fmt.Scanln(&s1)

	fmt.Print("Enter second string: ")
	fmt.Scanln(&s2)

	distance := levenshtein.LevenshteinDistance([]rune(s1), []rune(s2), cfg)

	fmt.Printf("\nLevenshtein distance: %d\n", distance)
}
