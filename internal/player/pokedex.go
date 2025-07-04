package player

import (
	"fmt"

	"github.com/DustinMeyer1010/pokedexcli/internal/models"
)

var pokedex map[string]models.Pokemon

func AddPokemon(pokemon models.Pokemon) error {
	pokedex[pokemon.Name] = pokemon
	return nil
}

func RetrievePokemon(pokemonName string) (*models.Pokemon, error) {
	var p models.Pokemon
	var exist bool

	if p, exist = pokedex[pokemonName]; !exist {
		return nil, fmt.Errorf("You do not own the pokemon")
	}

	return &p, nil
}
