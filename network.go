// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package names

import (
	"fmt"
	"regexp"
)

const NetworkTagKind = "network"

var ValidNetwork = regexp.MustCompile("^([a-z0-9]+(-[a-z0-9]+)*)$")

// IsNetwork reports whether name is a valid network name.
func IsNetwork(name string) bool {
	return ValidNetwork.MatchString(name)
}

type NetworkTag struct {
	name string
}

func (t NetworkTag) String() string { return t.Kind() + "-" + t.Id() }
func (t NetworkTag) Kind() string   { return NetworkTagKind }
func (t NetworkTag) Id() string     { return t.name }

// NewNetworkTag returns the tag of a network with the given name.
func NewNetworkTag(name string) NetworkTag {
	if !IsNetwork(name) {
		panic(fmt.Sprintf("%q is not a valid network name", name))
	}
	return NetworkTag{name: name}
}

// ParseNetworkTag parses a network tag string.
func ParseNetworkTag(networkTag string) (NetworkTag, error) {
	tag, err := ParseTag(networkTag)
	if err != nil {
		return NetworkTag{}, err
	}
	nt, ok := tag.(NetworkTag)
	if !ok {
		return NetworkTag{}, invalidTagError(networkTag, NetworkTagKind)
	}
	return nt, nil
}
