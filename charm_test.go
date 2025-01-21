// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	"fmt"

	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/names/v5"
)

type charmSuite struct{}

var _ = gc.Suite(&charmSuite{})

var validCharmURLs = []string{"charm",
	"charm",
	"charm-1",
	"amd64/charm",
	"amd64/charm-42",
	"seriesorarch/charm",
	"series/charm-1",
	"arch/series/charm",
	"arch/series/charm-1",
	"local:arch/series/charm-with-long2-name",
	"local:series/charm-with-long2-name-2",
	"local:arch/charm-with-long2-name-2",
	"ch:series/charm-with-long2-name-2",
	"ch:charm-with-long-name",
	"ch:charm-with-long-name-2",
	"ch:seriesorarch/charm-with-long-name",
	"ch:arch/series/charm-with-long-name",
	"ch:arch/series/charm-with-long-name-2",
}

func (s *charmSuite) TestValidCharmURLs(c *gc.C) {
	for _, url := range validCharmURLs {
		c.Logf("Processing tag %q", url)
		c.Assert(names.IsValidCharm(url), jc.IsTrue)
	}
}

func (s *charmSuite) TestInvalidCharmURLs(c *gc.C) {
	invalidURLs := []string{"",
		"local:~user/charm",              // false: user on local
		"local:~user/series/charm",       // false: user on local
		"local:~user/series/charm-1",     // false: user on local
		"ch:~user/series/charm-1",        // false: user on charmhub
		"local:charm--2",                 // false: only -1 is a valid negative revision
		"blah:charm-2",                   // false: invalid schema
		"local:series/charm-01",          // false: revision is funny
		"local:user/name/series/2",       // false: local charms can't have users
		"wrongschema:arch/series/name-1", // false: wrong schema
		"part0/part1/part2/part3/part4",  // false: too many parts
		"ch:part0/part1/part2/part3",     // false: too many parts, with schema
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

func (s *charmSuite) TestCharmFields(c *gc.C) {
	for _, test := range []struct {
		tag    string
		source string
		arch   string
		series string
		name   string
		rev    int
	}{
		{"amd64/mysql-42", "charmhub", "amd64", "", "mysql", 42},
		{"ch:amd64/mysql-42", "charmhub", "amd64", "", "mysql", 42},
		{"ch:amd64/focal/mysql-42", "charmhub", "amd64", "focal", "mysql", 42},
		{"ch:focal/mysql-42", "charmhub", "", "focal", "mysql", 42},
		{"ch:mysql-42", "charmhub", "", "", "mysql", 42},
		{"ch:mysql", "charmhub", "", "", "mysql", -1},
		{"local:trusty/mysql-42", "local", "", "trusty", "mysql", 42},
		{"local:mysql-42", "local", "", "", "mysql", 42},
	} {
		c.Logf("Processing tag %q", test.tag)
		tag, err := names.ParseCharmTag(names.CharmTagKind + "-" + test.tag)
		c.Assert(err, jc.ErrorIsNil)
		c.Check(tag.Source(), gc.Equals, test.source)
		c.Check(tag.Architecture(), gc.Equals, test.arch)
		c.Check(tag.Series(), gc.Equals, test.series)
		c.Check(tag.Name(), gc.Equals, test.name)
		c.Check(tag.Revision(), gc.Equals, test.rev)
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
