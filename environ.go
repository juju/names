// Copyright 2013 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"github.com/juju/utils"
)

const EnvironTagKind = "environment"

type EnvironTag struct {
	ID utils.UUID
}

// NewEnvironTag returns the tag of an environment with the given environment UUID.
func NewEnvironTag(id string) EnvironTag {
	uuid, err := utils.UUIDFromString(id)
	if err != nil {
		panic(err)
	}
	return EnvironTag{ID: uuid}
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
func (t EnvironTag) Id() string     { return t.ID.String() }

// IsValidEnvironment returns whether id is a valid environment UUID.
func IsValidEnvironment(id string) bool {
	return utils.IsValidUUIDString(id)
}
