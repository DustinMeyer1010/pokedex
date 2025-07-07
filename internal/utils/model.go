package utils

import (
	"github.com/DustinMeyer1010/pokedexcli/internal/commands"
	"github.com/DustinMeyer1010/pokedexcli/internal/models"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	mainView models.ViewState = iota
	mapView
	inputView
)

type command struct {
	name        string
	description string
	callback    func(config *commands.Config) ([]list.Item, error)
	view        models.ViewState
}

func (c command) GetName() string {
	return c.name
}

func (c command) Title() string       { return c.name }
func (c command) Description() string { return c.description }
func (c command) FilterValue() string { return c.name }

type model struct {
	currentView models.ViewState
	commandMenu list.Model
	mapMenu     list.Model
}

// Init implements tea.Model.
func (m model) Init() tea.Cmd {
	return nil
}

func InitialModel() model {
	commands := []list.Item{
		command{
			name:        "map",
			description: "Show the map locations",
			callback:    commands.CommandMap,
			view:        mapView,
		},
	}

	l := list.New(commands, itemDelegate{}, 50, 10)
	l.Title = "What would you like to do?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return model{
		currentView: mainView,
		commandMenu: l,
		mapMenu:     l,
	}
}
