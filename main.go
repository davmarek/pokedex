package main

import (
	"bufio"
	"fmt"
	"os"
)

var commandRegistry map[string]cliCommand

func main() {
	commandRegistry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		input := cleanInput(text)
		if len(input) > 0 {
			if c, ok := commandRegistry[input[0]]; ok {
				err := c.callback()
				if err != nil {
					fmt.Println(fmt.Errorf("%w", err))
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
