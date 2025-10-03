package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jonnie-z/boot.dev-pokedex/internal"
	"github.com/jonnie-z/boot.dev-pokedex/types"
)

var (
	cacheOnce      sync.Once
	cacheInstance  *internal.Cache
	configOnce     sync.Once
	configInstance *Config
)

func getCache() *internal.Cache {
	cacheOnce.Do(func() {
		cacheInstance = internal.NewCache(5 * time.Second)
	})

	return cacheInstance
}

func getConfig() *Config {
	configOnce.Do(func() {
		newConfig := Config{
			Next:     "",
			Previous: "",
			Pokemon:  map[string]types.Pokemon{},
		}
		configInstance = &newConfig
	})

	return configInstance
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokédex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		userCmd := input[0]
		args := input[1:]

		cmd, ok := getCommands()[userCmd]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(getConfig(), args)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}
}

func cleanInput(s string) []string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	return strings.Split(s, " ")
}
