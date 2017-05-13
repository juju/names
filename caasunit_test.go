// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	"fmt"

	gc "gopkg.in/check.v1"

	"gopkg.in/juju/names.v2"
)

type caasUnitSuite struct{}

var _ = gc.Suite(&caasUnitSuite{})

func (s *caasUnitSuite) TestCAASUnitTag(c *gc.C) {
	c.Assert(names.NewCAASUnitTag("wordpress/2").String(), gc.Equals, "caasunit-wordpress-2")
}

var caasUnitNameTests = []struct {
	pattern     string
	valid       bool
	application string
}{
	{pattern: "wordpress/42", valid: true, application: "wordpress"},
	{pattern: "rabbitmq-server/123", valid: true, application: "rabbitmq-server"},
	{pattern: "foo", valid: false},
	{pattern: "foo/", valid: false},
	{pattern: "bar/foo", valid: false},
	{pattern: "20/20", valid: false},
	{pattern: "foo-55", valid: false},
	{pattern: "foo-bar/123", valid: true, application: "foo-bar"},
	{pattern: "foo-bar/123/", valid: false},
	{pattern: "foo-bar/123-not", valid: false},
}

func (s *caasUnitSuite) TestCAASUnitNameFormats(c *gc.C) {
	for i, test := range caasUnitNameTests {
		c.Logf("test %d: %q", i, test.pattern)
		c.Assert(names.IsValidCAASUnit(test.pattern), gc.Equals, test.valid)
	}
}

func (s *caasUnitSuite) TestInvalidCAASUnitTagFormats(c *gc.C) {
	for i, test := range caasUnitNameTests {
		if !test.valid {
			c.Logf("test %d: %q", i, test.pattern)
			expect := fmt.Sprintf("%q is not a valid caasunit name", test.pattern)
			testCAASUnitTag := func() { names.NewCAASUnitTag(test.pattern) }
			c.Assert(testCAASUnitTag, gc.PanicMatches, expect)
		}
	}
}

func (s *applicationSuite) TestCAASUnitApplication(c *gc.C) {
	for i, test := range caasUnitNameTests {
		c.Logf("test %d: %q", i, test.pattern)
		if !test.valid {
			expect := fmt.Sprintf("%q is not a valid caasunit name", test.pattern)
			_, err := names.CAASUnitApplication(test.pattern)
			c.Assert(err, gc.ErrorMatches, expect)
		} else {
			result, err := names.CAASUnitApplication(test.pattern)
			c.Assert(err, gc.IsNil)
			c.Assert(result, gc.Equals, test.application)
		}
	}
}

var parseCAASUnitTagTests = []struct {
	tag      string
	expected names.Tag
	err      error
}{{
	tag: "",
	err: names.InvalidTagError("", ""),
}, {
	tag:      "caasunit-dave/0",
	expected: names.NewCAASUnitTag("dave/0"),
}, {
	tag: "dave",
	err: names.InvalidTagError("dave", ""),
}, {
	tag: "caasunit-dave",
	err: names.InvalidTagError("caasunit-dave", names.CAASUnitTagKind), // not a valid unit name either
}, {
	tag: "application-dave",
	err: names.InvalidTagError("application-dave", names.CAASUnitTagKind),
}}

func (s *caasUnitSuite) TestParseCAASUnitTag(c *gc.C) {
	for i, t := range parseCAASUnitTagTests {
		c.Logf("test %d: %s", i, t.tag)
		got, err := names.ParseCAASUnitTag(t.tag)
		if err != nil || t.err != nil {
			c.Check(err, gc.DeepEquals, t.err)
			continue
		}
		c.Check(got, gc.FitsTypeOf, t.expected)
		c.Check(got, gc.Equals, t.expected)
	}
}
