package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (c *Client) ListPokemons(area string) (LocationExploreResponse, error) {
	url := baseURL + fmt.Sprintf("/location-area/%s", area)
	body, err := c.get(url)

	if err != nil {
		return LocationExploreResponse{}, errors.New("Request failed with error: " + err.Error())
	}
	explorationResponse := LocationExploreResponse{}
	err = json.Unmarshal(body, &explorationResponse)

	if err != nil {
		return LocationExploreResponse{}, errors.New("Failed to parse response: " + err.Error())
	}

	return explorationResponse, nil
}
