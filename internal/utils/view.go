package utils

import "strings"

func (m model) View() string {
	var body strings.Builder

	switch m.currentView {
	case mainView:
		body.WriteString(m.commandMenu.View())

	case mapView:
		m.mapMenu.Title = "Locations Found"
		body.WriteString(m.mapMenu.View())
	default:
		body.WriteString("Loading")
	}

	body.WriteString("\n[n] Next Menu - [p] Previous Menu")

	return body.String()
}
