package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DustinMeyer1010/pokedexcli/internal/models"
	"github.com/charmbracelet/bubbles/list"
)

func CommandExplore(config *Config) (pokemon []list.Item, err error) {

	if config.Arugments == "" {
		return nil, fmt.Errorf("location missing\nUse map or mapb to get locations\nUsage: explore {location}")
	}

	res, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", config.Arugments))

	if err != nil {
		return nil, err
	}

	if res.Status == "404 Not Found" {
		return nil, fmt.Errorf("location not found")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var exploredLocation models.LocationExplore

	json.Unmarshal(body, &exploredLocation)

	for _, p := range exploredLocation.PokemonEncounters {
		pokemon = append(pokemon, models.Pokemon{Name: p.Pokemon.Name})
	}

	return pokemon, nil

}
