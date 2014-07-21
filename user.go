// Copyright 2013 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"regexp"
)

const UserTagKind = "user"

var validName = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9.-]*[a-zA-Z0-9]$")

// IsValidUser returns whether id is a valid user id.
func IsValidUser(name string) bool {
	return validName.MatchString(name)
}

type UserTag struct {
	name string
}

func (t UserTag) String() string { return t.Kind() + "-" + t.Id() }
func (t UserTag) Kind() string   { return UserTagKind }
func (t UserTag) Id() string     { return t.name }

// NewUserTag returns the tag for the user with the given name.
func NewUserTag(userName string) UserTag {
	return UserTag{name: userName}
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
