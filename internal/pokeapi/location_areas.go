package pokeapi

import (
	"encoding/json"
	"net/http"
)

func GetLocationAreas(url *string) (LocationAreas, string, error) {
	pageURL := baseURL + "location-area/"
	if url != nil {
		pageURL = *url
	}

	res, err := http.Get(pageURL)
	if err != nil {
		return LocationAreas{}, pageURL, err
	}
	defer res.Body.Close()

	var output LocationAreas
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&output)
	if err != nil {
		return LocationAreas{}, pageURL, err
	}

	return output, pageURL, nil
}
