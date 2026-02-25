package main

import "fmt"

func CommandInspect(c *config, args []string) error {
	if numArgs := len(args); numArgs != 1 {
		return fmt.Errorf("Expected 1 argument, received %v", numArgs)
	}

	pokemon, ok := c.pokedex[args[0]]
	if !ok {
		return fmt.Errorf("you have not caught the pokemon %s", args[0])
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Println(" -", pokeType.Type.Name)
	}

	return nil
}
