// Copyright 2016 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	gc "gopkg.in/check.v1"

	"gopkg.in/juju/names.v2"
)

type controllerSuite struct{}

var _ = gc.Suite(&controllerSuite{})

var parseControllerTagTests = []struct {
	title    string
	tag      string
	expected names.Tag
	err      error
}{{
	title: "empty tag fails",
	tag:   "",
	err:   names.InvalidTagError("", ""),
}, {
	title:    "valid controller-<uuid> tag",
	tag:      "controller-f47ac10b-58cc-4372-a567-0e02b2c3d479",
	expected: names.NewControllerTag("f47ac10b-58cc-4372-a567-0e02b2c3d479"),
}, {

	title: "invalid controller tag one word",
	tag:   "dave",
	err:   names.InvalidTagError("dave", ""),
}, {
	title: "invalid controller tag prefix only",
	tag:   "controller-",
	err:   names.InvalidTagError("controller-", names.ControllerTagKind),
}, {
	title: "invalid controller tag hyphen separated words",
	tag:   "application-dave",
	err:   names.InvalidTagError("application-dave", names.ControllerTagKind),
}, {
	title: "invalid controller tag non hyphen separated prefix",
	tag:   "controllerf47ac10b-58cc-4372-a567-0e02b2c3d479",
	err:   names.InvalidTagError("controllerf47ac10b-58cc-4372-a567-0e02b2c3d479", ""),
}, {
	title: "invalid controller tag non hyphen separated terms",
	tag:   "controllerf47ac10b58cc4372a5670e02b2c3d479",
	err:   names.InvalidTagError("controllerf47ac10b58cc4372a5670e02b2c3d479", ""),
}}

func (s *controllerSuite) TestParseControllerTag(c *gc.C) {
	for i, t := range parseControllerTagTests {
		c.Logf("test %d: %q %s", i, t.title, t.tag)
		got, err := names.ParseControllerTag(t.tag)
		if err != nil || t.err != nil {
			c.Check(err, gc.DeepEquals, t.err)
			continue
		}
		c.Check(got, gc.FitsTypeOf, t.expected)
		c.Check(got, gc.Equals, t.expected)
	}
}

var controllerNameTest = []struct {
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

func (s *controllerSuite) TestControllerName(c *gc.C) {
	for i, t := range controllerNameTest {
		c.Logf("test %d: %q", i, t.name)
		c.Check(names.IsValidControllerName(t.name), gc.Equals, t.expected)
	}
}
