package main

import (
	"io/fs"
	"strings"
)

type file struct {
	name     string
	isHidden bool
}

func newFile(entry fs.DirEntry) *file {
	return &file{
		name:     entry.Name(),
		isHidden: strings.HasPrefix(entry.Name(), "."),
	}
}

func (f *file) String() string {
	return f.name
}
