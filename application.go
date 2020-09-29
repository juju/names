// Copyright 2013 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/juju/errors"
)

// ApplicationTagKind defines a tag for identifying applications.
const ApplicationTagKind = "application"

const (
	// ApplicationSnippet is a non-compiled regexp that can be composed with
	// other snippets to form a valid application regexp.
	ApplicationSnippet = "(?:[a-z][a-z0-9]*(?:-[a-z0-9]*[a-z][a-z0-9]*)*)"
)

var (
	validApplication = regexp.MustCompile("^" + ApplicationSnippet + "$")
	tailNumberSuffix = regexp.MustCompile("-[0-9]+$")
)

// IsValidApplication returns whether name is a valid application name.
func IsValidApplication(name string) bool {
	return validApplication.MatchString(name)
}

// ValidateApplicationName takes a name and attempts to validate the application
// name, before returning a reason why it's not valid.
//
// This should supersede IsValidApplication. It provides a lot more granular
// information about why something might be wrong, which is a much better UX.
func ValidateApplicationName(name string) error {
	if IsValidApplication(name) {
		return nil
	}

	// If the application has uppercase characters, bail out and explain
	// why.
	if uppercaseChar.MatchString(name) {
		return errors.Errorf("invalid application name %q, unexpected uppercase character", name)
	}
	// If the application ends up being suffixed by a number, then we want
	// to mention it to users why.
	if tailNumberSuffix.MatchString(name) {
		return errors.Errorf("invalid application name %q, unexpected number(s) found after last hyphen", name)
	}

	index := strings.IndexFunc(name, invalidRuneForApplicationName)
	if index < 0 {
		return errors.Errorf("invalid application name %q", name)
	}

	// We have to ensure that we don't break up multi-rune characters, by
	// just selecting the index. Instead look at a slice of runes and use
	// the first one.
	invalidRune := []rune(name[index:])[0]
	return errors.Errorf("invalid application name %q, unexpected character %c", name, invalidRune)
}

// invalidRuneForApplicationName works out if there is a valid application rune.
func invalidRuneForApplicationName(r rune) bool {
	if (r >= 'a' && r <= 'z') || unicode.IsNumber(r) || r == '-' {
		return false
	}
	return true
}

// ApplicationTag defines a named tagged application.
type ApplicationTag struct {
	Name string
}

func (t ApplicationTag) String() string { return t.Kind() + "-" + t.Id() }

// Kind returns the application tag.
func (t ApplicationTag) Kind() string { return ApplicationTagKind }

// Id returns the underlying name of an application as the Id.
func (t ApplicationTag) Id() string { return t.Name }

// NewApplicationTag returns the tag for the application with the given name.
func NewApplicationTag(applicationName string) ApplicationTag {
	return ApplicationTag{Name: applicationName}
}

// ParseApplicationTag parses a application tag string.
func ParseApplicationTag(applicationTag string) (ApplicationTag, error) {
	tag, err := ParseTag(applicationTag)
	if err != nil {
		return ApplicationTag{}, err
	}
	st, ok := tag.(ApplicationTag)
	if !ok {
		return ApplicationTag{}, invalidTagError(applicationTag, ApplicationTagKind)
	}
	return st, nil
}
