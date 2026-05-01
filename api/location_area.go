package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (*LocationAreaResponse, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("error with response")
	}

	var result LocationAreaResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
