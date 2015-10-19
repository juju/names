// Copyright 2013 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
)

const (
	// PayloadTagKind is used as the prefix for the string
	// representation of payload tags.
	PayloadTagKind = "payload"

	validPayloadClass = `(?:[a-zA-Z](?:[-\w]*[a-zA-Z0-9])?)`
)

var (
	// TODO(ericsnow) Should we require that the class string be
	// a valid identifier ("[a-zA-Z]?
	validPayload = regexp.MustCompile(fmt.Sprintf(`^(?P<class>%s)/(?P<rawid>.+)$`, validPayloadClass))
)

// IsValidPayload returns whether fullID is a valid Juju ID for
// a charm payload. Examples of valid payload IDs include
// spam/eggs, spam/spam-eggs-and-spam, and spam/spam/spam/spam...
func IsValidPayload(fullID string) bool {
	return validPayload.MatchString(fullID)
}

// PayloadTag represents a charm payload.
type PayloadTag struct {
	// class is the name of the charm-defined payload class.
	class string
	// rawID uniquely identifies the payload to the underlying
	// technology of the payload's type.
	rawID string
	// encodedID is an encoded copy of rawID. The value is
	// base64-encoded and then "/" is replaced with ".".
	encodedID string
}

// NewPayloadTag returns the tag for a charm's payload with the given
// payload class and raw (from underlying tech, e.g. docker) ID.
func NewPayloadTag(class, rawID string) PayloadTag {
	encoded := encodePayloadRawID(rawID)
	return newPayloadTag(class, rawID, encoded)
}

func newPayloadTag(class, rawID, encodedID string) PayloadTag {
	return PayloadTag{
		class:     class,
		rawID:     rawID,
		encodedID: encodedID,
	}
}

// ParsePayloadFullID parses the given Juju-recognized ID for a charm
// payload and returns the corresponding PayloadTag.
func ParsePayloadFullID(fullID string) (PayloadTag, error) {
	class, rawID, err := parsePayloadFullID(fullID)
	if err != nil {
		return PayloadTag{}, err
	}
	return NewPayloadTag(class, rawID), nil
}

func parsePayloadFullIDEncoded(encoded string) (PayloadTag, error) {
	class, encodedID, err := parsePayloadFullID(encoded)
	if err != nil {
		return PayloadTag{}, err
	}
	rawID := decodePayloadRawID(encodedID)
	return newPayloadTag(class, rawID, encodedID), nil
}

func parsePayloadFullID(fullID string) (string, string, error) {
	parts := validPayload.FindStringSubmatch(fullID)
	if len(parts) != 3 {
		return "", "", fmt.Errorf("invalid payload ID %q", fullID)
	}
	return parts[1], parts[2], nil
}

// ParsePayloadTag parses a payload tag string.
// So ParsePayloadTag(tag.String()) === tag.
func ParsePayloadTag(tag string) (PayloadTag, error) {
	t, err := ParseTag(tag)
	if err != nil {
		return PayloadTag{}, err
	}
	pt, ok := t.(PayloadTag)
	if !ok {
		return PayloadTag{}, invalidTagError(tag, PayloadTagKind)
	}
	return pt, nil
}

// Kind implements Tag.
func (t PayloadTag) Kind() string {
	return PayloadTagKind
}

// Id implements Tag.Id. It always returns the same ID with which
// it was created. So NewPayloadTag(x).Id() == x for all valid x.
func (t PayloadTag) Id() string {
	return t.class + "/" + t.encodedID
}

// String implements Tag.
func (t PayloadTag) String() string {
	return tagString(t)
}

// Class identifies the payload's class, relative
// to the payload's charm.
func (t PayloadTag) Class() string {
	return t.class
}

// TODO(ericsnow) Find a better name than "RawID"?

// RawID uniquely identifies the payload to the underlying technology
// for the payload's type.
func (t PayloadTag) RawID() string {
	return t.rawID
}

// FullID returns a string represenatation of the tag
// that round-trips with ParsePayloadFullID().
func (t PayloadTag) FullID() string {
	return t.class + "/" + t.rawID
}

func encodePayloadRawID(rawID string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(rawID))
	// We try to keep the ID filename-friendly.
	return strings.Replace(encoded, "/", ".", -1)
}

func decodePayloadRawID(encoded string) string {
	encoded = strings.Replace(encoded, ".", "/", -1)
	// We can ignore the error return since we control the encoding.
	data, _ := base64.StdEncoding.DecodeString(encoded)
	return string(data)
}
