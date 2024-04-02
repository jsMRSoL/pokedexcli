package main

import "os"

func commandExit(c *config, _ []string) error {
		os.Exit(0)
		return nil
	}

