package main

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokédex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays a list of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous list of locations",
			callback:    commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Displays all the Pokémon in the area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "(Attempt to) catch a Pokémon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "lookit dat pokeminn",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "das alotta pokeminn (or not)",
			callback: commandPokedex,
		},
	}
}

/* // Logger represents a simple logging service.
type Logger struct {
	// Add any fields specific to your logger here
	logCount int
}

var (
	once     sync.Once
	instance *Logger
)

// GetLogger returns the singleton instance of Logger.
func GetLogger() *Logger {
	once.Do(func() {
		// This code block will be executed only once
		instance = &Logger{logCount: 0}
		fmt.Println("Logger initialized!")
	})
	return instance
}

// Log a message.
func (l *Logger) Log(message string) {
	l.logCount++
	fmt.Printf("Log #%d: %s\n", l.logCount, message)
}

func main() {
	// Get the logger instance multiple times
	logger1 := GetLogger()
	logger2 := GetLogger()

	// Both variables point to the same instance
	fmt.Printf("Logger1 address: %p\n", logger1)
	fmt.Printf("Logger2 address: %p\n", logger2)

	logger1.Log("First message")
	logger2.Log("Second message")
} */