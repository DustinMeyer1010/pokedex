package commands

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type Config struct {
	Next      string
	Previous  string
	Arugments string
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
	Commands["map"] = CliCommand{
		Name:        "map",
		Description: "Shows Next 20 locations",
		Callback:    commandMap,
	}
	Commands["mapb"] = CliCommand{
		Name:        "mapb",
		Description: "Shows Previous 20 locations",
		Callback:    commandMapb,
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

}
