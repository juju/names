// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package names

import (
	"regexp"
	"strings"
)

const MachineTagKind = "machine"

const (
	ContainerTypeSnippet = "[a-z]+"
	ContainerSnippet     = "(/" + ContainerTypeSnippet + "/" + NumberSnippet + ")"
	MachineSnippet       = NumberSnippet + ContainerSnippet + "*"
)

var validMachine = regexp.MustCompile("^" + MachineSnippet + "$")

// IsMachine returns whether id is a valid machine id.
func IsMachine(id string) bool {
	return validMachine.MatchString(id)
}

// IsContainerMachine returns whether id is a valid container machine id.
func IsContainerMachine(id string) bool {
	return validMachine.MatchString(id) && strings.Contains(id, "/")
}

type machineTag struct {
	id string
}

func (t machineTag) String() string { return strings.Replace(MachineTagKind+"-"+t.id, "/", "-", -1) }

// MachineTag returns the tag for the machine with the given id.
func MachineTag(id string) Tag {
	return machineTag{id: id}
}

func machineTagSuffixToId(s string) string {
	return strings.Replace(s, "-", "/", -1)
}
