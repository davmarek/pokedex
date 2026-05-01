package main

import (
	"time"

	"github.com/davmarek/pokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(time.Second * 5)
	startRepl(&config{pokeapiClient: client})
}
