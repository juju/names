// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	"fmt"

	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"gopkg.in/juju/names.v2"
)

type charmSuite struct{}

var _ = gc.Suite(&charmSuite{})

var validCharmURLs = []string{"charm",
	// V1 charm urls.
	"local:charm",
	"local:charm--1",
	"local:charm-1",
	"local:series/charm",
	"local:series/charm-3",
	"local:series/charm-0",
	"cs:~user/charm",
	"cs:~user/charm-1",
	"cs:~user/series/charm",
	"cs:~user/series/charm-1",
	"cs:series/charm",
	"cs:series/charm-with-long-name",
	"cs:series/charm-3",
	"cs:series/charm-0",
	"cs:charm",
	"cs:charm--1",
	"cs:charm-1",
	"charm",
	"charm-1",
	"series/charm",
	"series/charm-1",

	// V3 charm urls.
	"local:charm-with-long2-name/series/2",
	"local:charm-with-long2-name/series",
	"local:charm-with-long2-name/2",
	"cs:user/charm-with-long-name/series/2",
	"cs:charm-with-long2-name/series/2",
	"cs:user/charm-with-long-name/2",
	"cs:user/charm-with-long-name/series",
	"cs:charm-with-long-name/2",
	"cs:charm-with-long-name/series",
	"cs:user/charm-with-long-name",
	"user/charm-with-long-name/series/2",
	"charm-with-long2-name/series/2",
	"user/charm-with-long-name/2",
	"user/charm-with-long-name/series",
	"charm-with-long-name/2",
	"charm-with-long-name/series",
	"user/charm-with-long-name",
}

func (s *charmSuite) TestValidCharmURLs(c *gc.C) {
	for _, url := range validCharmURLs {
		c.Logf("Processing tag %q", url)
		c.Assert(names.IsValidCharm(url), jc.IsTrue)
	}
}

func (s *charmSuite) TestInvalidCharmURLs(c *gc.C) {
	invalidURLs := []string{"",
		"local:~user/charm",          // false: user on local
		"local:~user/series/charm",   // false: user on local
		"local:~user/series/charm-1", // false: user on local
		"local:charm--2",             // false: only -1 is a valid negative revision
		"blah:charm-2",               // false: invalid schema
		"local:series/charm-01",      // false: revision is funny
		"local:user/name/series/2",   // false: local charms can't have users
	}
	for _, url := range invalidURLs {
		c.Logf("Processing tag %q", url)
		c.Assert(names.IsValidCharm(url), jc.IsFalse)
	}
}

func (s *charmSuite) TestParseCharmTagValid(c *gc.C) {
	for _, tag := range validCharmURLs {
		c.Logf("Processing tag %q", tag)
		s.assertParseCharmTagValid(c, fmt.Sprintf("charm-%v", tag), names.NewCharmTag(tag))
	}
}

func (s *charmSuite) TestParseCharmTagInvalid(c *gc.C) {
	invalidTags := []string{"",
		"blah",
		"charm",
		"user-blah",
	}
	for _, aTag := range invalidTags {
		c.Logf("Processing tag %q", aTag)
		s.assertParseCharmTagInvalid(c, aTag)
	}
}

func (s *charmSuite) assertParseCharmTagValid(c *gc.C, tag string, expected names.Tag) {
	got, err := names.ParseCharmTag(tag)
	c.Assert(err, jc.ErrorIsNil)
	c.Check(got, gc.FitsTypeOf, expected)
	c.Check(got, gc.Equals, expected)
}

func (s *charmSuite) assertParseCharmTagInvalid(c *gc.C, tag string) {
	_, err := names.ParseCharmTag(tag)
	c.Check(err, gc.ErrorMatches, fmt.Sprintf(".*%q is not a valid.*", tag))
}
