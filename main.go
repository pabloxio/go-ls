package main

import (
	"fmt"

	"github.com/gosuri/uitable"
	flag "github.com/spf13/pflag"
)

func main() {
	config := newConfig()

	flag.BoolVarP(&config.showHiddenFiles, "all", "a", false, "show hidden files")
	flag.BoolVarP(&config.longListingFormat, "long", "l", false, "long listing format")
	flag.Parse()

	directory := flag.Arg(0)
	if directory == "" {
		directory = "."
	}

	files, err := listDirectory(directory, config)
	if err != nil {
		panic(err)
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

	fmt.Println(table)
}
