// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"fmt"
	"regexp"
)

const ActionTagKind = "action"

type ActionTag struct {
	// Tags that are serialized need to have fields exported.
	UUID string
}

// validActionUUID describes the valid UUID's for an ActionTag.
var validActionUUID = regexp.MustCompile(`^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`)

// NewActionTag returns the tag of an action with the given id UUID.
func NewActionTag(uuid string) ActionTag {
	if !IsValidAction(uuid) {
		panic(fmt.Sprintf("%q is not a valid action id", uuid))
	}
	return ActionTag{UUID: uuid}
}

// ParseActionTag parses an action tag string.
func ParseActionTag(actionTag string) (ActionTag, error) {
	tag, err := ParseTag(actionTag)
	if err != nil {
		return ActionTag{}, err
	}
	at, ok := tag.(ActionTag)
	if !ok {
		return ActionTag{}, invalidTagError(actionTag, ActionTagKind)
	}
	return at, nil
}

func (t ActionTag) String() string { return t.Kind() + "-" + t.Id() }
func (t ActionTag) Kind() string   { return ActionTagKind }
func (t ActionTag) Id() string     { return t.UUID }

// IsValidAction returns whether id is a valid action id UUID.
func IsValidAction(id string) bool {
	return validActionUUID.MatchString(id)
}

func ActionReceiverTag(name string) (Tag, error) {
	if IsValidUnit(name) {
		return NewUnitTag(name), nil
	}
	if IsValidService(name) {
		// TODO(jcw4) enable when leader elections complete
		//return NewServiceTag(name), nil
	}
	return nil, fmt.Errorf("invalid actionreceiver name %q", name)
}
