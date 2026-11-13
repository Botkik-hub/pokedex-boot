package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/botkik-hub/pokedex-boot/pokeapi"
)

func commandCatch(conf *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Wrong amount of arguments, required: pokemon name")
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := pokeapi.GetPokemonData(name, conf.cache)
	if err != nil || pokemon.Name == "" {
		return err
	}

	k := 50.0
	minC := 0.10
	maxC := 0.99

	baseExp := pokemon.BaseExperience
	chanceToCatch := max(minC, min(maxC, k/float64(baseExp)))

	//fmt.Printf("Chance to catch %s is %f\n", pokemon.Name, chanceToCatch)

	randNum := rand.Float64()
	if randNum < chanceToCatch {
		conf.pokemons[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
