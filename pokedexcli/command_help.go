package main

import "fmt"

func commandHelp(c *config) error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println("")
		cMap := getCommands()
		for k := range cMap {
			fmt.Printf("%s: %s\n", cMap[k].name, cMap[k].description)
		}
		return nil
	}

