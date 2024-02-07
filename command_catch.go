package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *context, args ...string) error {
	if len(args) == 0 {
		return errors.New("Please specify the pokemon you want to catch.")
	}

	if len(args) > 1 {
		return errors.New("You only can catch one pokemon at a time.")
	}

	pokemonName := args[0]

	pokemon, err := c.client.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s\n", pokemon.Name)
	catched := tryCatchPokemon(pokemon.BaseExperience)

	if !catched {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	c.pokedex[pokemon.Name] = pokemon
	return nil
}

func tryCatchPokemon(baseExp int) bool {
	expFactor := baseExp / 10

	generated := rand.Int31n(int32(expFactor))
	return generated < 10
}
