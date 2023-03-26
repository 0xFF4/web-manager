package module

import (
	"fmt"
	"reflect"
)

// Options struct that holds the Config instance
type Options struct {
	config *Config
}

// Back returns the parent Config instance
func (options *Options) Back() *Config {
	return options.config
}

// Options creates a new Options instance from the Config instance
// It returns an error if the Config instance's data has not been parsed
func (config *Config) Options() (*Options, error) {
	if config.Data == nil || len(config.Data) <= 0 {
		return nil, fmt.Errorf("must execute Parse() once with targets before calling this function")
	}

	return &Options{
		config: config,
	}, nil
}

// Get retrieves the value at the specified path in the Config instance's configs map
// It returns an error if the value at the specified path is not found or is not of the expected type
func (options *Options) Get(t reflect.Kind, p ...string) (interface{}, error) {
	scope := options.config.Data

	// Iterate over the path elements and traverse the configs map
	for i, element := range p {
		if i == len(p)-1 {
			// Check if the value at the specified path has the expected type
			if reflect.ValueOf(scope[element]).Kind() == t {
				return scope[element], nil
			}
			return nil, fmt.Errorf("item not found")
		}

		// Check if the value at the current path element is a map
		if reflect.ValueOf(scope[element]).Kind() == reflect.Map {
			scope = scope[element].(map[string]interface{})
			continue
		}

		return nil, fmt.Errorf("item not found")
	}

	return nil, fmt.Errorf("item not found")
}

// GetFromVectors retrieves the value at the specified path in the Config instance's configs map with the expected types
// It returns an error if the value at the specified path is not found or is not of any of the expected types
func (options *Options) GetFromVectors(vectors []reflect.Kind, p ...string) (interface{}, error) {
	// Iterate over the expected types
	for _, vector := range vectors {
		// Check if the value at the specified path has the expected type
		if item, err := options.Get(vector, p...); err == nil {
			return item, nil
		}
	}

	return nil, fmt.Errorf("item not found")
}
