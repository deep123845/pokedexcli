package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationInfo(locationName string) (LocationInfo, error) {
	url := baseURL + "location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationInfo := LocationInfo{}
		err := json.Unmarshal(val, &locationInfo)
		if err != nil {
			return LocationInfo{}, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationInfo{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationInfo{}, err
	}

	locationInfo := LocationInfo{}
	err = json.Unmarshal(dat, &locationInfo)
	if err != nil {
		return LocationInfo{}, err
	}

	c.cache.Add(url, dat)
	return locationInfo, nil
}
