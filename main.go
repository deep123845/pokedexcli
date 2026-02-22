package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := getCommands()

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

		err := command.callback()
		if err != nil {
			fmt.Print(err)
		}
	}
}
