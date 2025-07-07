package utils

import (
	"github.com/DustinMeyer1010/pokedexcli/internal/commands"
	tea "github.com/charmbracelet/bubbletea"
)

var config commands.Config = commands.Config{
	Next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
	Previous: "",
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.currentView {
		case mainView:
			switch msg.String() {
			case "ctrl-c":
				return m, tea.Quit
			case "enter":
				item := m.commandMenu.SelectedItem().(command) // Get the current slice
				locations, _ := item.callback(&config)
				m.mapMenu.SetItems(locations)
				m.currentView = item.view

			}
		case mapView:
			switch msg.String() {
			case "ctrl-c":
				return m, tea.Quit
			case "n":
				if config.Next != "" {
					locations, _ := commands.CommandMap(&config)
					m.mapMenu.SetItems(locations)
				}
			case "p":
				if config.Previous != "" {
					locations, _ := commands.CommandMapb(&config)
					m.mapMenu.SetItems(locations)
				}
			case "b":
				m.currentView = mainView

			}
		}

	}

	if m.currentView == mainView {
		m.commandMenu, cmd = m.commandMenu.Update(msg)
	}

	if m.currentView == mapView {
		m.mapMenu, cmd = m.mapMenu.Update(msg)
	}

	return m, cmd
}
