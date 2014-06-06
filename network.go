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

type NetworkTag struct {
	name string
}

func (t NetworkTag) String() string {
	return NetworkTagKind + "-" + t.name
}

// NewNetworkTag returns the tag of a network with the given name.
func NewNetworkTag(name string) Tag {
	if !IsNetwork(name) {
		panic(fmt.Sprintf("%q is not a valid network name", name))
	}
	return NetworkTag{name: name}
}
