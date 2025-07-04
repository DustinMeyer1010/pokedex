package main

import (
	"fmt"
	"testing"

	"github.com/DustinMeyer1010/pokedexcli/internal/utils"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"Charmander", "Bulbasaur", "PIKACHU"},
		},
		{
			input:    "hello  world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "              Charmander Bulbasaur    PIKACHU",
			expected: []string{"Charmander", "Bulbasaur", "PIKACHU"},
		},
		{
			input:    "h e l l o",
			expected: []string{"h", "e", "l", "l", "o"},
		},
		{
			input:    "Charmander      Bulbasaur    PIKACHU                       ",
			expected: []string{"Charmander", "Bulbasaur", "PIKACHU"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "l",
			expected: []string{"l"},
		},
	}

	for _, c := range cases {
		actual := utils.CleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%s), \nReturn Length: %d, \nExpected Length: %d", c.input, len(c.input), len(c.expected))
		}
		for i := range actual {
			fmt.Println(i)
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("cleanInput(%s),\nReturn: %s,\nExpected: %s", c.input, word, expectedWord)
			}
		}
	}

}
