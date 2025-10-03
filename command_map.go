package main

import "fmt"

func commandMap(c *Config, args []string) error {
	if c.Next == "" {
		c.Next = firstNext
	}
	urlEnd := "" + c.Next

	locations, err := getLocations(urlEnd)
	if err != nil {
		return fmt.Errorf("error getting locations: %w", err)
	}

	c.Next = getNewUrl(locations.Next)
	c.Previous = getNewUrl(locations.Previous)

	for _, result := range locations.Results {
		fmt.Printf("%s\n", result.Name)
	}

	return nil
}
