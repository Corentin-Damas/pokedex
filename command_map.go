package main

import (
	"fmt"
)

func commandMap(cfg *config , args ...string) error {

	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas available: ")
	for _, area := range res.Results {
		fmt.Printf(" - %s \n", area.Name)
	}

	cfg.nextLocationAreaURL = res.Next
	cfg.prevLocationAreaURL = res.Previous
	return nil
}
