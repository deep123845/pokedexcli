package pokeapi

import (
	"encoding/json"
	"net/http"
)

func GetLocationAreas(url *string) (locationAreas, error) {
	pageURL := baseURL + "location-area/"
	if url != nil {
		pageURL = *url
	}

	res, err := http.Get(pageURL)
	if err != nil {
		return locationAreas{}, err
	}
	defer res.Body.Close()

	var output locationAreas
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&output)
	if err != nil {
		return locationAreas{}, err
	}

	return output, nil
}
