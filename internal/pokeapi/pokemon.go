package pokeapi

import "encoding/json"

func (c *Client) GetPokemon(pokemonName string) (PokemonResponse, error) {
	body, err := c.get(baseURL + "/pokemon/" + pokemonName)

	if err != nil {
		return PokemonResponse{}, err
	}

	pokemon := PokemonResponse{}
	err = json.Unmarshal(body, &pokemon)

	if err != nil {
		return pokemon, err
	}

	return pokemon, nil

}
