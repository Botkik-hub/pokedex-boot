package main

import (
	"fmt"

	"github.com/botkik-hub/pokedex-boot/pokeapi"
)

func commandExplore(conf *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Wrong amount of arguments, required: name")
	}
	name := args[0]

	locations, err := pokeapi.GetLocationData(name, conf.cache)
	if err != nil {
		return err
	}
	for _, poc := range locations.PokemonEncounters {
		fmt.Println(poc.Pokemon.Name)
	}
	conf.lastLocationPage++
	return nil
}
