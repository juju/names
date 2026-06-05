// Copyright 2015 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// uuid represents a universally unique identifier with 16 octets.
// This is a package-local type used to avoid a dependency on
// github.com/juju/utils for UUID parsing and validation.
type uuid [16]byte

// String returns the canonical hyphenated hex representation of the UUID.
func (u uuid) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:16])
}

// validUUIDFull is an anchored regexp for full-string UUID validation.
// It reuses the UUIDv7Snippet constant defined in tag.go.
var validUUIDFull = regexp.MustCompile("^" + UUIDv7Snippet + "$")

// isValidUUIDString returns true if s is a valid UUID string.
func isValidUUIDString(s string) bool {
	return validUUIDFull.MatchString(s)
}

// uuidFromString parses a UUID from its canonical string representation.
func uuidFromString(s string) (uuid, error) {
	if !isValidUUIDString(s) {
		return uuid{}, fmt.Errorf("invalid UUID: %q", s)
	}
	s = strings.Replace(s, "-", "", 4)
	raw, err := hex.DecodeString(s)
	if err != nil {
		return uuid{}, err
	}
	var u uuid
	copy(u[:], raw)
	return u, nil
}

// newUUID generates a new random (version 4) UUID.
func newUUID() (uuid, error) {
	var u uuid
	if _, err := io.ReadFull(rand.Reader, u[:]); err != nil {
		return uuid{}, err
	}
	// Set version (4) and variant (2) per RFC 4122.
	u[6] = (4 << 4) | (u[6] & 0x0f)
	u[8] = (8 << 4) | (u[8] & 0x0f)
	return u, nil
}
