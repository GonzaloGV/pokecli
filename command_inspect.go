package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *context, args ...string) error {
	if len(args) == 0 {
		return errors.New("You must specify a pokemon")
	}

	pokedex := c.pokedex
	pokemonName := args[0]

	pokemon, catched := pokedex[pokemonName]

	if !catched {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Height: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil
}
