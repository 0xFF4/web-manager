package module

import (
	"encoding/json"
	"fmt"

	"github.com/BurntSushi/toml"
)

// Inclusion is a function type that can be used to include configuration from a file
type Inclusion func([]byte, string, map[string]interface{}) error

// Defaults is a map that contains default inclusion functions for different file types
var Defaults = map[string]Inclusion{
	// This default inclusion function unmarshals json-formatted configuration into a map
	".json": func(b []byte, p string, m map[string]any) error {
		return json.Unmarshal(b, &m)
	},

	// This default inclusion function unmarshals toml-formatted configuration into a map
	".toml": func(b []byte, p string, m map[string]interface{}) error {
		return toml.Unmarshal(b, &m)
	},
}

// NewInclusion is a method that adds a new inclusion function for a given file type
// It returns an error if the file type already has an inclusion function
func (config *Config) NewInclusion(ext string, exec Inclusion) error {
	if _, ok := config.Inclusions[ext]; ok {
		return fmt.Errorf("file type %s already included", ext)
	}

	config.Inclusions[ext] = exec
	return nil
}
