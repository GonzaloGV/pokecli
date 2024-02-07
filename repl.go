package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(c *context) {
	for {
		words := waitForUserCommand()

		if len(words) == 0 {
			fmt.Println("Please enter a command")
			continue
		}

		userCommand := words[0]
		args := words[1:]

		command, exists := c.availableCommands[userCommand]

		if !exists {
			fmt.Println("Command doesn't exist")
			continue
		}

		err := command.callback(c, args...)
		if err != nil {
			fmt.Printf("command %s failed with error -> %s\n", command.name, err.Error())
		}
	}
}

func waitForUserCommand() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Pokedex > ")
	scanner := bufio.NewScanner(reader)
	scanner.Scan()

	return cleanInput(scanner.Text())
}

func cleanInput(rawInput string) []string {
	out := strings.ToLower(rawInput)
	words := strings.Fields(out)

	return words
}
