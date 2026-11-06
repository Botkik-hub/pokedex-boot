package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)

		firstWord := ""
		if len(cleaned) != 0 {
			firstWord = cleaned[0]
		}

		cmd, exists := conf.commands[firstWord]
		if exists {
			cmd.callback(conf)
		} else {
			fmt.Println("Unknown command")
		}
	}
}
