// Copyright 2016 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	gc "gopkg.in/check.v1"

	"gopkg.in/juju/names.v2"
)

type cloudSuite struct{}

var _ = gc.Suite(&cloudSuite{})

func (s *cloudSuite) TestCloudTag(c *gc.C) {
	for i, t := range []struct {
		input  string
		string string
	}{
		{
			input:  "bob",
			string: "cloud-bob",
		},
	} {
		c.Logf("test %d: %s", i, t.input)
		cloudTag := names.NewCloudTag(t.input)
		c.Check(cloudTag.String(), gc.Equals, t.string)
		c.Check(cloudTag.Id(), gc.Equals, t.input)
	}
}

func (s *cloudSuite) TestIsValidCloud(c *gc.C) {
	for i, t := range []struct {
		string string
		expect bool
	}{
		{"", false},
		{"bob", true},
		{"Bob", true},
		{"bOB", true},
		{"b^b", false},
		{"bob1", true},
		{"bob-1", true},
		{"bob.1", true},
		{"1bob", true},
		{"1-bob", true},
		{"1+bob", false},
		{"1.bob", true},
		{"a", true},
		{"0foo", true},
		{"foo bar", false},
		{"bar{}", false},
		{"bar+foo", false},
		{"bar_foo", true},
		{"bar_", true},
		{"bar!", false},
		{"bar^", false},
		{"bar*", false},
		{"foo=bar", false},
		{"foo?", false},
		{"[bar]", false},
		{"'foo'", false},
		{"%bar", false},
		{"&bar", false},
		{"#1foo", false},
		{"bar@", false},
		{"@local", false},
		{"not/valid", false},
	} {
		c.Logf("test %d: %s", i, t.string)
		c.Assert(names.IsValidCloud(t.string), gc.Equals, t.expect, gc.Commentf("%s", t.string))
	}
}

func (s *cloudSuite) TestParseCloudTag(c *gc.C) {
	for i, t := range []struct {
		tag      string
		expected names.Tag
		err      error
	}{{
		tag: "",
		err: names.InvalidTagError("", ""),
	}, {
		tag:      "cloud-aws",
		expected: names.NewCloudTag("aws"),
	}, {
		tag: "aws",
		err: names.InvalidTagError("aws", ""),
	}, {
		tag: "unit-aws",
		err: names.InvalidTagError("unit-aws", names.UnitTagKind), // not a valid unit name either
	}, {
		tag: "application-aws",
		err: names.InvalidTagError("application-aws", names.CloudTagKind),
	}} {
		c.Logf("test %d: %s", i, t.tag)
		got, err := names.ParseCloudTag(t.tag)
		if err != nil || t.err != nil {
			c.Check(err, gc.DeepEquals, t.err)
			continue
		}
		c.Check(got, gc.FitsTypeOf, t.expected)
		c.Check(got, gc.Equals, t.expected)
	}
}
