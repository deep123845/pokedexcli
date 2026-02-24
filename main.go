package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/deep123845/pokedexcli/internal/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := getCommands()

	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	config := &config{
		pokeapiClient: pokeClient,
	}

	for {
		fmt.Println()
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()
		cleanedText := cleanInput(text)
		if len(cleanedText) < 1 {
			continue
		}

		word := cleanedText[0]
		command, ok := commands[word]
		if !ok {
			fmt.Print("Unknown Command")
		}

		err := command.callback(config)
		if err != nil {
			fmt.Print(err)
		}
	}
}
