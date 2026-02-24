package main

import (
	"fmt"
)

func commandMap(c *config) error {
	url := c.Next

	output, err := c.pokeapiClient.GetLocations(url)
	if err != nil {
		return err
	}

	for _, location := range output.Results {
		fmt.Println(location.Name)
	}

	c.Next = output.Next
	c.Previous = output.Previous

	return nil
}

func commandMapb(c *config) error {
	url := c.Previous

	output, err := c.pokeapiClient.GetLocations(url)
	if err != nil {
		return err
	}

	for _, location := range output.Results {
		fmt.Println(location.Name)
	}

	c.Next = output.Next
	c.Previous = output.Previous

	return nil
}
