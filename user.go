// Copyright 2013 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"fmt"
	"regexp"
)

const (
	UserTagKind   = "user"
	LocalProvider = "local"
)

var validPart = "[a-zA-Z][a-zA-Z0-9.-]*[a-zA-Z0-9]"

var validName = regexp.MustCompile(fmt.Sprintf("^(?P<name>%s)(?:@(?P<provider>%s))?$", validPart, validPart))
var validUserName = regexp.MustCompile("^" + validPart + "$")

// IsValidUser returns whether id is a valid user id.
func IsValidUser(name string) bool {
	return validName.MatchString(name)
}

// IsValidUserName returns whether the user's name is a valid.
func IsValidUserName(name string) bool {
	return validUserName.MatchString(name)
}

type UserTag struct {
	name     string
	provider string
}

func (t UserTag) Kind() string   { return UserTagKind }
func (t UserTag) String() string { return UserTagKind + "-" + t.Id() }

func (t UserTag) Id() string {
	if t.provider == "" {
		return t.name
	}
	return t.name + "@" + t.provider
}

func (t UserTag) Username() string { return t.name + "@" + t.Provider() }
func (t UserTag) Name() string     { return t.name }

func (t UserTag) Provider() string {
	if t.provider == "" {
		return LocalProvider
	}
	return t.provider
}

// NewUserTag returns the tag for the user with the given name.
func NewUserTag(userName string) UserTag {
	parts := validName.FindStringSubmatch(userName)
	if len(parts) != 3 {
		panic(fmt.Sprintf("Invalid user tag %q", userName))
	}
	return UserTag{name: parts[1], provider: parts[2]}
}

// ParseUserTag parser a user tag string.
func ParseUserTag(tag string) (UserTag, error) {
	t, err := ParseTag(tag)
	if err != nil {
		return UserTag{}, err
	}
	ut, ok := t.(UserTag)
	if !ok {
		return UserTag{}, invalidTagError(tag, UserTagKind)
	}
	return ut, nil
}
