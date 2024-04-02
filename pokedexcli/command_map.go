package main

import (
	"fmt"

	"github.com/jsMRSoL/pokedexcli/internal/locations"
)

func mapForward(c *config) error {
	ld, err := api.GetLocationsData(c.next)
	if err != nil {
		return err
	}

	c.next = ld.Next
	c.prev = ld.Previous

	for _, v := range ld.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func mapBackward(c *config) error {
	if c.prev == nil {
		fmt.Println("Already at the start!")
		return nil
	}

	ld, err := api.GetLocationsData(c.prev)
	if err != nil {
		return err
	}

	c.next = ld.Next
	c.prev = ld.Previous

	for _, v := range ld.Results {
		fmt.Println(v.Name)
	}

	return nil
}
