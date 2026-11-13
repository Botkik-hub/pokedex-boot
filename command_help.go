package main

import "fmt"

func commandHelp(conf *config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("Too much arguments, require 0")
	}
	line := `Welcome to the Pokedex!
Usage:

`

	for _, cmd := range conf.commands {
		line += cmd.name + ": " + cmd.description + "\n"
	}
	fmt.Print(line)
	return nil
}
