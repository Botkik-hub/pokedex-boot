package main

import (
	"fmt"
)

func commandPokedex(conf *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, val := range conf.pokemons {
		fmt.Printf("\t-%s\n", val.Name)
	}
	return nil
}
