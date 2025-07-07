package commands

import (
	"github.com/DustinMeyer1010/pokedexcli/internal/models"
	"github.com/charmbracelet/bubbles/list"
)

type Config struct {
	Next      string
	Previous  string
	Arugments string
}

type Command struct {
	Name     string
	Desc     string
	Callback func(config *Config) ([]list.Item, error)
	View     models.ViewState
}

func (c Command) GetName() string {
	return c.Name
}

func (c Command) Title() string       { return c.Name }
func (c Command) Description() string { return c.Desc }
func (c Command) FilterValue() string { return c.Name }

/*
type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}


var Commands map[string]CliCommand = map[string]CliCommand{}

func InitCommands() {
	Commands["help"] = CliCommand{
		Name:        "help",
		Description: "Displays help message",
		Callback:    commandHelp,
	}
	Commands["exit"] = CliCommand{
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    commandExit,
	}
	Commands["explore"] = CliCommand{
		Name:        "explore",
		Description: "Explore a location (Usage: explore {location})",
		Callback:    commandExplore,
	}
	Commands["catch"] = CliCommand{
		Name:        "catch",
		Description: "Catch a pokemon (Usage: catch {pokemon})",
		Callback:    commandCatch,
	}
	Commands["inspect"] = CliCommand{
		Name:        "inspect",
		Description: "Inspect pokemon in pokedex (Usage: inspect {pokemon})",
		Callback:    commandInspect,
	}
	Commands["pokedex"] = CliCommand{
		Name:        "pokedex",
		Description: "Show all pokemon in pokedex",
		Callback:    commandPokedex,
	}
}
*/
