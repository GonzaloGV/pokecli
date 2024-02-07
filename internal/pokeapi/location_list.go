package pokeapi

import (
	"encoding/json"
)

func (c *Client) ListLocations(pageUrl *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if *pageUrl != "" {
		url = *pageUrl
	}

	body, err := c.get(url)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	locations := LocationAreaResponse{}
	err = json.Unmarshal(body, &locations)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locations, nil
}
