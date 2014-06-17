// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package names

import (
	"regexp"
)

const (
	// ActionTagKind is used to identify the Tag type
	ActionTagKind = "action"

	// ActionMarker is the identifier used to join filterable
	// prefixes for Action Id's with unique suffixes
	ActionMarker  = "_a_"
)

const (
	validActionNameRegex = ServiceSnippet + "([-/]" + NumberSnippet + ")?" + ActionMarker + NumberSnippet
)

var validAction = regexp.MustCompile("^" + validActionNameRegex + "$")

// IsAction returns whether name is a valid action name.
// Valid action names include the names.ActionMarker token that delimits
// a prefix that can be used for filtering, and a suffix that should be
// unique.  The prefix should match the name rules for units and services
func IsAction(name string) bool {
	return validAction.MatchString(name)
}

// ActionTag is a Tag type for representing Action entities, which
// are records of queued actions for a given service or unit
type ActionTag struct {
	name string
}

// String returns a string that shows the type and id of an ActionTag
func (t ActionTag) String() string { return t.Kind() + "-" + t.Id() }

// Kind exposes the ActionTagKind value to identify what kind of Tag this is
func (t ActionTag) Kind() string   { return ActionTagKind }

// Id returns the name of the Action this Tag represents
func (t ActionTag) Id() string     { return t.name }

// NewActionTag returns the tag for the action with the given name.
func NewActionTag(actionName string) Tag {
	return ActionTag{name: actionName}
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
