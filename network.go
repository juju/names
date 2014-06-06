// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package names

import (
	"fmt"
	"regexp"
)

const NetworkTagKind = "network"

var validNetwork = regexp.MustCompile("^([a-z0-9]+(-[a-z0-9]+)*)$")

// IsNetwork reports whether name is a valid network name.
var IsNetwork = validNetwork.MatchString

type networkTag struct {
	name string
}

func (t networkTag) String() string {
	return NetworkTagKind + "-" + t.name
}

// NetworkTag returns the tag of a network with the given name.
func NetworkTag(name string) Tag {
	if !IsNetwork(name) {
		panic(fmt.Sprintf("%q is not a valid network name", name))
	}
	return networkTag{name: name}
}
