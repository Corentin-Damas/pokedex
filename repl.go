package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanedText := cleanInput(text)

		if len(cleanedText) == 0 {
			continue
		}
		commandName := cleanedText[0]
		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Explore next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Comeback previous 20 locations",
			callback:    commandMapB,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
