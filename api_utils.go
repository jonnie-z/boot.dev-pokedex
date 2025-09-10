package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/jonniez/pokedexcli/types"
)

func getLocations(urlEnd string) (types.LocationResponse, error) {
	url := locationAreaURL + urlEnd

	cache := getCache()
	if val, exists := cache.Get(url); exists {
		locations := types.LocationResponse{}
		decoder := json.NewDecoder(bytes.NewBuffer(val))
		err := decoder.Decode(&locations)
		if err != nil {
			return types.LocationResponse{}, fmt.Errorf("issue decoding json: %w", err)
		}

		return locations, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return types.LocationResponse{}, fmt.Errorf("issue getting %s: %w", locationAreaURL, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return types.LocationResponse{}, fmt.Errorf("issue reading data: %w", err)
	}

	locations := types.LocationResponse{}
	decoder := json.NewDecoder(bytes.NewBuffer(data))
	err = decoder.Decode(&locations)
	if err != nil {
		return types.LocationResponse{}, fmt.Errorf("issue decoding json: %w", err)
	}

	cache.Add(url, data)

	return locations, nil
}

func getLocationPokemon(areaName string) (types.LocationAreaResponse, error) {
	url := locationAreaURL + areaName

	cache := getCache()
	if val, exists := cache.Get(url); exists {
		locations := types.LocationAreaResponse{}
		decoder := json.NewDecoder(bytes.NewBuffer(val))
		err := decoder.Decode(&locations)
		if err != nil {
			return types.LocationAreaResponse{}, fmt.Errorf("issue decoding json: %w", err)
		}

		return locations, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return types.LocationAreaResponse{}, fmt.Errorf("issue getting %s: %w", locationAreaURL, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return types.LocationAreaResponse{}, fmt.Errorf("issue reading data: %w", err)
	}

	locations := types.LocationAreaResponse{}
	decoder := json.NewDecoder(bytes.NewBuffer(data))
	err = decoder.Decode(&locations)
	if err != nil {
		return types.LocationAreaResponse{}, fmt.Errorf("issue decoding json: %w", err)
	}

	cache.Add(url, data)

	return locations, nil
}

func getPokemon(pokemonName string) (types.Pokemon, error) {
	url := pokemonURL + pokemonName

	cache := getCache()
	if val, exists := cache.Get(url); exists {
		pokemon := types.Pokemon{}
		decoder := json.NewDecoder(bytes.NewBuffer(val))
		err := decoder.Decode(&pokemon)
		if err != nil {
			return types.Pokemon{}, fmt.Errorf("issue decoding json: %w", err)
		}

		return pokemon, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return types.Pokemon{}, fmt.Errorf("issue getting %s: %w", locationAreaURL, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return types.Pokemon{}, fmt.Errorf("issue reading data: %w", err)
	}

	pokemon := types.Pokemon{}
	decoder := json.NewDecoder(bytes.NewBuffer(data))
	err = decoder.Decode(&pokemon)
	if err != nil {
		return types.Pokemon{}, fmt.Errorf("issue decoding json: %w", err)
	}

	cache.Add(url, data)

	return pokemon, nil
}


func getNewUrl(s string) string {
	if s != "" {
		return strings.TrimPrefix(s, locationAreaURL)
	}

	return ""
}
