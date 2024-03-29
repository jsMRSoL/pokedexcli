package main

import (
	"fmt"
	"github.com/jsMRSoL/pokedexcli/internal/api"
)

func mapForward(c *config) error {
	var url string
	if c.index == -1 {
		url = "https://pokeapi.co/api/v2/location-area/"
		urls := c.urls
		urls = append(urls, url)
		c.urls = urls
		c.index = 0
	} else {
		c.index++
		url = c.urls[c.index]
	}

	ld, err := api.GetLocationsData(url)
	if err != nil {
		return err
	}

	c.urls = append(c.urls, ld.Next)

	for _, v := range ld.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func mapBackward(c *config) error {
	var url string
	if c.index == -1 {
		mapForward(c)
		return nil
	}
	if c.index == 0 {
		fmt.Println("Already at the start!")
		return nil
	}

	c.index--
	url = c.urls[c.index]

	ld, err := api.GetLocationsData(url)
	if err != nil {
		return err
	}

	for _, v := range ld.Results {
		fmt.Println(v.Name)
	}

	return nil
}
