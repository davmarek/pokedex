package main

import (
	"fmt"

	"github.com/davmarek/pokedex/api"
)

func commandMapf(cfg *config) error {
	url := ""
	if cfg.nextLocationsUrl != nil {
		url = *cfg.nextLocationsUrl
	}

	response, err := api.GetLocationAreas(url)
	if err != nil {
		return err
	}

	cfg.nextLocationsUrl = response.Next
	cfg.previousLocationsUrl = response.Previous

	for _, location := range response.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previousLocationsUrl == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	response, err := api.GetLocationAreas(*cfg.previousLocationsUrl)
	if err != nil {
		return err
	}
	for _, location := range response.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}
