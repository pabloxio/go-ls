package main

import (
	"flag"
	"fmt"
)

func main() {
	config := newConfig()

	flag.BoolVar(&config.showHiddenFiles, "a", false, "show hidden files")
	flag.BoolVar(&config.longListingFormat, "l", false, "long listing format")
	flag.Parse()

	directory := flag.Arg(0)
	if directory == "" {
		directory = "."
	}

	files, err := listDirectory(directory, config)
	if err != nil {
		panic(err)
	}

	fmt.Println(files)
}
