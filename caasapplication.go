// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"regexp"
)

const CAASApplicationTagKind = "caasapplication"

const (
	CAASApplicationSnippet = "(?:[a-z][a-z0-9]*(?:-[a-z0-9]*[a-z][a-z0-9]*)*)"
)

var validCAASApplication = regexp.MustCompile("^" + CAASApplicationSnippet + "$")

// IsValidCAASApplication returns whether name is a valid application name.
func IsValidCAASApplication(name string) bool {
	return validCAASApplication.MatchString(name)
}

type CAASApplicationTag struct {
	Name string
}

func (t CAASApplicationTag) String() string { return t.Kind() + "-" + t.Id() }
func (t CAASApplicationTag) Kind() string   { return CAASApplicationTagKind }
func (t CAASApplicationTag) Id() string     { return t.Name }

// NewCAASApplicationTag returns the tag for the application with the given name.
func NewCAASApplicationTag(applicationName string) CAASApplicationTag {
	return CAASApplicationTag{Name: applicationName}
}

// ParseCAASApplicationTag parses a application tag string.
func ParseCAASApplicationTag(applicationTag string) (CAASApplicationTag, error) {
	tag, err := ParseTag(applicationTag)
	if err != nil {
		return CAASApplicationTag{}, err
	}
	st, ok := tag.(CAASApplicationTag)
	if !ok {
		return CAASApplicationTag{}, invalidTagError(applicationTag, CAASApplicationTagKind)
	}
	return st, nil
}
