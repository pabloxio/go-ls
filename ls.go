package main

import "os"

func listDirectory(directory string, config *config) ([]*file, error) {
	dirContent, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	files := make([]*file, 0, len(dirContent))
	for _, entry := range dirContent {
		file, err := newFile(entry)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}
