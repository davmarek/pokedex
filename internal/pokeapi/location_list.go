package pokeapi

import "encoding/json"

type RespLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationAreas(pageUrl *string) (RespLocations, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	res, err := c.httpClient.Get(url)
	if err != nil {
		return RespLocations{}, err
	}
	defer res.Body.Close()

	var result RespLocations
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&result); err != nil {
		return RespLocations{}, err
	}
	return result, nil
}
