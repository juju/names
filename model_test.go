// Copyright 2016 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	gc "gopkg.in/check.v1"

	"gopkg.in/juju/names.v2"
)

type modelSuite struct{}

var _ = gc.Suite(&modelSuite{})

var parseModelTagTests = []struct {
	tag      string
	expected names.Tag
	err      error
}{{
	tag: "",
	err: names.InvalidTagError("", ""),
}, {
	tag:      "model-f47ac10b-58cc-4372-a567-0e02b2c3d479",
	expected: names.NewModelTag("f47ac10b-58cc-4372-a567-0e02b2c3d479"),
}, {
	tag: "dave",
	err: names.InvalidTagError("dave", ""),
}, {
	tag: "model-",
	err: names.InvalidTagError("model-", names.ModelTagKind),
}, {
	tag: "application-dave",
	err: names.InvalidTagError("application-dave", names.ModelTagKind),
}}

func (s *modelSuite) TestParseModelTag(c *gc.C) {
	for i, t := range parseModelTagTests {
		c.Logf("test %d: %s", i, t.tag)
		got, err := names.ParseModelTag(t.tag)
		if err != nil || t.err != nil {
			c.Check(err, gc.DeepEquals, t.err)
			continue
		}
		c.Check(got, gc.FitsTypeOf, t.expected)
		c.Check(got, gc.Equals, t.expected)
	}
}

var modelNameTest = []struct {
	test     string
	name     string
	expected bool
}{{
	test:     "Hyphenated true",
	name:     "foo-bar",
	expected: true,
}, {
	test:     "Whitespsce false",
	name:     "foo bar",
	expected: false,
}, {
	test:     "Capital false",
	name:     "fooBar",
	expected: false,
}, {
	test:     "At sign false",
	name:     "foo@bar",
	expected: false,
}}

func (s *modelSuite) TestModelName(c *gc.C) {
	for i, t := range modelNameTest {
		c.Logf("test %d: %q", i, t.name)
		c.Check(names.IsValidModelName(t.name), gc.Equals, t.expected)
	}
}
