package main

import (
	"fmt"
	"lab3/levenshtein"
)

func main() {
	var s1, s2 string
	fmt.Print("Enter first string: ")
	fmt.Scanln(&s1)
	fmt.Print("Enter second string: ")
	fmt.Scanln(&s2)

	distance:= levenshtein.LevenshteinDistance([]rune(s1), []rune(s2))

	fmt.Printf("\nLevenshtein distance: %d\n", distance)
}
