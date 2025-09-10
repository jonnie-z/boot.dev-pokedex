package main

import "fmt"

func commandPokedex(c *Config, args []string) error {
	pokemonCaught := len(c.Pokemon)
	if pokemonCaught == 0 {
		fmt.Printf("You have not caught any pokeminn :<:<:<:<:<:<:<:<\n")
		return nil
	}

	fmt.Println("LOOKIT DOSE POKEMINN")
	fmt.Println("======POKEDEX======")
	fmt.Println("Pokémon Caught:")
	for _, p := range c.Pokemon {
		fmt.Printf("  - %s\n", p.Name)
	}

	return nil
}