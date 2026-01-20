package main

import (
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	var commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Get help",
			callback:    commandHelp,
		},
	}

	return commands
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")

	availableCmds := getCommands()

	for _, cmd := range availableCmds {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

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
