package main

import (
	"io/fs"
	"strings"
	"time"
)

type file struct {
	name             string
	isHidden         bool
	mode             string
	user             string
	group            string
	size             int64
	modificationTime time.Time
}

func newFile(entry fs.DirEntry) (*file, error) {
	info, err := entry.Info()
	if err != nil {
		return nil, err
	}

	return &file{
		name:             entry.Name(),
		isHidden:         strings.HasPrefix(entry.Name(), "."),
		mode:             info.Mode().String(),
		size:             info.Size(),
		modificationTime: info.ModTime(),

		// FIXME: hardcoded user and group
		user:  "user",
		group: "group",
	}, nil
}
