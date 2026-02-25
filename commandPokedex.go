package main

import "fmt"

func CommandPokedex(c *config, _ []string) error {
	if len(c.pokedex) == 0 {
		fmt.Println("Your Pokedex is empty")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range c.pokedex {
		fmt.Println(" -", name)
	}

	return nil
}
