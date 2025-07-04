package commands

import (
	"fmt"
	"net/http"
)

func commandExplore(config *Config) error {

	if config.Arugments == "" {
		return fmt.Errorf("location missing\nUse map or mapb to get locations\nUsage: explore {location}")
	}

	res, _ := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", config.Arugments))

	if res.Status == "404 Not Found" {
		return fmt.Errorf("location not found")
	}

	fmt.Printf("Exploring %s\n", config.Arugments)

	return nil

}
