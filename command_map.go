package main

import (
	"fmt"

	"github.com/botkik-hub/pokedex-boot/pokeapi"
)

func commandMap(conf *config) error {
	locations, err := pokeapi.GetLocations(conf.lastLocationPage)
	if err != nil {
		return err
	}
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	conf.lastLocationPage++
	return nil
}

func commandMapb(conf *config) error {
	nextPage := conf.lastLocationPage - 1
	if nextPage < 0 {
		nextPage = 0
	}
	locations, err := pokeapi.GetLocations(nextPage)
	if err != nil {
		return err
	}
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	conf.lastLocationPage = nextPage
	return nil
}
