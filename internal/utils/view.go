package utils

import (
	"strings"
)

func (m model) View() string {
	var body strings.Builder

	switch m.currentView {
	case mainView:
		body.WriteString(m.commandMenu.View())
		body.WriteString("")
	case mapView:
		m.mapMenu.Title = "Map Locations"
		body.WriteString(m.mapMenu.View())
		body.WriteString("\n[n] Next - [p] Previous [up/down] - Change Selection")
	case pokemonView:
		m.pokemonMenu.Title = "Pokemon Found"
		return m.pokemonMenu.View()
		//body.WriteString(m.pokemonMenu.View())
	default:
		body.WriteString("Loading")
	}
	body.WriteString("\n[enter] Select - [q] Quit")

	return body.String()
}
