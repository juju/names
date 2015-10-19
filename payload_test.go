// Copyright 2013 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/names"
)

var _ = gc.Suite(&payloadSuite{})

type payloadSuite struct{}

type payloadTest struct {
	input string
	id    string
	class string
	rawID string
}

func (t payloadTest) check(c *gc.C, tag names.PayloadTag) {
	c.Check(tag.Id(), gc.Equals, t.id)
	c.Check(tag.String(), gc.Equals, names.PayloadTagKind+"-"+t.id)
	c.Check(tag.Class(), gc.Equals, t.class)
	c.Check(tag.RawID(), gc.Equals, t.rawID)
}

func (s *payloadSuite) TestPayloadTag(c *gc.C) {
	for i, test := range []payloadTest{
		{
			input: "spam/eggs",
			id:    "spam/ZWdncw==",
			class: "spam",
			rawID: "eggs",
		}, {
			input: "spam/spam-eggs-and-spam",
			id:    "spam/c3BhbS1lZ2dzLWFuZC1zcGFt",
			class: "spam",
			rawID: "spam-eggs-and-spam",
		}, {
			input: "spam/spam/spam/spam...",
			id:    "spam/c3BhbS9zcGFtL3NwYW0uLi4=",
			class: "spam",
			rawID: "spam/spam/spam...",
		}, {
			input: "spam/f47ac10b-58cc-4372-a567-0e02b2c3d479",
			id:    "spam/ZjQ3YWMxMGItNThjYy00MzcyLWE1NjctMGUwMmIyYzNkNDc5",
			class: "spam",
			rawID: "f47ac10b-58cc-4372-a567-0e02b2c3d479",
		}, {
			input: "spam/3f9064e777bfd5ffc24553580f95111bb0ec82ed",
			id:    "spam/M2Y5MDY0ZTc3N2JmZDVmZmMyNDU1MzU4MGY5NTExMWJiMGVjODJlZA==",
			class: "spam",
			rawID: "3f9064e777bfd5ffc24553580f95111bb0ec82ed",
		},
	} {
		c.Logf("test %d: %s", i, test.input)
		tag := names.NewPayloadTag(test.class, test.rawID)
		parsed, err := names.ParsePayloadFullID(test.input)
		c.Assert(err, jc.ErrorIsNil)

		test.check(c, tag)
		test.check(c, parsed)
	}
}

func (s *payloadSuite) TestIsValidPayload(c *gc.C) {
	for i, test := range []struct {
		fullID string
		expect bool
	}{
		{"", false},
		{"spam", false},
		{"spam/", false},
		{"/eggs", false},

		{"spam/eggs", true},
		{"Spam/eggs", true},
		{"sPaM/eggs", true},
		{"SPAM/eggs", true},
		{"spam_spam_spam/eggs", true},
		{"spam-spam-spam/eggs", true},
		{"spam1/eggs", true},
		{"spam-/eggs", false},
		{"spam_/eggs", false},
		{"spam?/eggs", false},
		{"s.p.a.m/eggs", false},
		{"_spam_/eggs", false},
		{"__/eggs", false},
		{".../eggs", false},
		{"@!$#/eggs", false},

		{"spam/eggs", true},
		{"spam/Eggs", true},
		{"spam/eGgS", true},
		{"spam/EGGS", true},
		{"spam/eggs", true},
		{"spam/eggs?", true},
		{"spam/e.g.g.s", true},
		{"spam/_eggs_", true},
		{"spam/__", true},
		{"spam/...", true},
		{"spam/@!$#", true},

		{"spam/eggs/", true},
		{"spam/a/b/c", true},
		{"spam/spam/spam/spam", true},
		{"spam/_/_/_", true},
	} {
		c.Logf("test %d: %s", i, test.fullID)
		ok := names.IsValidPayload(test.fullID)

		c.Check(ok, gc.Equals, test.expect, gc.Commentf("%s", test.fullID))
	}
}

func (s *payloadSuite) TestParsePayloadFullIDOkay(c *gc.C) {
	for i, fullID := range []string{
		"spam/eggs",
		"Spam/eggs",
		"sPaM/eggs",
		"SPAM/eggs",
		"spam_spam_spam/eggs",
		"spam-spam-spam/eggs",
		"spam1/eggs",

		"spam/eggs",
		"spam/Eggs",
		"spam/eGgS",
		"spam/EGGS",
		"spam/eggs?",
		"spam/e.g.g.s",
		"spam/_eggs_",
		"spam/__",
		"spam/...",
		"spam/@!$#",

		"spam/eggs/",
		"spam/a/b/c",
		"spam/spam/spam/spam",
		"spam/_/_/_",
	} {
		c.Logf("test %d: %s", i, fullID)
		tag, err := names.ParsePayloadFullID(fullID)
		c.Assert(err, jc.ErrorIsNil)

		c.Check(tag.FullID(), gc.Equals, fullID)
	}
}

func (s *payloadSuite) TestParsePayloadFullIDInvalid(c *gc.C) {
	for i, fullID := range []string{
		"",
		"spam",
		"spam/",
		"/eggs",

		"spam-/eggs",
		"spam_/eggs",
		"spam?/eggs",
		"s.p.a.m/eggs",
		"_spam_/eggs",
		"__/eggs",
		".../eggs",
		"@!$#/eggs",
	} {
		c.Logf("test %d: %s", i, fullID)
		_, err := names.ParsePayloadFullID(fullID)

		c.Check(err, gc.ErrorMatches, `invalid payload ID .*`)
	}
}

func (s *payloadSuite) TestParsePayloadTag(c *gc.C) {
	for i, test := range []struct {
		tag      string
		expected names.Tag
		err      error
	}{{
		tag: "",
		err: names.InvalidTagError("", ""),
	}, {
		tag: "payload-",
		err: names.InvalidTagError("payload-", names.PayloadTagKind),
	}, {
		tag: "payload-/",
		err: names.InvalidTagError("payload-/", names.PayloadTagKind),
	}, {
		tag: "payload-spam/",
		err: names.InvalidTagError("payload-spam/", names.PayloadTagKind),
	}, {
		tag: "payload-/eggs",
		err: names.InvalidTagError("payload-/eggs", names.PayloadTagKind),
	}, {
		tag:      "payload-spam/ZWdncw==",
		expected: names.NewPayloadTag("spam", "eggs"),
	}, {
		tag: "spam/eggs",
		err: names.InvalidTagError("spam/eggs", ""),
	}, {
		tag: "unit-spam/eggs",
		err: names.InvalidTagError("unit-spam/eggs", names.UnitTagKind),
	}, {
		tag: "service-spam",
		err: names.InvalidTagError("service-spam", names.PayloadTagKind),
	}} {
		c.Logf("test %d: %s", i, test.tag)
		got, err := names.ParsePayloadTag(test.tag)
		if test.err != nil {
			c.Check(err, jc.DeepEquals, test.err)
		} else {
			c.Check(err, jc.ErrorIsNil)
			c.Check(got, jc.DeepEquals, test.expected)
		}
	}
}
