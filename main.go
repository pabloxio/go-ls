package main

import "flag"

func main() {
	config := newConfig()

	flag.BoolVar(&config.showHiddenFiles, "a", false, "show hidden files")
	flag.Parse()

	directory := flag.Arg(0)
	if directory == "" {
		directory = "."
	}

	err := listDirectory(directory, config)
	if err != nil {
		panic(err)
	}
}
