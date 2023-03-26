package module

import "fmt"

// Config struct that holds the parsed configuration data
type Config struct {
	// Data is a map that stores the parsed configuration data
	Data map[string]interface{}

	// Inclusions is a map that maps file types to inclusion functions
	Inclusions map[string]Inclusion

	// Renders is a map that maps file paths to the raw contents of the file
	Renders map[string][]byte
}

// NewConfig is a constructor for Config
// It creates a new instance of Config with the default inclusion functions and empty data and renders maps
func NewConfig() *Config {
	return &Config{
		// Set the default inclusion functions
		Inclusions: Defaults,
		// Initialize an empty data map
		Data: make(map[string]interface{}),
		// Initialize an empty renders map
		Renders: make(map[string][]byte),
	}
}

// Parse is a method that parses configuration data from a list of directories
// It initializes the data map if it has not been initialized yet and parses each directory
func (config *Config) Parse(directories ...string) error {
	// Initialize the data map if it is not already initialized
	if config.Data == nil {
		config.Data = make(map[string]interface{})
	}

	// Parse each directory
	for _, dir := range directories {
		if err := config.execDir(dir); err != nil {
			return fmt.Errorf("error executing directory %s: %v", dir, err)
		}
	}

	return nil
}

// Files is a method that returns the raw contents of all parsed configuration files
func (config *Config) Files() map[string][]byte {
	return config.Renders
}
