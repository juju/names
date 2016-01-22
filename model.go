// Copyright 2016 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"regexp"
)

const ModelTagKind = "model"

// EnvironModel represents a tag used to describe a model.
type EnvironModel struct {
	uuid string
}

var validUUID = regexp.MustCompile(`[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`)

// NewModelTag returns the tag of an model with the given model UUID.
func NewModelTag(uuid string) EnvironModel {
	return EnvironModel{uuid: uuid}
}

// ParseModelTag parses an environ tag string.
func ParseModelTag(modelTag string) (EnvironModel, error) {
	tag, err := ParseTag(modelTag)
	if err != nil {
		return EnvironModel{}, err
	}
	et, ok := tag.(EnvironModel)
	if !ok {
		return EnvironModel{}, invalidTagError(modelTag, ModelTagKind)
	}
	return et, nil
}

func (t EnvironModel) String() string { return t.Kind() + "-" + t.Id() }
func (t EnvironModel) Kind() string   { return ModelTagKind }
func (t EnvironModel) Id() string     { return t.uuid }

// IsValidModel returns whether id is a valid model UUID.
func IsValidModel(id string) bool {
	return validUUID.MatchString(id)
}
