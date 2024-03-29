package main

import (
	"time"

	"github.com/gv/pokedexcli/internal/pokeapi"
)

func createContext() context {
	return context{
		availableCommands: getAvailableCommands(),
		client:            pokeapi.NewClient(time.Duration(time.Duration.Seconds(5))),
		mapState:          mapState{},
		pokedex:           map[string]pokeapi.PokemonResponse{},
	}

}

type commandsMap = map[string]cliCommand

type mapState struct {
	prev    string
	next    string
	current []string
}

type context struct {
	availableCommands commandsMap
	client            pokeapi.Client
	mapState          mapState
	pokedex           map[string]pokeapi.PokemonResponse
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *context, args ...string) error
}

func getAvailableCommands() commandsMap {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Show available map options. Locations are batched by 20 maps, every subsequent map will show the next 20 locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show the previous batch of maps",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore pokemons in a given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemons. The chance of caught the pokemon decrease with the pokemon's base expirience",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Use pokedex to see pokemon stats",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "show pokemons in pokedex",
			callback:    commandPokedex,
		},
	}
}
