// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package names

import (
	"regexp"
)

const UserTagKind = "user"

var validName = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9.-]*[a-zA-Z0-9]$")

// IsUser returns whether id is a valid user id.
var IsUser = validName.MatchString

type userTag struct {
	name string
}

func (t userTag) String() string {
	return UserTagKind + "-" + t.name
}

// UserTag returns the tag for the user with the given name.
func UserTag(userName string) Tag {
	return userTag{name: userName}
}
