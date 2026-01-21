package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Auxiguilar/go-pokedex/internal/pokeapi"
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

func startRepl() {
	availableCmds := getCommands()
	scanner := bufio.NewScanner(os.Stdin)

	config := pokeapi.NewConfig()

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		cleanedInput := cleanInput(input)
		if cleanedInput == nil {
			continue
		}

		userCmd := cleanedInput[0]
		if cmd, ok := availableCmds[userCmd]; ok {
			err := cmd.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
