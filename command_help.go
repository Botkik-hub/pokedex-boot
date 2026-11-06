package main

import "fmt"

func commandHelp(conf *config) error {
	line := `Welcome to the Pokedex!
Usage:

`

	for _, cmd := range conf.commands {
		line += cmd.name + ": " + cmd.description + "\n"
	}
	fmt.Print(line)
	return nil
}
