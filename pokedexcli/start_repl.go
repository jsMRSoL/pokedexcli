package main

import (
  "fmt"
  "bufio"
  "os"
)

func startRepl() {
	cmds := getCommands()
	for {
		fmt.Printf("pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if _, ok := cmds[input]; !ok {
			fmt.Println("Command not recognised. Try help!")
      fmt.Println()
		} else {
			cmds[input].callback()
      fmt.Println()
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

