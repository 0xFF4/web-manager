package config

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

// ScopeDir will attempt to locate all files and folder within the passed directory
func ScopeDir(dir ...string) error {
	//files, err := os.ReadDir(filepath.Join(dir...))
	//if err != nil {
	//	return err
	//}
	//
	//// Range trough all files/folders inside the directory
	//for _, file := range files {
	//	if file.IsDir() {
	//		fmt.Printf("Found /%s in %s\r\n", file.Name(), filepath.Join(dir...))
	//		fmt.Println(ScopeDir("resources/" + file.Name()))
	//	} else {
	//		fmt.Printf("Found %s in %s\r\n", file.Name(), filepath.Join(dir...))
	//	}
	//}
	return filepath.Walk(filepath.Join(dir...), func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			fmt.Printf("Found /%s in %s\r\n", info.Name(), filepath.Join(dir...))
		} else {
			fmt.Printf("Found %s in %s\r\n", info.Name(), filepath.Join(dir...))
		}
		return nil
	})

}
