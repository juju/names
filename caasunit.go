// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"fmt"
	"regexp"
	"strings"
)

const CAASUnitTagKind = "caasunit"

var validCAASUnit = regexp.MustCompile("^(" + ApplicationSnippet + ")/" + NumberSnippet + "$")

type CAASUnitTag struct {
	name string
}

func (t CAASUnitTag) String() string { return t.Kind() + "-" + t.name }
func (t CAASUnitTag) Kind() string   { return CAASUnitTagKind }
func (t CAASUnitTag) Id() string     { return caasUnitTagSuffixToId(t.name) }

// NewCAASUnitTag returns the tag for the caasUnit with the given name.
// It will panic if the given caasunit name is not valid.
func NewCAASUnitTag(caasUnitName string) CAASUnitTag {
	tag, ok := tagFromCAASUnitName(caasUnitName)
	if !ok {
		panic(fmt.Sprintf("%q is not a valid caasunit name", caasUnitName))
	}
	return tag
}

// ParseCAASUnitTag parses a caasUnit tag string.
func ParseCAASUnitTag(caasUnitTag string) (CAASUnitTag, error) {
	tag, err := ParseTag(caasUnitTag)
	if err != nil {
		return CAASUnitTag{}, err
	}
	ut, ok := tag.(CAASUnitTag)
	if !ok {
		return CAASUnitTag{}, invalidTagError(caasUnitTag, CAASUnitTagKind)
	}
	return ut, nil
}

// IsValidCAASUnit returns whether name is a valid caasunit name.
func IsValidCAASUnit(name string) bool {
	return validCAASUnit.MatchString(name)
}

// CAASUnitApplication returns the name of the application that the caasUnit is
// associated with. It returns an error if caasUnitName is not a valid caasunit name.
func CAASUnitApplication(caasUnitName string) (string, error) {
	s := validCAASUnit.FindStringSubmatch(caasUnitName)
	if s == nil {
		return "", fmt.Errorf("%q is not a valid caasunit name", caasUnitName)
	}
	return s[1], nil
}

func tagFromCAASUnitName(caasUnitName string) (CAASUnitTag, bool) {
	// Replace only the last "/" with "-".
	i := strings.LastIndex(caasUnitName, "/")
	if i <= 0 || !IsValidCAASUnit(caasUnitName) {
		return CAASUnitTag{}, false
	}
	caasUnitName = caasUnitName[:i] + "-" + caasUnitName[i+1:]
	return CAASUnitTag{name: caasUnitName}, true
}

func caasUnitTagSuffixToId(s string) string {
	// Replace only the last "-" with "/", as it is valid for application
	// names to contain hyphens.
	if i := strings.LastIndex(s, "-"); i > 0 {
		s = s[:i] + "/" + s[i+1:]
	}
	return s
}
