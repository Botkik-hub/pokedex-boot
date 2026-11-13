package main

import (
	"time"

	"github.com/botkik-hub/pokedex-boot/pokeapi"
	"github.com/botkik-hub/pokedex-boot/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	commands         map[string]cliCommand
	lastLocationPage int
	cache            *pokecache.Cache
	pokemons         map[string]pokeapi.PokemonData
}

func initConfig() config {
	return config{
		commands: map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "Displays a help message",
				callback:    commandHelp,
			},
			"map": {
				name:        "map",
				description: "Shows next page of locations",
				callback:    commandMap,
			},
			"mapb": {
				name:        "mapb",
				description: "Shows previous page of locations",
				callback:    commandMapb,
			},
			"explore": {
				name:        "explore",
				description: "Shows pokemon names that can be encountred at the selected location. Arguments: Name of the location",
				callback:    commandExplore,
			},
			"catch": {
				name:        "catch",
				description: "Throws a pokeball and tries to catch a pokemon, chance depends on the pokemon level",
				callback:    commandCatch,
			},
			"inspect": {
				name:        "inspect",
				description: "shows stats of the caught pokemon",
				callback:    commandInspect,
			},
		},
		lastLocationPage: 0,
		cache:            pokecache.NewCache(5 * time.Minute),
		pokemons:         map[string]pokeapi.PokemonData{},
	}
}
