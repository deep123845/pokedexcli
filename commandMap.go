package main

import (
	"encoding/json"
	"fmt"

	"github.com/deep123845/pokedexcli/internal/pokeapi"
)

func commandMap(c *config) error {
	url := c.Next

	output, ok := getCachedLocationAreas(c, url)
	if !ok {
		var err error
		var newURL string
		output, newURL, err = pokeapi.GetLocationAreas(url)
		if err != nil {
			return err
		}

		err = addCachedLocationAreas(c, newURL, output)
		if err != nil {
			return err
		}
	}

	c.Next = output.Next
	c.Previous = output.Previous

	for _, location := range output.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(c *config) error {
	url := c.Previous

	output, ok := getCachedLocationAreas(c, url)
	if !ok {
		var err error
		var newURL string
		output, newURL, err = pokeapi.GetLocationAreas(url)
		if err != nil {
			return err
		}

		err = addCachedLocationAreas(c, newURL, output)
		if err != nil {
			return err
		}
	}

	c.Next = output.Next
	c.Previous = output.Previous

	for _, location := range output.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func getCachedLocationAreas(c *config, url *string) (pokeapi.LocationAreas, bool) {
	if url == nil {
		return pokeapi.LocationAreas{}, false
	}

	data, ok := c.Cache.Get(*url)
	if !ok {
		return pokeapi.LocationAreas{}, false
	}

	var areas pokeapi.LocationAreas
	err := json.Unmarshal(data, &areas)
	if err != nil {
		return pokeapi.LocationAreas{}, false
	}

	return areas, true
}

func addCachedLocationAreas(c *config, url string, locationAreas pokeapi.LocationAreas) error {
	data, err := json.Marshal(locationAreas)
	if err != nil {
		return err
	}

	c.Cache.Add(url, data)
	return nil
}
