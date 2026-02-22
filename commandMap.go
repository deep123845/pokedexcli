package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.Next != "" {
		url = c.Next
	}

	return getMapAreas(c, url)
}

func commandMapb(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.Previous != "" {
		url = c.Previous
	}

	return getMapAreas(c, url)
}

func getMapAreas(c *config, url string) error {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		return err
	}
	defer res.Body.Close()

	var output locationAreas
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&output)
	if err != nil {
		fmt.Printf("%s", err)
		return err
	}

	c.Next = output.Next
	c.Previous = output.Previous

	for _, l := range output.Results {
		fmt.Printf("%s\n", l.Name)
	}

	return nil
}
