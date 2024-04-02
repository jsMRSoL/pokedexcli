package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jsMRSoL/pokedexcli/internal/locations"
)

func mapForward(c *config, args []string) error {
	data, found := c.cache.Get(*c.next)
	var ld locations.LocationsData
	if found {
		if err := json.Unmarshal(data, &ld); err != nil {
			log.Println("Error decoding cached JSON: ", err)
			return err
		}
    // log.Println("Returning data from cache...")
	} else {
		var err error
    var bytes []byte
		ld, bytes, err = locations.GetLocationsData(c.next)
		if err != nil {
			return err
		}
    // log.Println("Returning data from api...")
    // log.Println("Caching url: ", *c.next)
    c.cache.Add(*c.next, bytes)
	}

	c.next = ld.Next
	c.prev = ld.Previous

	for _, v := range ld.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func mapBackward(c *config, _ []string) error {
	if c.prev == nil {
		fmt.Println("Already at the start!")
		return nil
	}

	data, found := c.cache.Get(*c.prev)
	var ld locations.LocationsData
	if found {
		if err := json.Unmarshal(data, &ld); err != nil {
			log.Println("Error decoding JSON: ", err)
			return err
		}
    // log.Println("Returning data from cache...")
	} else {
    var err error
    var bytes []byte
		ld, bytes, err = locations.GetLocationsData(c.prev)
		if err != nil {
			return err
		}
    // log.Println("Returning data from api...")
    // This is safe because we checked if c.prev == nil above
    // log.Println("Caching url: ", *c.prev)
    c.cache.Add(*c.prev, bytes)
	}

	c.next = ld.Next
	c.prev = ld.Previous

	for _, v := range ld.Results {
		fmt.Println(v.Name)
	}

	return nil
}
