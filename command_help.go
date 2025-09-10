package main

import "fmt"

func commandHelp(c *Config, args []string) error {
	fmt.Println("Welcome to the Pokédex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, value := range getCommands() {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}

	return nil
}