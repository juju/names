// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package names

import (
	"strings"
)

const EnvironTagKind = "environment"

type EnvironTag struct {
	uuid string
}

// NewEnvironTag returns the tag of an environment with the given environment UUID.
func NewEnvironTag(uuid string) EnvironTag {
	return EnvironTag{uuid: uuid}
}

// ParseEnvironTag parses an environ tag string.
func ParseEnvironTag(environTag string) (EnvironTag, error) {
	tag, err := ParseTag(environTag)
	if err != nil {
		return EnvironTag{}, err
	}
	et, ok := tag.(EnvironTag)
	if !ok {
		return EnvironTag{}, invalidTagError(environTag, EnvironTagKind)
	}
	return et, nil
}

func (t EnvironTag) String() string { return t.Kind() + "-" + t.Id() }
func (t EnvironTag) Kind() string   { return EnvironTagKind }
func (t EnvironTag) Id() string     { return t.uuid }

// IsEnvironment returns whether id is a valid environment UUID.
func IsEnvironment(id string) bool {
	// TODO(axw) 2013-12-04 #1257587
	// We should not accept environment tags that
	// do not look like UUIDs.
	return !strings.Contains(id, "/")
}
