package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/Auxiguilar/go-pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config, string) error
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
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon you have caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View your pokedex",
			callback:    commandPokedex,
		},
	}
}

func commandExit(cfg *pokeapi.Config, s string) error {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)
	return err
}

func commandHelp(cfg *pokeapi.Config, s string) error {
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

func commandMap(cfg *pokeapi.Config, s string) error {
	if cfg.UrlNext == nil {
		_, err := fmt.Println("You are already on the last page")
		return err
	}

	data, err := cfg.GetAreaData(*cfg.UrlNext)
	if err != nil {
		return fmt.Errorf("Getting area data: %w", err)
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

func commandMapB(cfg *pokeapi.Config, s string) error {
	if cfg.UrlPrevious == nil {
		_, err := fmt.Println("You are already on the first page")
		return err
	}

	data, err := cfg.GetAreaData(*cfg.UrlPrevious)
	if err != nil {
		return fmt.Errorf("Getting area data: %w", err)
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

func commandExplore(cfg *pokeapi.Config, areaName string) error {
	// access pokemon: locationData.PokemonEncounters[i].Name (string)
	data, err := cfg.GetLocationData(areaName)
	if err != nil {
		return fmt.Errorf("Getting location data: %w", err)
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range data.PokemonEncounters {
		name := pokemon.Pokemon.Name
		_, err := fmt.Printf("- %s\n", name)
		if err != nil {
			return err
		}
	}

	return nil
}

func commandCatch(cfg *pokeapi.Config, pokeName string) error {
	// access pokemon: pokemonData.Name
	data, err := cfg.GetPokemonData(pokeName)
	if err != nil {
		return fmt.Errorf("Getting pokemon data: %w", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", data.Name)

	// failure
	if rand.Intn(1000) < data.BaseExperience {
		fmt.Printf("%s escaped!\n", data.Name)
		return nil
	}

	// success
	fmt.Printf("%s was caught!\n", data.Name)
	cfg.Pokemon[data.Name] = data

	return nil
}

func commandInspect(cfg *pokeapi.Config, pokeName string) error {
	pokemon, ok := cfg.Pokemon[pokeName]
	if !ok {
		_, err := fmt.Println("You have not caught that pokemon yet")
		return err
	}

	// no, I guess I won't be checking err...
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *pokeapi.Config, s string) error {
	for _, pokemon := range cfg.Pokemon {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	fmt.Println("...")

	return nil
}
