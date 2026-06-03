package main

import (
	"fmt"
)

func commandMapf(cfg *config) error {
	response, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationsUrl)
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

	response, err := cfg.pokeapiClient.GetLocationAreas(cfg.previousLocationsUrl)
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
