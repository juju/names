// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package names

import (
	"strconv"
	"strings"
)

const (
	// ActionTagKind is used to identify the Tag type
	ActionTagKind = "action"

	// ActionMarker is the identifier used to join filterable
	// prefixes for Action Id's with unique suffixes
	ActionMarker = "_a_"
)

// IsAction returns whether actionId is a valid actionId
// Valid action ids include the names.ActionMarker token that delimits
// a prefix that can be used for filtering, and a suffix that should be
// unique.  The prefix should match the name rules for units
func IsAction(actionId string) bool {
	_, ok := parseActionId(actionId)
	return ok
}

// ActionTag is a Tag type for representing Action entities, which
// are records of queued actions for a given unit
type ActionTag struct {
	id string
}

// String returns a string that shows the type and id of an ActionTag
func (t ActionTag) String() string {
	if len(t.id) > 0 {
		return t.Kind() + "-" + t.Id()
	}
	return ""
}

// Kind exposes the ActionTagKind value to identify what kind of Tag this is
func (t ActionTag) Kind() string { return ActionTagKind }

// Id returns the id of the Action this Tag represents
func (t ActionTag) Id() string { return t.id }

// NewActionTag returns the tag for the action with the given id.
func NewActionTag(actionId string) ActionTag {
	if IsAction(actionId) {
		return ActionTag{id: actionId}
	}
	return ActionTag{}
}

// UnitTag will extract and return the UnitTag from the ActionTag
func (t ActionTag) UnitTag() (UnitTag, bool) {
	if parts, ok := parseActionId(t.id); ok {
		return parts.Unit, true
	}
	return UnitTag{}, false
}

// actionIdParts is a convenience struct for holding parsed
// actionId's
type actionIdParts struct {
	Unit     UnitTag
	Sequence int
}

// parseActionId extracts the UnitTag and the unique sequence from the
// actionId.  It returns false if the actionId cannot be parsed otherwise
// true.
func parseActionId(actionId string) (actionIdParts, bool) {
	bad := actionIdParts{}
	parts := strings.Split(actionId, ActionMarker)
	// must have exactly one ActionMarker token
	if len(parts) != 2 {
		return bad, false
	}
	// first part must be a unit name
	tag, ok := unitNameToTag(parts[0])
	if !ok {
		return bad, false
	}

	sl := len(parts[1])
	// sequence has to be at least one digit long
	if sl == 0 {
		return bad, false
	}
	// sequence cannot have leading zero if more than
	// one digit long
	if sl > 1 && strings.HasPrefix(parts[1], "0") {
		return bad, false
	}
	// sequence must be a number (it's generated as int)
	sequence, err := strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		return bad, false
	}
	return actionIdParts{Unit: tag, Sequence: int(sequence)}, true
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
