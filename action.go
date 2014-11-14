// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"fmt"
	"strings"
)

const (
	// ActionTagKind is used to identify the Tag type
	ActionTagKind = "action"

	// actionMarker is the identifier used to join filterable
	// prefixes for Action Id's with unique suffixes
	actionMarker = "_a_"
)

//
// ActionTag
//

// ActionTag is a Tag type for representing Action entities, which
// are records of queued actions for a given unit
type ActionTag struct {
	IdPrefixer
}

var _ PrefixTag = (*ActionTag)(nil)

// NewActionTag returns the tag for the action with the given id.
func NewActionTag(id string) ActionTag {
	tag, ok := newActionTag(id)
	if !ok {
		panic(fmt.Sprintf("%q is not a valid action id", id))
	}
	return tag
}

// JoinActionTag reconstitutes an ActionTag from it's prefix and UUID
func JoinActionTag(prefix string, uuid string) ActionTag {
	actionId := fmt.Sprintf("%s%s%s", prefix, actionMarker, uuid)
	tag, ok := newActionTag(actionId)
	if !ok {
		panic("bad prefix or uuid")
	}
	return tag
}

// IsValidAction returns whether actionId is a valid actionId
// Valid action ids include the names.actionMarker token that delimits
// a prefix that can be used for filtering, and a suffix that should be
// unique.  The prefix should match the name rules for units
func IsValidAction(actionId string) bool {
	return isValidIdPrefixTag(actionId, actionMarker)
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
	if !isValidIdPrefixTag(actionId, actionMarker) {
		return ActionTag{}, false
	}
	prefixer := IdPrefixer{
		Id_:     actionId,
		Kind_:   ActionTagKind,
		Marker_: actionMarker,
	}
	return ActionTag{IdPrefixer: prefixer}, true
}

//
// IdPrefixer
//

type PrefixTag interface {
	Tag
	Prefix() string
	UUID() string
	PrefixTag() Tag
}

// IdPrefixer is a common type for representing tags that have
// structured prefixes.
// Note: this type is only used as an embedded type in other Tags that
// use structured prefixes, like ActionTag, and ActionResultTag. We
// have to export this type however so that gccgo is able to serialize
// the embedding types without dropping these fields.
// see: https://bugs.launchpad.net/juju-core/+bug/1381626
type IdPrefixer struct {
	Id_     string
	Kind_   string
	Marker_ string
}

var _ PrefixTag = (*IdPrefixer)(nil)

// Id returns the id of the type this Tag represents
func (t IdPrefixer) Id() string { return t.Id_ }

// String returns a string that shows the type and id of the Tag
func (t IdPrefixer) String() string { return t.Kind_ + "-" + t.Id() }

// Kind exposes the value to identify what kind of Tag this is
func (t IdPrefixer) Kind() string { return t.Kind_ }

// Prefix returns the string representation of the prefix of the Tag
func (t IdPrefixer) Prefix() string {
	prefix, _, ok := splitId(t.Id(), t.Marker_)
	if !ok {
		return ""
	}
	return prefix
}

// UUID returns the unique suffix of the Tag
func (t IdPrefixer) UUID() string {
	_, uuid, ok := splitId(t.Id(), t.Marker_)
	if !ok {
		return ""
	}
	return uuid
}

// PrefixTag returns a Tag representing the Entity matching the id
// prefix
func (t IdPrefixer) PrefixTag() Tag {
	prefix, _, ok := splitId(t.Id(), t.Marker_)
	if !ok {
		return nil
	}

	var tag Tag
	var err error

	switch {
	case IsValidUnit(prefix):
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

// isValidIdPrefixTag signals whether the id is a validly formatted id
// for an IdPrefixer with the given marker
func isValidIdPrefixTag(id, marker string) bool {
	prefix, _, ok := splitId(id, marker)
	if !ok {
		return false
	}
	switch {
	case IsValidUnit(prefix):
	case IsValidService(prefix):
	default:
		return false
	}
	return true
}

// splitId extracts the prefix and suffix from the id using the marker
// token
func splitId(id, marker string) (string, string, bool) {
	parts := strings.Split(id, marker)
	if len(parts) != 2 {
		return "", "", false
	}
	if len(parts[1]) < 1 {
		return "", "", false
	}
	return parts[0], parts[1], true
}
