package commands

/*
import (
	"fmt"

	"github.com/DustinMeyer1010/pokedexcli/internal/models"
	"github.com/DustinMeyer1010/pokedexcli/internal/player"
)

func commandInspect(config *Config) error {

	var err error
	var pokemon *models.Pokemon

	if config.Arugments == "" {
		return fmt.Errorf("Please choose a pokemon from pokedex\nUsage: inspect {pokemon}")
	}

	if pokemon, err = player.RetrievePokemon(config.Arugments); err != nil {
		fmt.Errorf("%s not in Pokedex", config.Arugments)
	}

	fmt.Printf(`
Name: %s
Height: %d
Weight: %d
`,
		pokemon.Name, pokemon.Height, pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, kind := range pokemon.Types {
		fmt.Printf("- %s\n", kind.Type.Name)
	}

	return nil
}
*/
