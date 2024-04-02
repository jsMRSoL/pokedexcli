package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/jsMRSoL/pokedexcli/internal/pokemon"
)

func getUrlString(args []string) (string, error) {
	//construct area name from args
	var areaName string
	if len(args) == 0 {
		log.Println("You have to include an area name after the command,")
		log.Println("E.g. explore canalave-city-area")
		log.Println("or   explore canalave city area")
		return "", errors.New("Area name could not be parsed.")
	}
	if len(args) == 1 {
		areaName = args[0]
	}
	if len(args) > 1 {
		areaName = strings.Join(args, "-")
	}
	areaName = strings.ToLower(areaName)

	return "https://pokeapi.co/api/v2/location-area/" + areaName, nil
}

func exploreArea(c *config, args []string) error {
	url, err := getUrlString(args)
	if err != nil {
		return err
	}

	data, found := c.cache.Get(url)
	var pe pokemon.PokemonEncounters
	if found {
		if err := json.Unmarshal(data, &pe); err != nil {
			log.Println("Error decoding cached JSON: ", err)
			return err
		}
		// log.Println("Returning data from cache...")
	} else {
		var err error
		var bytes []byte
		pe, bytes, err = pokemon.GetPokemonList(url)
		if err != nil {
			return err
		}
		// log.Println("Returning data from location-area api...")
		// log.Println("Caching url: ", url)
		c.cache.Add(url, bytes)
	}

	fmt.Printf("Exploring %s...\n", pe.Name)
	fmt.Println("Found Pokemon:")
	for _, v := range pe.PokemonEncounters {
		fmt.Println(" - ", v.Pokemon.Name)
	}

	return nil
}
