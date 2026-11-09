package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/botkik-hub/pokedex-boot/pokecache"
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

func GetLocations(page int, cache *pokecache.Cache) (locationArea, error) {
	apiURL := "https://pokeapi.co/api/v2/location-area"

	if page < 0 {
		page = 0
	}

	pageSize := 20
	pageOffset := page * pageSize

	apiURL += fmt.Sprintf("?limit=%d&offset=%d", pageSize, pageOffset)

	needRequest := false
	var data []byte

	if cache == nil {
		needRequest = true
	} else {
		var found bool
		data, found = cache.Get(apiURL)
		if !found {
			needRequest = true
		} else {
			fmt.Println("Cached result hit instead of response")
		}
	}

	if needRequest {
		res, err := http.Get(apiURL)
		if err != nil {
			return locationArea{}, err
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return locationArea{}, err
		}
	}

	var locations locationArea
	err := json.Unmarshal(data, &locations)
	if err != nil {
		return locationArea{}, err
	}

	if needRequest && cache != nil {
		cache.Add(apiURL, data)
	}

	return locations, nil
}
