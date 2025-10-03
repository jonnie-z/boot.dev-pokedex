package main

import "github.com/jonnie-z/boot.dev-pokedex/types"

type Config struct {
	Next     string
	Previous string
	Pokemon  map[string]types.Pokemon
}
