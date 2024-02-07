package main

import (
	"fmt"
)

func commandHelp(c *context, args ...string) error {
	fmt.Print("\n\nWelcome to the pokedex!\nUsage:\n\n")
	for _, command := range c.availableCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
