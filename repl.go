package main

import (
	"strings"
)

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	words := strings.Split(text, " ")
	var result []string

	for _, w := range words {
		word := strings.ToLower(strings.TrimSpace(w))
		if word != "" {
			result = append(result, word)
		}
	}

	return result
}
