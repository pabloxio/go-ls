package main

import (
	"fmt"
	"os"
)

func listDirectory(directory string, config *config) error {
	dirContent, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	files := make([]*file, 0, len(dirContent))
	for _, entry := range dirContent {
		files = append(files, newFile(entry))
	}

	for _, file := range files {
		if !config.showHiddenFiles && file.isHidden {
			continue
		}

		fmt.Printf("%s\n", file)
	}

	return nil
}
