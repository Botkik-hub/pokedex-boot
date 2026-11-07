package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationArea struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations(page int) (locationArea, error) {
	apiURL := "https://pokeapi.co/api/v2/location-area"

	if page < 0 {
		page = 0
	}

	pageSize := 20
	pageOffset := page * pageSize

	apiURL += fmt.Sprintf("?limit=%d&offset=%d", pageSize, pageOffset)

	res, err := http.Get(apiURL)
	if err != nil {
		return locationArea{}, err
	}

	defer res.Body.Close()

	var locations locationArea

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&locations); err != nil {
		return locationArea{}, err
	}

	return locations, nil
}
