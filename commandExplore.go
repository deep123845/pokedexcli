package main

import "fmt"

func CommandExplore(c *config, args []string) error {
	if numArgs := len(args); numArgs != 1 {
		return fmt.Errorf("Expected 1 argument, received %v", numArgs)
	}

	output, err := c.pokeapiClient.GetLocationInfo(args[0])
	if err != nil {
		return err
	}

	fmt.Println("Exploring " + args[0])
	fmt.Println("Found Pokemon:")
	for _, encounter := range output.PokemonEncounters {
		fmt.Println("- " + encounter.Pokemon.Name)
	}

	return nil
}
