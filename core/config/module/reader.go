package module

import (
	"os"
	"path/filepath"
)

// execDir recursively scans the given directory and applies the registered inclusion functions to all the supported file types
func (config *Config) execDir(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		// Check if a registered inclusion function exists for the file's extension
		file, ok := config.Inclusions[filepath.Ext(info.Name())]
		if !ok {
			return nil
		}

		// Read the file and apply the registered inclusion function
		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		config.Renders[path] = contents
		return file(contents, path, config.Data)
	})
}
