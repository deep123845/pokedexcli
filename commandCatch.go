package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func CommandCatch(c *config, args []string) error {
	if numArgs := len(args); numArgs != 1 {
		return fmt.Errorf("Expected 1 argument, received %v", numArgs)
	}

	args[0] = strings.ToLower(args[0])

	output, err := c.pokeapiClient.GetPokemonInfo(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	catchRequirment := math.Min(float64(output.BaseExperience)/float64(300), 0.90)
	catchHit := rand.Float64()
	if catchHit < catchRequirment {
		fmt.Println(args[0], "escaped!")
		return nil
	}

	fmt.Println(args[0], "was caught!")
	c.pokedex[args[0]] = output
	return nil
}
