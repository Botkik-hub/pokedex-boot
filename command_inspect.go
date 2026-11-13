package main

import (
	"fmt"

	"github.com/botkik-hub/pokedex-boot/pokeapi"
)

func printPokemon(p *pokeapi.PokemonData) {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)

	fmt.Println("Stats:")
	for _, s := range p.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
}

func commandInspect(conf *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Wrong amount of arguments, required: name")
	}
	name := args[0]

	pokemon, found := conf.pokemons[name]
	if !found {
		fmt.Println("you have not caught this pokemon")
		return nil
	}
	printPokemon(&pokemon)
	return nil
}
