package config

import (
	"secure-me/source/config/module"
)

var (
	Options *module.Options
)

// ParseConfiguration parses the configuration from the specified folder
func ParseConfiguration() error {
	config := module.NewConfig()
	err := config.Parse("resources")
	if err != nil {
		return err
	}

	// Retrieve the options from the parsed configuration
	Options, err = config.Options()
	if err != nil {
		return err
	}

	return nil
}
