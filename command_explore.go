package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *context, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must specify an area")
	}
	area := args[0]

	pokemonResponse, err := c.client.ListPokemons(area)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")
	for _, encounter := range pokemonResponse.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil

}
