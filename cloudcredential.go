// Copyright 2016 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"fmt"
	"regexp"
	"strings"
)

const CloudCredentialTagKind = "cloudcred"

var (
	cloudCredentialNameSnippet = "[a-zA-Z][a-zA-Z0-9.-]*"
	validCloudCredentialName   = regexp.MustCompile("^" + cloudCredentialNameSnippet + "$")
	validCloudCredential       = regexp.MustCompile(
		"^" +
			"(" + cloudSnippet + ")" +
			"/(" + validUserSnippet + ")" + // credential owner
			"/(" + cloudCredentialNameSnippet + ")" +
			"$",
	)
)

type CloudCredentialTag struct {
	cloud CloudTag
	owner UserTag
	name  string
}

// Kind is part of the Tag interface.
func (t CloudCredentialTag) Kind() string { return CloudCredentialTagKind }

// Id is part of the Tag interface.
func (t CloudCredentialTag) Id() string {
	return t.id(false)
}

// String is part of the Tag interface.
func (t CloudCredentialTag) String() string {
	return fmt.Sprintf("%s-%s-%s-%s", t.Kind(), t.cloud.Id(), t.owner.Id(), t.name)
}

// Canonical returns the cloud credential ID in canonical form.
// Specifically, the user tag portion will be canonicalized.
func (t CloudCredentialTag) Canonical() string {
	return t.id(true)
}

func (t CloudCredentialTag) id(canonical bool) string {
	var ownerId string
	if canonical {
		ownerId = t.owner.Canonical()
	} else {
		ownerId = t.owner.Id()
	}
	return fmt.Sprintf("%s/%s/%s", t.cloud.Id(), ownerId, t.name)
}

// Cloud returns the tag of the cloud to which the credential pertains.
func (t CloudCredentialTag) Cloud() CloudTag {
	return t.cloud
}

// Owner returns the tag of the user that owns the credential.
func (t CloudCredentialTag) Owner() UserTag {
	return t.owner
}

// Name returns the cloud credential name, excluding the
// cloud and owner qualifiers.
func (t CloudCredentialTag) Name() string {
	return t.name
}

// NewCloudCredentialTag returns the tag for the cloud with the given ID.
// It will panic if the given cloud ID is not valid.
func NewCloudCredentialTag(id string) CloudCredentialTag {
	parts := validCloudCredential.FindStringSubmatch(id)
	if len(parts) != 4 {
		panic(fmt.Sprintf("%q is not a valid cloud credential ID", id))
	}
	cloud := NewCloudTag(parts[1])
	owner := NewUserTag(parts[2])
	return CloudCredentialTag{cloud, owner, parts[3]}
}

// ParseCloudCredentialTag parses a cloud tag string.
func ParseCloudCredentialTag(s string) (CloudCredentialTag, error) {
	tag, err := ParseTag(s)
	if err != nil {
		return CloudCredentialTag{}, err
	}
	dt, ok := tag.(CloudCredentialTag)
	if !ok {
		return CloudCredentialTag{}, invalidTagError(s, CloudCredentialTagKind)
	}
	return dt, nil
}

// IsValidCloudCredential returns whether id is a valid cloud credential ID.
func IsValidCloudCredential(id string) bool {
	return validCloudCredential.MatchString(id)
}

// IsValidCloudCredentialName returns whether name is a valid cloud credential name.
func IsValidCloudCredentialName(name string) bool {
	return validCloudCredentialName.MatchString(name)
}

func cloudCredentialTagSuffixToId(s string) string {
	return strings.Replace(s, "-", "/", -1)
}
