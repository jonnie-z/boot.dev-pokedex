package main

import (
	"fmt"
	"os"
)

func commandExit(c *Config, args []string) error {
	fmt.Println("Closing the Pokédex...goodbye!")
	os.Exit(0)
	return nil
}