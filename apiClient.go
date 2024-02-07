package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type locationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type pokeapiClient struct {
	http.Client
}

func createClient() pokeapiClient {
	return pokeapiClient{}
}

func extractLocationNames(locations locationAreaResponse) []string {
	eLocations := make([]string, 0)
	for _, location := range locations.Results {
		eLocations = append(eLocations, location.Name)
	}

	return eLocations
}

func getLocations(client *pokeapiClient, batch int) ([]string, error) {
	res, err := client.Get("https://pokeapi.co/api/v2/location-area/" + fmt.Sprintf("?offset=%v&limit=20", batch*10))
	if err != nil {
		return nil, errors.New("failed to get locations with error: " + err.Error())
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body))
	}
	if err != nil {
		return nil, errors.New("Failed to read body with error: " + err.Error())
	}

	locations := locationAreaResponse{}
	err = json.Unmarshal(body, &locations)

	if err != nil {
		return nil, errors.New("Failed to unmarshal response to json with error: " + err.Error())
	}

	return extractLocationNames(locations), nil
}
