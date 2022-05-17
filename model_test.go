// Copyright 2016 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	"github.com/juju/names/v4"
	gc "gopkg.in/check.v1"
)

type modelSuite struct{}

var _ = gc.Suite(&modelSuite{})

var parseModelTagTests = []struct {
	tagString       string
	expectedTag     names.Tag
	expectedId      string
	expectedShortId string
	err             error
}{{
	tagString: "",
	err:       names.InvalidTagError("", ""),
}, {
	tagString:       "model-f47ac10b-58cc-4372-a567-0e02b2c3d479",
	expectedTag:     names.NewModelTag("f47ac10b-58cc-4372-a567-0e02b2c3d479"),
	expectedId:      "f47ac10b-58cc-4372-a567-0e02b2c3d479",
	expectedShortId: "f47ac1",
}, {
	tagString: "dave",
	err:       names.InvalidTagError("dave", ""),
}, {
	tagString: "model-",
	err:       names.InvalidTagError("model-", names.ModelTagKind),
}, {
	tagString: "application-dave",
	err:       names.InvalidTagError("application-dave", names.ModelTagKind),
}}

func (s *modelSuite) TestParseModelTag(c *gc.C) {
	for i, t := range parseModelTagTests {
		c.Logf("test %d: %s", i, t.tagString)
		got, err := names.ParseModelTag(t.tagString)
		if err != nil || t.err != nil {
			c.Check(err, gc.DeepEquals, t.err)
			continue
		}
		c.Check(got, gc.FitsTypeOf, t.expectedTag)
		c.Check(got, gc.Equals, t.expectedTag)
		c.Check(got.String(), gc.Equals, t.tagString)
		c.Check(got.Kind(), gc.Equals, "model")
		c.Check(got.Id(), gc.Equals, t.expectedId)
		c.Check(got.ShortId(), gc.Equals, t.expectedShortId)
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
