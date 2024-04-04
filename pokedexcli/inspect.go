package main

import (
	"errors"
	"fmt"

	"github.com/jsMRSoL/pokedexcli/internal/pokemon"
)

func inspectPokemon(c *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("You have to include a pokemon name.")
		fmt.Println("E.g. inspect tentacool")
		return errors.New("No pokemon name submitted for inspection")
	}
	pokemon_name := args[0]

	var pd pokemon.PokemonData
	pd, found := c.pokedex.Get(pokemon_name)
	if !found {
		fmt.Printf("You have not caught %s\n", pokemon_name)
	}

	fmt.Printf("Name: %s!\n", pd.Name)
	fmt.Printf("Height: %d\n", pd.Height)
	fmt.Printf("Weight: %d\n", pd.Weight)
	fmt.Println("Stats:")
	for _, stat := range pd.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, _type := range pd.Types {
		fmt.Printf("  - %s\n", _type.Type.Name)
	}

	return nil
}
