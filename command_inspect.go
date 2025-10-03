package main

import "fmt"

func commandInspect(c *Config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("wrong number of args")
	}

	val, ok := c.Pokemon[args[0]]
	if !ok {
		fmt.Printf("You have not caught a --%s--...\n", args[0])
		return nil
	}

	fmt.Printf("LOOKIT DAT %s\n", val.Name)

	fmt.Printf("=========%s=========\n", val.Name)
	fmt.Printf("Name: %v\n", val.Name)
	fmt.Printf("Height: %v\n", val.Height)
	fmt.Printf("Weight: %v\n", val.Weight)
	fmt.Println("Stats:")
	for _, s := range val.Stats {
		fmt.Printf("  - %s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range val.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	fmt.Printf("=========%s=========\n", val.Name)

	return nil
}
