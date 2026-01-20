package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	availableCmds := getCommands()
	scanner := bufio.NewScanner(os.Stdin)

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
			cmd.callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
}
