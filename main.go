package main

import (
	"time"

	pokeapi "github.com/Corentin-Damas/pokedexcli/internal"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)

}

// https://www.youtube.com/watch?v=8yrmAGcCnKg

// test all test files : go test ./...
// go build
// ./pokedexcli
