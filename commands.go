package main

import (
	"fmt"
	"os"

	"github.com/Auxiguilar/go-pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of areas",
			callback:    commandMapB,
		},
	}
}

func commandExit(cfg *pokeapi.Config) error {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)
	return err
}

func commandHelp(cfg *pokeapi.Config) error {
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

func commandMap(cfg *pokeapi.Config) error {
	if cfg.UrlNext == nil {
		_, err := fmt.Println("You are already on the last page")
		return err
	}

	data, err := cfg.GetAreaData(*cfg.UrlNext)
	if err != nil {
		return err
	}

	for _, result := range data.Results {
		_, err := fmt.Println(result.Name)
		if err != nil {
			return err
		}
	}

	cfg.UrlNext = data.Next
	cfg.UrlPrevious = data.Previous

	return nil
}

func commandMapB(cfg *pokeapi.Config) error {
	if cfg.UrlPrevious == nil {
		_, err := fmt.Println("You are already on the first page")
		return err
	}

	data, err := cfg.GetAreaData(*cfg.UrlPrevious)
	if err != nil {
		return err
	}

	for _, result := range data.Results {
		_, err := fmt.Println(result.Name)
		if err != nil {
			return err
		}
	}

	cfg.UrlNext = data.Next
	cfg.UrlPrevious = data.Previous

	return nil
}
