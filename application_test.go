// Copyright 2013 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	gc "gopkg.in/check.v1"

	"gopkg.in/juju/names.v2"
)

type applicationSuite struct{}

var _ = gc.Suite(&applicationSuite{})

var applicationNameTests = []struct {
	pattern string
	valid   bool
}{
	{pattern: "", valid: false},
	{pattern: "wordpress", valid: true},
	{pattern: "foo42", valid: true},
	{pattern: "doing55in54", valid: true},
	{pattern: "%not", valid: false},
	{pattern: "42also-not", valid: false},
	{pattern: "but-this-works", valid: true},
	{pattern: "so-42-far-not-good", valid: false},
	{pattern: "foo/42", valid: false},
	{pattern: "is-it-", valid: false},
	{pattern: "broken2-", valid: false},
	{pattern: "foo2", valid: true},
	{pattern: "foo-2", valid: false},
}

func (s *applicationSuite) TestApplicationNameFormats(c *gc.C) {
	assertApplication := func(s string, expect bool) {
		c.Assert(names.IsValidApplication(s), gc.Equals, expect)
		// Check that anything that is considered a valid application name
		// is also (in)valid if a(n) (in)valid unit designator is added
		// to it.
		c.Assert(names.IsValidUnit(s+"/0"), gc.Equals, expect)
		c.Assert(names.IsValidUnit(s+"/99"), gc.Equals, expect)
		c.Assert(names.IsValidUnit(s+"/-1"), gc.Equals, false)
		c.Assert(names.IsValidUnit(s+"/blah"), gc.Equals, false)
		c.Assert(names.IsValidUnit(s+"/"), gc.Equals, false)
	}

	for i, test := range applicationNameTests {
		c.Logf("test %d: %q", i, test.pattern)
		assertApplication(test.pattern, test.valid)
	}
}

var parseApplicationTagTests = []struct {
	tag      string
	expected names.Tag
	err      error
}{{
	tag: "",
	err: names.InvalidTagError("", ""),
}, {
	tag:      "application-dave",
	expected: names.NewApplicationTag("dave"),
}, {
	tag: "dave",
	err: names.InvalidTagError("dave", ""),
}, {
	tag: "application-dave/0",
	err: names.InvalidTagError("application-dave/0", names.ApplicationTagKind),
}, {
	tag: "application",
	err: names.InvalidTagError("application", ""),
}, {
	tag: "user-dave",
	err: names.InvalidTagError("user-dave", names.ApplicationTagKind),
}}

func (s *applicationSuite) TestParseApplicationTag(c *gc.C) {
	for i, t := range parseApplicationTagTests {
		c.Logf("test %d: %s", i, t.tag)
		got, err := names.ParseApplicationTag(t.tag)
		if err != nil || t.err != nil {
			c.Check(err, gc.DeepEquals, t.err)
			continue
		}
		c.Check(got, gc.FitsTypeOf, t.expected)
		c.Check(got, gc.Equals, t.expected)
	}
}
