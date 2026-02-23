package main

import (
	"fmt"

	"github.com/deep123845/pokedexcli/internal/pokeapi"
)

func commandMap(c *config) error {
	output, err := pokeapi.GetLocationAreas(c.Next)
	if err != nil {
		return err
	}

	c.Next = output.Next
	c.Previous = output.Previous

	for _, location := range output.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(c *config) error {
	output, err := pokeapi.GetLocationAreas(c.Previous)
	if err != nil {
		return err
	}

	c.Next = output.Next
	c.Previous = output.Previous

	for _, location := range output.Results {
		fmt.Println(location.Name)
	}

	return nil
}
