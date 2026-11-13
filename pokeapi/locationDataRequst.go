package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/botkik-hub/pokedex-boot/pokecache"
)

type LocationData struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocationData(locName string, cache *pokecache.Cache) (LocationData, error) {
	apiURL := "https://pokeapi.co/api/v2/location-area/"

	if locName == "" {
		return LocationData{}, fmt.Errorf("Empty argument, need to provide name of the location")
	}
	apiURL += locName

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
			return LocationData{}, err
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationData{}, err
		}
	}

	var locations LocationData
	err := json.Unmarshal(data, &locations)
	if err != nil {
		return LocationData{}, err
	}

	if needRequest && cache != nil {
		cache.Add(apiURL, data)
	}

	return locations, nil
}
