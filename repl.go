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

	fmt.Print(
		"List available locations with 'map' and 'mapb',\n",
		"then 'explore' one that catches your eye for Pokemon\n",
		"to 'catch' and add to your 'pokedex'!\n\n",
		"You may 'inspect' any Pokemon you've caught.\n\n",
		"Use 'help' to see available commands.\n",
	)

	for {
		fmt.Print("\npokedex > ")
		scanner.Scan()
		input := scanner.Text()

		cleanedInput := cleanInput(input)
		if cleanedInput == nil {
			continue
		}

		userCmd := cleanedInput[0]
		cmdArg := ""
		if len(cleanedInput) > 1 {
			cmdArg = cleanedInput[1]
		}

		if cmd, ok := availableCmds[userCmd]; ok {
			err := cmd.callback(&config, cmdArg)
			if err != nil {
				fmt.Printf("Error calling command: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
