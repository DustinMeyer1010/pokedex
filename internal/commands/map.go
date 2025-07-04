package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/DustinMeyer1010/pokedexcli/internal/cache"
)

type locationAreas struct {
	Count    int
	Next     string
	Previous string
	Results  []location
}

type location struct {
	Name string
	Url  string
}

var Cache cache.Cache = *cache.NewCache(time.Second * 60)

func commandMap(config *Config) error {
	var body []byte
	var exist bool

	if body, exist = Cache.Get(config.Next); !exist {
		res, err := http.Get(config.Next)

		if err != nil {
			return fmt.Errorf("fail to retrieve locations")
		}
		body, err = io.ReadAll(res.Body)
		Cache.Add(config.Next, body)

		if err != nil {
			return fmt.Errorf("failed to read response body")
		}

	}

	var data locationAreas
	err := json.Unmarshal(body, &data)

	if err != nil {
		return fmt.Errorf("fail to parse locations")
	}

	config.Next = data.Next
	config.Previous = data.Previous

	for _, results := range data.Results {
		fmt.Println(results.Name)
	}

	if exist {
		fmt.Println("pulling cached data")
	}
	return err
}

func commandMapb(config *Config) error {
	var body []byte
	var exist bool

	if config.Previous == "" {
		return fmt.Errorf("no previous locations")
	}

	if body, exist = Cache.Get(config.Previous); !exist {
		res, err := http.Get(config.Previous)

		if err != nil {
			return fmt.Errorf("fail to retrieve locations")
		}

		body, err = io.ReadAll(res.Body)
		Cache.Add(config.Previous, body)

		if err != nil {
			return fmt.Errorf("failed to read response body")
		}
	}

	var data locationAreas
	err := json.Unmarshal(body, &data)

	if err != nil {
		return fmt.Errorf("fail to parse locations")
	}

	config.Next = data.Next
	config.Previous = data.Previous

	for _, results := range data.Results {
		fmt.Println(results.Name)
	}

	if exist {
		fmt.Println("pulling cached data")
	}
	return err
}
