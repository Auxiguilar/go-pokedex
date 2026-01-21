package main

import (
	"fmt"
	"os"

	"github.com/Auxiguilar/go-pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// contains the next and previous urls for map command
type config struct {
	Next     *string
	Previous *string
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
		"map": {
			name:        "map",
			description: "Map the world!",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "!dlrow eht paM",
			callback:    commandMapB,
		},
	}

	return commands
}

func commandExit(c *config) error {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)
	return err
}

func commandHelp(c *config) error {
	_, err := fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	if err != nil {
		return err
	}

	availableCmds := getCommands()

	for _, cmd := range availableCmds {
		_, err := fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		if err != nil {
			return err
		}
	}

	return nil
}

// maps the next 20 locations
func commandMap(c *config) error {
	if c.Next == nil {
		_, err := fmt.Println("You are already on the last page")
		return err
	}

	// is it dereferencing??
	data, err := pokeapi.GetAreaData(*c.Next)
	if err != nil {
		return err
	}

	for _, result := range data.Results {
		_, err := fmt.Println(result.Name)
		if err != nil {
			return err
		}
	}

	c.Next = data.Next
	c.Previous = data.Previous

	return nil
}

// maps the previous 20 locations
func commandMapB(c *config) error {
	if c.Previous == nil {
		_, err := fmt.Println("You are already on the first page")
		return err
	}

	data, err := pokeapi.GetAreaData(*c.Previous)
	if err != nil {
		return err
	}

	for _, result := range data.Results {
		_, err := fmt.Println(result.Name)
		if err != nil {
			return err
		}
	}

	c.Next = data.Next
	c.Previous = data.Previous

	return nil
}
