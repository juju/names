// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"fmt"
	"regexp"
)

const DiskTagKind = "disk"

var validDisk = regexp.MustCompile("^" + NumberSnippet + "$")

type DiskTag struct {
	name string
}

func (t DiskTag) String() string { return t.Kind() + "-" + t.name }
func (t DiskTag) Kind() string   { return DiskTagKind }
func (t DiskTag) Id() string     { return t.name }

// NewDiskTag returns the tag for the disk with the given name.
// It will panic if the given disk name is not valid.
func NewDiskTag(diskName string) DiskTag {
	tag, ok := tagFromDiskName(diskName)
	if !ok {
		panic(fmt.Sprintf("%q is not a valid disk name", diskName))
	}
	return tag
}

// ParseDiskTag parses a disk tag string.
func ParseDiskTag(diskTag string) (DiskTag, error) {
	tag, err := ParseTag(diskTag)
	if err != nil {
		return DiskTag{}, err
	}
	dt, ok := tag.(DiskTag)
	if !ok {
		return DiskTag{}, invalidTagError(diskTag, DiskTagKind)
	}
	return dt, nil
}

// IsValidDisk returns whether name is a valid disk name.
func IsValidDisk(name string) bool {
	return validDisk.MatchString(name)
}

func tagFromDiskName(diskName string) (DiskTag, bool) {
	if !IsValidDisk(diskName) {
		return DiskTag{}, false
	}
	return DiskTag{diskName}, true
}
