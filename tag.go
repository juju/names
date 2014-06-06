// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package names

import (
	"fmt"
	"strings"
)

// A Tag tags things that are taggable.
type Tag interface {
	fmt.Stringer // all Tags should be able to print themselves
}

// TagKind returns one of the *TagKind constants for the given tag, or
// an error if none matches.
func tagKind(tag string) (string, error) {
	i := strings.Index(tag, "-")
	if i <= 0 || !validKinds(tag[:i]) {
		return "", fmt.Errorf("%q is not a valid tag", tag)
	}
	return tag[:i], nil
}

func validKinds(kind string) bool {
	switch kind {
	case UnitTagKind, MachineTagKind, ServiceTagKind, EnvironTagKind, UserTagKind, RelationTagKind, NetworkTagKind:
		return true
	}
	return false
}

func splitTag(tag string) (kind, rest string, err error) {
	kind, err = tagKind(tag)
	if err != nil {
		return "", "", err
	}
	return kind, tag[len(kind)+1:], nil
}

// ParseTag parses a tag into its kind and identifier
// components. It returns an error if the tag is malformed,
// or if expectKind is not empty and the kind is
// not as expected.
func ParseTag(tag, expectKind string) (kind, id string, err error) {
	kind, id, err = splitTag(tag)
	if err != nil {
		return "", "", invalidTagError(tag, expectKind)
	}
	if expectKind != "" && kind != expectKind {
		return "", "", invalidTagError(tag, expectKind)
	}
	switch kind {
	case UnitTagKind:
		id = unitTagSuffixToId(id)
	case MachineTagKind:
		id = machineTagSuffixToId(id)
	case RelationTagKind:
		id = relationTagSuffixToKey(id)
	}
	switch kind {
	case UnitTagKind:
		if !IsUnit(id) {
			return "", "", invalidTagError(tag, kind)
		}
	case MachineTagKind:
		if !IsMachine(id) {
			return "", "", invalidTagError(tag, kind)
		}
	case ServiceTagKind:
		if !IsService(id) {
			return "", "", invalidTagError(tag, kind)
		}
	case UserTagKind:
		if !IsUser(id) {
			return "", "", invalidTagError(tag, kind)
		}
	case EnvironTagKind:
		if !IsEnvironment(id) {
			return "", "", invalidTagError(tag, kind)
		}
	case RelationTagKind:
		if !IsRelation(id) {
			return "", "", invalidTagError(tag, kind)
		}
	case NetworkTagKind:
		if !IsNetwork(id) {
			return "", "", invalidTagError(tag, kind)
		}
	}
	return kind, id, nil
}

func invalidTagError(tag, kind string) error {
	if kind != "" {
		return fmt.Errorf("%q is not a valid %s tag", tag, kind)
	}
	return fmt.Errorf("%q is not a valid tag", tag)
}
