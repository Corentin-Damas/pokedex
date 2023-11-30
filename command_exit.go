package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("Now Quiting Pokedex ...")
	os.Exit(0)
	return nil
}
