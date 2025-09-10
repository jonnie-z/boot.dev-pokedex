package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("wrong number of args")
	}
	
	pokemon, err := getPokemon(args[0])
	if err != nil {
		return fmt.Errorf("error getting Pokémon %s: %w", args[0], err)
	}

	randNum := rand.Intn(100) + 1

	caught := calculatePercentage(pokemon.BaseExperience, pokemon.Weight) < randNum

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	if caught {
		c.Pokemon[pokemon.Name] = pokemon
		fmt.Printf("%s caught!!\n", pokemon.Name)
	} else {
		fmt.Printf("Oooh, so close! But no %s...\n", pokemon.Name)
	}

	return nil
}


func calculatePercentage(base_experience int, weight int) int {
	baseRate := (base_experience / 3) + (weight / 10 / 4)


	return 1 + ((baseRate - rateMin) * 99) / (rateMax - rateMin)
}