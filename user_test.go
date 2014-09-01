// Copyright 2013 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	gc "launchpad.net/gocheck"

	"github.com/juju/names"
)

type userSuite struct{}

var _ = gc.Suite(&userSuite{})

func (s *userSuite) TestUserTag(c *gc.C) {
	for i, t := range []struct {
		input    string
		string   string
		name     string
		provider string
		username string
	}{
		{
			input:    "bob",
			string:   "user-bob",
			name:     "bob",
			provider: "local",
			username: "bob@local",
		}, {
			input:    "bob@local",
			string:   "user-bob@local",
			name:     "bob",
			provider: "local",
			username: "bob@local",
		}, {
			input:    "bob@foo",
			string:   "user-bob@foo",
			name:     "bob",
			provider: "foo",
			username: "bob@foo",
		},
	} {
		c.Logf("test %d: %s", i, t.input)
		userTag := names.NewUserTag(t.input)
		c.Check(userTag.String(), gc.Equals, t.string)
		c.Check(userTag.Id(), gc.Equals, t.input)
		c.Check(userTag.Name(), gc.Equals, t.name)
		c.Check(userTag.Provider(), gc.Equals, t.provider)
		c.Check(userTag.Username(), gc.Equals, t.username)
	}
}

func (s *userSuite) TestIsValidUser(c *gc.C) {
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
		{"bob+1", false},
		{"bob.1", true},
		{"1bob", false},
		{"1-bob", false},
		{"1+bob", false},
		{"1.bob", false},
		{"jim.bob+99-1.", false},
		{"a", false},
		{"0foo", false},
		{"foo bar", false},
		{"bar{}", false},
		{"bar+foo", false},
		{"bar_foo", false},
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
		{"bar@ram.u", true},
		{"bar@", false},
		{"@local", false},
		{"not/valid", false},
	} {
		c.Logf("test %d: %s", i, t.string)
		c.Assert(names.IsValidUser(t.string), gc.Equals, t.expect, gc.Commentf("%s", t.string))
	}
}

func (s *userSuite) TestIsValidUserName(c *gc.C) {
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
		{"bob+1", false},
		{"bob.1", true},
		{"1bob", false},
		{"1-bob", false},
		{"1+bob", false},
		{"1.bob", false},
		{"jim.bob+99-1.", false},
		{"a", false},
		{"0foo", false},
		{"foo bar", false},
		{"bar{}", false},
		{"bar+foo", false},
		{"bar_foo", false},
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
		{"bar@ram.u", false},
		{"bar@local", false},
		{"bar@ubuntuone", false},
		{"bar@", false},
		{"@local", false},
		{"not/valid", false},
	} {
		c.Logf("test %d: %s", i, t.string)
		c.Assert(names.IsValidUserName(t.string), gc.Equals, t.expect, gc.Commentf("%s", t.string))
	}
}

var parseUserTagTests = []struct {
	tag      string
	expected names.Tag
	err      error
}{{
	tag: "",
	err: names.InvalidTagError("", ""),
}, {
	tag:      "user-dave",
	expected: names.NewUserTag("dave"),
}, {
	tag:      "user-dave@local",
	expected: names.NewUserTag("dave@local"),
}, {
	tag:      "user-dave@foobar",
	expected: names.NewUserTag("dave@foobar"),
}, {
	tag: "dave",
	err: names.InvalidTagError("dave", ""),
}, {
	tag: "unit-dave",
	err: names.InvalidTagError("unit-dave", names.UnitTagKind), // not a valid unit name either
}, {
	tag: "service-dave",
	err: names.InvalidTagError("service-dave", names.UserTagKind),
}}

func (s *userSuite) TestParseUserTag(c *gc.C) {
	for i, t := range parseUserTagTests {
		c.Logf("test %d: %s", i, t.tag)
		got, err := names.ParseUserTag(t.tag)
		if err != nil || t.err != nil {
			c.Check(err, gc.DeepEquals, t.err)
			continue
		}
		c.Check(got, gc.FitsTypeOf, t.expected)
		c.Check(got, gc.Equals, t.expected)
	}
}
