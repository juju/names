// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package names

import (
	"strings"
)

const EnvironTagKind = "environment"

type environTag struct {
	uuid string
}

// EnvironTag returns the tag of an environment with the given environment UUID.
func EnvironTag(uuid string) Tag {
	return environTag{uuid: uuid}
}

func (t environTag) String() string { return EnvironTagKind + "-" + t.uuid }

// IsEnvironment returns whether id is a valid environment UUID.
func IsEnvironment(id string) bool {
	// TODO(axw) 2013-12-04 #1257587
	// We should not accept environment tags that
	// do not look like UUIDs.
	return !strings.Contains(id, "/")
}
