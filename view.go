package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/DustinMeyer1010/pokedexcli/internal/commands"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type itemDelegate struct{}

type item interface {
	GetName() string
}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.GetName())

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("-> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type viewState int

const (
	mainView viewState = iota
	mapView
	inputView
)

type command struct {
	name        string
	description string
	callback    func(config *commands.Config) error
	view        viewState
}

func (c command) GetName() string {
	return c.name
}

func (c command) Title() string       { return c.name }
func (c command) Description() string { return c.description }
func (c command) FilterValue() string { return c.name }

type location struct {
	name string
	view viewState
}

func (l location) GetName() string {
	return l.name
}

func (l location) Title() string       { return l.name }
func (l location) Description() string { return "test" }
func (l location) FilterValue() string { return l.name }

type model struct {
	currentView viewState
	commandMenu list.Model
	mapMenu     list.Model
}

// Init implements tea.Model.
func (m model) Init() tea.Cmd {
	return nil
}

func initialModel() model {
	commands := []list.Item{
		command{
			name:        "map",
			description: "Show the map locations",
			callback:    commands.CommandMap,
			view:        mapView,
		},
		command{
			name:        "map2",
			description: "Show the map locations",
			callback:    commands.CommandMap,
			view:        mapView,
		},
		command{
			name:        "map3",
			description: "Show the map locations",
			callback:    commands.CommandMap,
			view:        mapView,
		},
	}

	l := list.New(commands, itemDelegate{}, 50, 10)
	l.Title = "What do you want for dinner?"
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
				if m.currentView == mainView {
					item := m.commandMenu.SelectedItem().(command) // Get the current slice
					m.mapMenu.SetItems([]list.Item{location{"test location", mapView}, location{"test location2", mapView}})
					m.currentView = item.view
				}
			}
		case mapView:
			switch msg.String() {
			case "ctrl-c":
				return m, tea.Quit
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

func (m model) View() string {
	switch m.currentView {
	case mainView:
		return m.commandMenu.View()

	case mapView:

		return m.mapMenu.View() + fmt.Sprint("frameworks item count:", len(m.mapMenu.Items()))
	default:
		return "Loading..."
	}
}
