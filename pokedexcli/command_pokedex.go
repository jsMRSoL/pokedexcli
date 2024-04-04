package main

import (
	"fmt"
)

func inspectPokedex(c *config, args []string) error {

	if c.pokedex.Len() == 0 {
		fmt.Println("You have no Pokemon in your Pokedex.")
		return nil
	}

	c.pokedex.List()

	return nil
}
