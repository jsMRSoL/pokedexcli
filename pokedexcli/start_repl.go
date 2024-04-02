package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		cmd, args, err := normalizeInput(input)

		if err == nil {
			if _, ok := cmds[cmd]; !ok {
				fmt.Println("Command not recognised. Try help!")
				fmt.Println()
			} else {
				cmds[cmd].callback(config, args)
				fmt.Println()
			}
		}
	}
}

func normalizeInput(input string) (cmd string, args []string, err error) {
	wds := strings.Fields(input)
	if len(wds) == 0 {
		fmt.Println("You have to enter a command. Try help!")
		fmt.Println()
		return "", args, err
	}

	cmd = wds[0]

	if len(wds) > 1 {
		args = wds[1:]
	}

	return cmd, args, nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	prev  *string
	next  *string
	cache *pokecache.PokeCache
}

func newConfig() *config {
	first := "https://pokeapi.co/api/v2/location-area/"
	// first := "notaurl.noreally"
	pc := pokecache.NewCache(time.Minute * 5)
	return &config{
		prev:  nil,
		next:  &first,
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
		"explore": {
			name:        "explore",
			description: "Explore an area <name>, looking for Pokemon",
			callback:    exploreArea,
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
