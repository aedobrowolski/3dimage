package main

import "math/rand"

// shuffleText shuffles INPUT returning the first GOAL characters
func shuffleText(input string, goal int) string {
	if len(input) == 0 {
		panic("Invalid argument")
	}

	// Extend the input if needed
	for len(input) < goal {
		input = input + input
	}

	// Shuffle using the Fisher-Yates algorithm
	a := []byte(input)
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}

	return string(a[:goal])
}
