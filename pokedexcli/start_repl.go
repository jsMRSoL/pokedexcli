package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/jsMRSoL/pokedexcli/internal/pokecache"
)

func startRepl() {
	cmds := getCommands()
	config := newConfig()
	for {
		fmt.Printf("pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if _, ok := cmds[input]; !ok {
			fmt.Println("Command not recognised. Try help!")
			fmt.Println()
		} else {
			cmds[input].callback(config)
			fmt.Println()
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	prev *string
	next *string
  cache *pokecache.PokeCache
}

func newConfig() *config {
	first := "https://pokeapi.co/api/v2/location-area/"
	// first := "notaurl.noreally"
  pc := pokecache.NewCache(time.Minute * 5)
	return &config{
		prev: nil,
		next: &first,
    cache: pc,
	}
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Display the next 20 locations in the Pokemon world",
			callback:    mapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 locations in the Pokemon world",
			callback:    mapBackward,
		},
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
