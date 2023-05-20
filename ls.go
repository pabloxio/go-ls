package main

import (
	"fmt"
	"os"

	"github.com/gosuri/uitable"
)

func listDirectory(directory string, config *config) (*uitable.Table, error) {
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

	table := uitable.New()
	table.Separator = " "

	for _, f := range files {
		if !config.showHiddenFiles && f.isHidden {
			continue
		}

		if config.longListingFormat {
			// Right align the size column
			table.RightAlign(3)

			table.AddRow(
				f.mode,
				f.user,
				f.group,
				fmt.Sprintf("%d", f.size),
				f.modificationTime.Format("Jan 02 15:04"),
				f.name,
			)
		} else {
			table.AddRow(f.name)
		}
	}

	return table, nil
}
