package main

import (
	"fmt"

	"github.com/botkik-hub/pokedex-boot/pokeapi"
)

func commandMap(conf *config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("Too much arguments, require 0")
	}
	locations, err := pokeapi.GetLocations(conf.lastLocationPage, conf.cache)
	if err != nil {
		return err
	}
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	conf.lastLocationPage++
	return nil
}

func commandMapb(conf *config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("Too much arguments, require 0")
	}
	nextPage := conf.lastLocationPage - 1
	if nextPage < 0 {
		nextPage = 0
	}
	locations, err := pokeapi.GetLocations(nextPage, conf.cache)
	if err != nil {
		return err
	}
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	conf.lastLocationPage = nextPage
	return nil
}
