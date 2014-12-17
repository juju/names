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

var validUserPart = "[a-zA-Z0-9][a-zA-Z0-9.-]*[a-zA-Z0-9]"

var validName = regexp.MustCompile(fmt.Sprintf("^(?P<name>%s)(?:@(?P<provider>%s))?$", validUserPart, validUserPart))
var validUserName = regexp.MustCompile("^" + validUserPart + "$")

// IsValidUser returns whether id is a valid user id.
func IsValidUser(name string) bool {
	return validName.MatchString(name)
}

// IsValidUserName returns whether the user's name is a valid.
func IsValidUserName(name string) bool {
	return validUserName.MatchString(name)
}

// UserTag represents a user that may be stored in the local database, or provided
// through some remote identity provider.
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

// IsLocal returns true if the tag represents a local user.
func (t UserTag) IsLocal() bool {
	return t.Provider() == LocalProvider
}

// Provider returns the name of the user provider. Users in the local database
// are from the LocalProvider. Other users are considered 'remote' users.
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
		panic(fmt.Sprintf("invalid user tag %q", userName))
	}
	return UserTag{name: parts[1], provider: parts[2]}
}

// NewLocalUserTag returns the tag for a local user with the given name.
func NewLocalUserTag(name string) UserTag {
	if !IsValidUserName(name) {
		panic(fmt.Sprintf("invalid user name %q", name))
	}
	return UserTag{name: name, provider: LocalProvider}
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
