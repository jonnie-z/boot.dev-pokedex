package main

import "github.com/jonniez/pokedexcli/types"

type Config struct {
	Next     string
	Previous string
	Pokemon  map[string]types.Pokemon
}
