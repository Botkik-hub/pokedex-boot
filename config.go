package main

import (
	"time"

	"github.com/botkik-hub/pokedex-boot/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands         map[string]cliCommand
	lastLocationPage int
	cache            *pokecache.Cache
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
		},
		lastLocationPage: 0,
		cache:            pokecache.NewCache(5 * time.Minute),
	}
}
