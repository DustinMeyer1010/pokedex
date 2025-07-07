package commands

/*
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DustinMeyer1010/pokedexcli/internal/models"
)

func commandExplore(config *Config) error {

	if config.Arugments == "" {
		return fmt.Errorf("location missing\nUse map or mapb to get locations\nUsage: explore {location}")
	}

	res, _ := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", config.Arugments))

	if res.Status == "404 Not Found" {
		return fmt.Errorf("location not found")
	}

	body, _ := io.ReadAll(res.Body)

	var exploredLocation models.LocationExplore

	json.Unmarshal(body, &exploredLocation)

	fmt.Printf("Exploring %s...\n", config.Arugments)

	fmt.Println("Found Pokemon: ")

	for _, pokemon := range exploredLocation.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil

}
*/
