// Copyright 2016 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"regexp"
)

const ModelTagKind = "model"

// EnvironModelTag represents a tag used to describe a model.
type EnvironModelTag struct {
	uuid string
}

var validUUID = regexp.MustCompile(`[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`)

// NewModelTag returns the tag of an model with the given model UUID.
func NewModelTag(uuid string) EnvironModelTag {
	return EnvironModelTag{uuid: uuid}
}

// ParseModelTag parses an environ tag string.
func ParseModelTag(ModelTag string) (EnvironModelTag, error) {
	tag, err := ParseTag(ModelTag)
	if err != nil {
		return EnvironModelTag{}, err
	}
	et, ok := tag.(EnvironModelTag)
	if !ok {
		return EnvironModelTag{}, invalidTagError(ModelTag, ModelTagKind)
	}
	return et, nil
}

func (t EnvironModelTag) String() string { return t.Kind() + "-" + t.Id() }
func (t EnvironModelTag) Kind() string   { return ModelTagKind }
func (t EnvironModelTag) Id() string     { return t.uuid }

// IsValidModel returns whether id is a valid model UUID.
func IsValidModel(id string) bool {
	return validUUID.MatchString(id)
}
