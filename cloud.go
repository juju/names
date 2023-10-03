// Copyright 2016 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"fmt"
	"regexp"
)

const CloudTagKind = "cloud"

var (
	cloudSnippet = "[a-zA-Z0-9][a-zA-Z0-9._-]*"
	validCloud   = regexp.MustCompile("^" + cloudSnippet + "$")
)

// CloudTag is a names.Tag the represents a Cloud in the Juju domain.
type CloudTag struct {
	id string
}

// String implements Tag.String.
func (t CloudTag) String() string { return t.Kind() + "-" + t.id }

// Kind implements Tag.Kind returning a CloudTag's unique kind value.
func (t CloudTag) Kind() string { return CloudTagKind }

// Id returns the id of this CloudTag. This is the name of the cloud as
// represented by Juju. Examples:
// aws
// azure
// kubernetes
func (t CloudTag) Id() string { return t.id }

// NewCloudTag returns the tag for the cloud with the given ID.
// It will panic if the given cloud ID is not valid.
func NewCloudTag(id string) CloudTag {
	if !IsValidCloud(id) {
		panic(fmt.Sprintf("%q is not a valid cloud ID", id))
	}
	return CloudTag{id}
}

// ParseCloudTag parses a cloud tag string.
func ParseCloudTag(cloudTag string) (CloudTag, error) {
	tag, err := ParseTag(cloudTag)
	if err != nil {
		return CloudTag{}, err
	}
	dt, ok := tag.(CloudTag)
	if !ok {
		return CloudTag{}, invalidTagError(cloudTag, CloudTagKind)
	}
	return dt, nil
}

// IsValidCloud returns whether id is a valid cloud ID.
func IsValidCloud(id string) bool {
	return validCloud.MatchString(id)
}
