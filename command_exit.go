package main

import (
	"fmt"
	"os"
)

func commandExit(conf *config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("Too much arguments, require 0")
	}
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
