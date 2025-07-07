package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/DustinMeyer1010/pokedexcli/internal/cache"
	"github.com/DustinMeyer1010/pokedexcli/internal/models"
	"github.com/charmbracelet/bubbles/list"
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

func CommandMap(config *Config) (locations []list.Item, err error) {
	var body []byte
	var exist bool

	if body, exist = Cache.Get(config.Next); !exist {
		res, err := http.Get(config.Next)

		if err != nil {
			return nil, fmt.Errorf("fail to retrieve locations")
		}
		body, err = io.ReadAll(res.Body)
		Cache.Add(config.Next, body)

		if err != nil {
			return nil, fmt.Errorf("failed to read response body")
		}

	}

	var data locationAreas
	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, fmt.Errorf("fail to parse locations")
	}

	config.Next = data.Next
	config.Previous = data.Previous

	for _, loc := range data.Results {
		locations = append(locations, models.Location{Name: loc.Name, Desc: ""})
	}

	return locations, err
}

func CommandMapb(config *Config) (locations []list.Item, err error) {
	var body []byte
	var exist bool

	if config.Previous == "" {
		return nil, fmt.Errorf("no previous locations")
	}

	if body, exist = Cache.Get(config.Previous); !exist {
		res, err := http.Get(config.Previous)

		if err != nil {
			return nil, fmt.Errorf("fail to retrieve locations")
		}

		body, err = io.ReadAll(res.Body)
		Cache.Add(config.Previous, body)

		if err != nil {
			return nil, fmt.Errorf("failed to read response body")
		}
	}

	var data locationAreas
	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, fmt.Errorf("fail to parse locations")
	}

	config.Next = data.Next
	config.Previous = data.Previous

	for _, loc := range data.Results {
		locations = append(locations, models.Location{Name: loc.Name, Desc: ""})
	}

	return locations, err

}
