package main

import (
	"errors"
	"fmt"

	"github.com/gv/pokedexcli/internal/pokeapi"
)

func commandMap(c *context, args ...string) error {
	mapState := &c.mapState

	if len(mapState.current) != 0 && mapState.next == "" {
		return errors.New("There are no more maps")
	}

	locations, err := c.client.ListLocations(&mapState.next)

	if err != nil {
		return err
	}

	mapState.next = locations.Next
	mapState.prev = locations.Previous

	displayLocations(&locations)

	return nil
}

func commandMapBack(c *context, args ...string) (err error) {
	mapState := &c.mapState

	if mapState.prev == "" {
		return errors.New("There are no previous maps")
	}

	locations, err := c.client.ListLocations(&mapState.prev)

	if err != nil {
		return err
	}

	mapState.next = locations.Next
	mapState.prev = locations.Previous

	displayLocations(&locations)

	return nil
}

func displayLocations(locations *pokeapi.LocationAreaResponse) {
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
}
