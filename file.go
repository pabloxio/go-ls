package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"strings"
	"syscall"
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

	user, group := getUserAndGroupOwners(&info)

	return &file{
		name:             entry.Name(),
		isHidden:         strings.HasPrefix(entry.Name(), "."),
		mode:             info.Mode().String(),
		user:             user,
		group:            group,
		size:             info.Size(),
		modificationTime: info.ModTime(),
	}, nil
}

func getUserAndGroupOwners(fileInfo *os.FileInfo) (userName, groupName string) {
	statSyscall := (*fileInfo).Sys().(*syscall.Stat_t)

	userName = fmt.Sprint(statSyscall.Uid)
	groupName = fmt.Sprint(statSyscall.Gid)

	// Getting the user
	owner, err := user.LookupId(userName)
	if err != nil {
		return
	}
	userName = owner.Username

	// Getting the group
	group, err := user.LookupGroupId(groupName)
	if err != nil {
		return
	}
	groupName = group.Name

	return
}
