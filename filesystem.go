// Copyright 2015 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	FilesystemTagKind = "filesystem"
	filenameSnippet   = "[\\w. ][\\w\\-. ]*"
)

var validFilesystem = regexp.MustCompile("^" + filenameSnippet + "$")

type FilesystemTag struct {
	name string
}

func (t FilesystemTag) String() string { return t.Kind() + "-" + t.name }
func (t FilesystemTag) Kind() string   { return FilesystemTagKind }
func (t FilesystemTag) Id() string     { return t.name }

// NewFilesystemTag returns the tag for the filesystem with the given name.
// It will panic if the given filesystem name is not valid.
func NewFilesystemTag(filesystemName string) FilesystemTag {
	tag, ok := tagFromFilesystemName(filesystemName)
	if !ok {
		panic(fmt.Sprintf("%q is not a valid filesystem name", filesystemName))
	}
	return tag
}

// ParseFilesystemTag parses a filesystem tag string.
func ParseFilesystemTag(filesystemTag string) (FilesystemTag, error) {
	tag, err := ParseTag(filesystemTag)
	if err != nil {
		return FilesystemTag{}, err
	}
	fstag, ok := tag.(FilesystemTag)
	if !ok {
		return FilesystemTag{}, invalidTagError(filesystemTag, FilesystemTagKind)
	}
	return fstag, nil
}

// IsValidFilesystem returns whether name is a valid filesystem name.
func IsValidFilesystem(name string) bool {
	return validFilesystem.MatchString(name)
}

func tagFromFilesystemName(filesystemName string) (FilesystemTag, bool) {
	if !IsValidFilesystem(filesystemName) {
		return FilesystemTag{}, false
	}
	return FilesystemTag{filesystemName}, true
}

func filesystemTagSuffixToId(s string) string {
	// Replace only the last "-" with "/", as it is valid for filesystem
	// names to contain hyphens.
	if i := strings.LastIndex(s, "-"); i > 0 {
		s = s[:i] + "/" + s[i+1:]
	}
	return s
}
