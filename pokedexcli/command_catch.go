package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/jsMRSoL/pokedexcli/internal/pokemon"
)

func catchPokemon(c *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("You have to include a pokemon name.")
		fmt.Println("E.g. catch tentacool")
		return errors.New("No pokemon name submitted")
	}
	pokemon_name := args[0]

	var pd pokemon.PokemonData
	pd, found := c.pokedex.Get(pokemon_name)
	if !found {
		urlString := "https://pokeapi.co/api/v2/pokemon/" + pokemon_name
		var err error
		pd, _, err = pokemon.GetPokemonData(urlString)
		if err != nil {
			return err
		}
	}

	catchChance := rand.Float32() * (1.0 + float32(pd.BaseExperience)/255)
	name := pd.Name
	if catchChance >= 0.5 {
		fmt.Printf("You caught %s!", name)
		c.pokedex.Add(pokemon_name, pd)
	} else {
		fmt.Printf("%s escaped!", name)
	}

	return nil
}
