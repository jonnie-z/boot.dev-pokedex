package main

import "fmt"

func commandExplore(c *Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("wrong number of args")
	}

	fmt.Printf("Exploring %s...", args[0])

	pokemon, err := getLocationPokemon(args[0])
	if err != nil {
		return fmt.Errorf("error getting pokémon: %w", err)
	}

	fmt.Printf("Found the following Pokémon:\n")
	for _, result := range pokemon.PokemonEncounters {
		fmt.Printf("  - %s\n", result.Pokemon.Name)
	}

	return nil
}