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
		commands.Command{
			Name:     "map",
			Desc:     "Show the map locations",
			Callback: commands.CommandMap,
			View:     mapView,
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
