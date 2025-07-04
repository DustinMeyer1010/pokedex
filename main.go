package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DustinMeyer1010/pokedexcli/internal/commands"
	"github.com/DustinMeyer1010/pokedexcli/internal/utils"
)

var cfg commands.Config = commands.Config{Next: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20", Previous: ""}

func main() {
	commands.InitCommands()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := utils.CleanInput(scanner.Text())
			if len(input) == 0 {
				continue
			}
			command := input[0]

			if c, exist := commands.Commands[command]; exist {
				if len(input) >= 2 {
					cfg.Arugments = input[1]
				}
				err := c.Callback(&cfg)
				if err != nil {
					fmt.Println(err)
				}
				continue
			}
			fmt.Printf("Unknown Command %s\n", command)
		}

	}
}
