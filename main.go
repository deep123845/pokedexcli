package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/deep123845/pokedexcli/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := getCommands()

	cache := pokecache.NewCache(time.Minute)
	config := config{
		Cache: &cache,
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

		err := command.callback(&config)
		if err != nil {
			fmt.Print(err)
		}
	}
}
