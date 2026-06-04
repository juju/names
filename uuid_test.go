// Copyright 2026 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"strings"

	gc "gopkg.in/check.v1"
)

type uuidSuite struct{}

var _ = gc.Suite(&uuidSuite{})

// isValidUUIDString tests

var isValidUUIDStringTests = []struct {
	description string
	input       string
	expected    bool
}{
	// Valid UUIDs
	{
		description: "canonical v4 UUID",
		input:       "f47ac10b-58cc-4372-a567-0e02b2c3d479",
		expected:    true,
	},
	{
		description: "all-zeros UUID",
		input:       "00000000-0000-0000-0000-000000000000",
		expected:    true,
	},
	{
		description: "all-f UUID",
		input:       "ffffffff-ffff-ffff-ffff-ffffffffffff",
		expected:    true,
	},
	{
		description: "mixed digits and letters",
		input:       "deadbeef-1234-5678-90ab-cdef01234567",
		expected:    true,
	},
	// Invalid: wrong structure
	{
		description: "empty string",
		input:       "",
		expected:    false,
	},
	{
		description: "too short",
		input:       "f47ac10b-58cc-4372-a567",
		expected:    false,
	},
	{
		description: "too long",
		input:       "f47ac10b-58cc-4372-a567-0e02b2c3d479-extra",
		expected:    false,
	},
	{
		description: "no hyphens",
		input:       "f47ac10b58cc4372a5670e02b2c3d479",
		expected:    false,
	},
	{
		description: "wrong hyphen positions",
		input:       "f47ac10b-58cc4372-a567-0e02b2c3d479",
		expected:    false,
	},
	// Invalid: uppercase letters (UUIDv7Snippet only accepts [0-9a-f])
	{
		description: "uppercase hex digits",
		input:       "F47AC10B-58CC-4372-A567-0E02B2C3D479",
		expected:    false,
	},
	// Invalid: non-hex characters
	{
		description: "non-hex character",
		input:       "g47ac10b-58cc-4372-a567-0e02b2c3d479",
		expected:    false,
	},
	// Invalid: anchoring — valid UUID surrounded by extra text
	{
		description: "valid UUID with leading text",
		input:       "xf47ac10b-58cc-4372-a567-0e02b2c3d479",
		expected:    false,
	},
	{
		description: "valid UUID with trailing text",
		input:       "f47ac10b-58cc-4372-a567-0e02b2c3d479x",
		expected:    false,
	},
}

func (s *uuidSuite) TestIsValidUUIDString(c *gc.C) {
	for _, tt := range isValidUUIDStringTests {
		c.Logf("test: %s", tt.description)
		c.Check(isValidUUIDString(tt.input), gc.Equals, tt.expected)
	}
}

// uuidFromString tests

func (s *uuidSuite) TestUUIDFromStringValid(c *gc.C) {
	input := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	u, err := uuidFromString(input)
	c.Assert(err, gc.IsNil)
	c.Check(u.String(), gc.Equals, input)
}

func (s *uuidSuite) TestUUIDFromStringAllZeros(c *gc.C) {
	input := "00000000-0000-0000-0000-000000000000"
	u, err := uuidFromString(input)
	c.Assert(err, gc.IsNil)
	c.Check(u.String(), gc.Equals, input)
	// All bytes should be zero.
	for i, b := range u {
		c.Check(b, gc.Equals, byte(0), gc.Commentf("byte %d", i))
	}
}

func (s *uuidSuite) TestUUIDFromStringByteLayout(c *gc.C) {
	// Use a UUID where each block is distinguishable.
	// 11223344-5566-7788-99aa-bbccddeeff00
	input := "11223344-5566-7788-99aa-bbccddeeff00"
	u, err := uuidFromString(input)
	c.Assert(err, gc.IsNil)
	expected := [16]byte{
		0x11, 0x22, 0x33, 0x44, // block1
		0x55, 0x66, // block2
		0x77, 0x88, // block3
		0x99, 0xaa, // block4
		0xbb, 0xcc, 0xdd, 0xee, 0xff, 0x00, // block5
	}
	c.Check([16]byte(u), gc.DeepEquals, expected)
}

func (s *uuidSuite) TestUUIDFromStringInvalid(c *gc.C) {
	invalid := []string{
		"",
		"not-a-uuid",
		"42",
		"f47ac10b-58cc-4372-a567", // truncated
	}
	for _, input := range invalid {
		_, err := uuidFromString(input)
		c.Check(err, gc.NotNil, gc.Commentf("input: %q", input))
		c.Check(err.Error(), gc.Equals, `invalid UUID: "`+input+`"`)
	}
}

// uuid.String tests

func (s *uuidSuite) TestUUIDString(c *gc.C) {
	// Manually construct a uuid and verify String() format.
	u := uuid{
		0x12, 0x34, 0x56, 0x78,
		0x9a, 0xbc,
		0xde, 0xf0,
		0x12, 0x34,
		0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
	}
	c.Check(u.String(), gc.Equals, "12345678-9abc-def0-1234-56789abcdef0")
}

func (s *uuidSuite) TestUUIDStringRoundtrip(c *gc.C) {
	// Parse then re-serialise should be identity.
	inputs := []string{
		"f47ac10b-58cc-4372-a567-0e02b2c3d479",
		"00000000-0000-0000-0000-000000000000",
		"ffffffff-ffff-ffff-ffff-ffffffffffff",
	}
	for _, input := range inputs {
		u, err := uuidFromString(input)
		c.Assert(err, gc.IsNil)
		c.Check(u.String(), gc.Equals, input)
	}
}

// newUUID tests

func (s *uuidSuite) TestNewUUIDVersion4(c *gc.C) {
	u, err := newUUID()
	c.Assert(err, gc.IsNil)
	// Version is encoded in the high nibble of byte 6 and must be 4.
	c.Check(u[6]>>4, gc.Equals, byte(4))
}

func (s *uuidSuite) TestNewUUIDVariant(c *gc.C) {
	u, err := newUUID()
	c.Assert(err, gc.IsNil)
	// RFC 4122 variant: high two bits of byte 8 must be 10xx (i.e. 0x8x–0xbx).
	c.Check(u[8]&0xc0, gc.Equals, byte(0x80))
}

func (s *uuidSuite) TestNewUUIDIsValid(c *gc.C) {
	u, err := newUUID()
	c.Assert(err, gc.IsNil)
	c.Check(isValidUUIDString(u.String()), gc.Equals, true)
}

func (s *uuidSuite) TestNewUUIDUnique(c *gc.C) {
	const n = 10
	seen := make(map[string]bool, n)
	for i := 0; i < n; i++ {
		u, err := newUUID()
		c.Assert(err, gc.IsNil)
		s := u.String()
		c.Assert(seen[s], gc.Equals, false, gc.Commentf("duplicate UUID %q", s))
		seen[s] = true
	}
}

func (s *uuidSuite) TestNewUUIDStringFormat(c *gc.C) {
	u, err := newUUID()
	c.Assert(err, gc.IsNil)
	str := u.String()
	// Must be exactly 36 characters: 8-4-4-4-12.
	c.Check(len(str), gc.Equals, 36)
	parts := strings.Split(str, "-")
	c.Assert(len(parts), gc.Equals, 5)
	c.Check(len(parts[0]), gc.Equals, 8)
	c.Check(len(parts[1]), gc.Equals, 4)
	c.Check(len(parts[2]), gc.Equals, 4)
	c.Check(len(parts[3]), gc.Equals, 4)
	c.Check(len(parts[4]), gc.Equals, 12)
}
