package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)

		if cleanedInput == nil {
			fmt.Println("invalid input")
		}

		firstWord := cleanedInput[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}
}
