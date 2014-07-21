// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// ActionTagKind is used to identify the Tag type
	ActionTagKind = "action"

	// actionMarker is the identifier used to join filterable
	// prefixes for Action Id's with unique suffixes
	actionMarker = "_a_"
)

// ActionTag is a Tag type for representing Action entities, which
// are records of queued actions for a given unit
type ActionTag struct {
	id string
}

// NewActionTag returns the tag for the action with the given id.
func NewActionTag(id string) ActionTag {
	tag, ok := newActionTag(id)
	if !ok {
		return ActionTag{}
	}
	return tag
}

// JoinActionTag reconstitutes an ActionTag from it's prefix and sequence
func JoinActionTag(prefix string, sequence int) ActionTag {
	actionId := fmt.Sprintf("%s%s%d", prefix, actionMarker, sequence)
	return ActionTag{id: actionId}
}

// String returns a string that shows the type and id of an ActionTag
func (t ActionTag) String() string {
	return t.Kind() + "-" + t.Id()
}

// Kind exposes the ActionTagKind value to identify what kind of Tag this is
func (t ActionTag) Kind() string { return ActionTagKind }

// Id returns the id of the Action this Tag represents
func (t ActionTag) Id() string { return t.id }

// PrefixTag returns a Tag representing the ActionReceiver this action is
// queued for
func (t ActionTag) PrefixTag() Tag {
	prefix, _, ok := splitActionId(t.Id())
	if !ok {
		return nil
	}

	var tag Tag
	var err error

	switch {
	case IsUnit(prefix):
		tag = NewUnitTag(prefix)
	case IsValidService(prefix):
		tag = NewServiceTag(prefix)
	default:
		tag, err = ParseTag(prefix)
		if err != nil {
			tag = nil
		}
	}
	return tag
}

func (t ActionTag) Sequence() int {
	_, sequence, ok := splitActionId(t.Id())
	if !ok {
		return -1
	}
	return sequence
}

// IsValidAction returns whether actionId is a valid actionId
// Valid action ids include the names.actionMarker token that delimits
// a prefix that can be used for filtering, and a suffix that should be
// unique.  The prefix should match the name rules for units
func IsValidAction(actionId string) bool {
	_, ok := newActionTag(actionId)
	return ok
}

// ParseActionTag parses a action tag string.
func ParseActionTag(actionTag string) (ActionTag, error) {
	tag, err := ParseTag(actionTag)
	if err != nil {
		return ActionTag{}, err
	}
	st, ok := tag.(ActionTag)
	if !ok {
		return ActionTag{}, invalidTagError(actionTag, ActionTagKind)
	}
	return st, nil
}

func newActionTag(actionId string) (ActionTag, bool) {
	bad := ActionTag{}
	prefix, _, ok := splitActionId(actionId)
	if !ok {
		return bad, false
	}
	switch {
	case IsUnit(prefix):
	case IsValidService(prefix):
	default:
		return bad, false
	}
	return ActionTag{id: actionId}, true
}

func splitActionId(id string) (string, int, bool) {
	parts := strings.Split(id, actionMarker)
	if len(parts) != 2 {
		return "", -(len(parts) + 100), false
	}
	if len(parts[1]) > 1 && parts[1][:1] == "0" {
		return "", -2, false
	}
	seq, err := strconv.ParseInt(parts[1], 10, 0)
	if err != nil {
		return "", -3, false
	}
	return parts[0], int(seq), true
}
