package main

import "fmt"

func commandHelp(_ *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	commands := getCommands()

	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.name, command.description)
	}

	return nil
}
