package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jonniez/pokedexcli/internal"
)

var (
	once  sync.Once
	cacheInstance *internal.Cache
)

func getCache() *internal.Cache {
	once.Do(func() {
		cacheInstance = internal.NewCache(5 * time.Second)
	})

	return cacheInstance
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	c := Config{
		Next:     "",
		Previous: "",
	}

	for {
		fmt.Print("Pokédex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		userCmd := input[0]
		args := input [1:]

		cmd, ok := getCommands()[userCmd]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		cmd.callback(&c, args)
	}
}

func cleanInput(s string) []string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	return strings.Split(s, " ")
}
