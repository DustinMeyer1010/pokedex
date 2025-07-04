package commands

import "fmt"

func commandHelp(config *Config) error {
	helpMessage :=
		`
Welcome to  the Pokedex!
Usage:

`
	for _, c := range Commands {
		helpMessage += fmt.Sprintf("%s: %s\n", c.Name, c.Description)
	}
	fmt.Println(helpMessage)
	return nil
}
