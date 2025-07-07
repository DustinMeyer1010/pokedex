package main

import (
	"fmt"
	"os"

	"github.com/DustinMeyer1010/pokedexcli/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(utils.InitialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
