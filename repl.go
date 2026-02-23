package main

import (
	"strings"

	"github.com/deep123845/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	Previous *string
	Next     *string
	Cache    *pokecache.Cache
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
			description: "Displays the next 20 locations from the world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations from the world",
			callback:    commandMapb,
		},
	}
}
