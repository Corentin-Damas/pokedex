package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Pokedex is a command line where you type your the pokemon you want to know about")
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s \n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
