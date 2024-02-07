package main

import "fmt"

func commandPokedex(c *context, args ...string) error {
	fmt.Println("Your Pokedex:")
	if len(c.pokedex) == 0 {
		fmt.Println("0 Pokemons")
		return nil
	}

	for _, pokemon := range c.pokedex {
		fmt.Println(" - " + pokemon.Name)
	}

	return nil
}
