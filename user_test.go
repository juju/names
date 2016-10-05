// Copyright 2013 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	"fmt"

	gc "gopkg.in/check.v1"

	"gopkg.in/juju/names.v2"
)

type userSuite struct{}

var _ = gc.Suite(&userSuite{})

func (s *userSuite) TestUserTag(c *gc.C) {
	for i, t := range []struct {
		input    string
		string   string
		name     string
		domain   string
		username string
	}{
		{
			input:    "bob",
			string:   "user-bob",
			name:     "bob",
			username: "bob",
		}, {
			input:    "bob@local",
			string:   "user-bob",
			name:     "bob",
			username: "bob",
		}, {
			input:    "bob@foo",
			string:   "user-bob@foo",
			name:     "bob",
			domain:   "foo",
			username: "bob@foo",
		},
	} {
		c.Logf("test %d: %s", i, t.input)
		userTag := names.NewUserTag(t.input)
		c.Check(userTag.String(), gc.Equals, t.string)
		c.Check(userTag.Id(), gc.Equals, t.username)
		c.Check(userTag.Name(), gc.Equals, t.name)
		c.Check(userTag.Domain(), gc.Equals, t.domain)
		c.Check(userTag.IsLocal(), gc.Equals, t.domain == "")
	}
}

var withDomainTests = []struct {
	id       string
	domain   string
	expectId string
}{{
	id:       "bob",
	domain:   names.LocalUserDomain,
	expectId: "bob",
}, {
	id:       "bob@local",
	domain:   "foo",
	expectId: "bob@foo",
}, {
	id:     "bob@local",
	domain: "",
}, {
	id:       "bob@foo",
	domain:   names.LocalUserDomain,
	expectId: "bob",
}, {
	id:     "bob",
	domain: "@foo",
}}

func (s *userSuite) TestWithDomain(c *gc.C) {
	for i, test := range withDomainTests {
		c.Logf("test %d: id %q; domain %q", i, test.id, test.domain)
		tag := names.NewUserTag(test.id)
		if test.expectId == "" {
			c.Assert(func() {
				tag.WithDomain(test.domain)
			}, gc.PanicMatches, fmt.Sprintf("invalid user domain %q", test.domain))
		} else {
			c.Assert(tag.WithDomain(test.domain).Id(), gc.Equals, test.expectId)
		}
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
		{"bob+1", true},
		{"bob+", false},
		{"+bob", false},
		{"bob.1", true},
		{"1bob", true},
		{"1-bob", true},
		{"1+bob", true},
		{"1.bob", true},
		{"jim.bob+99-1.", false},
		{"a", false},
		{"0foo", true},
		{"foo bar", false},
		{"bar{}", false},
		{"bar+foo", true},
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

func (s *userSuite) TestIsValidUserNameOrDomain(c *gc.C) {
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
		{"bob+1", true},
		{"bob+", false},
		{"+bob", false},
		{"bob.1", true},
		{"1bob", true},
		{"1-bob", true},
		{"1+bob", true},
		{"1.bob", true},
		{"jim.bob+99-1.", false},
		{"a", false},
		{"0foo", true},
		{"foo bar", false},
		{"bar{}", false},
		{"bar+foo", true},
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
		c.Assert(names.IsValidUserDomain(t.string), gc.Equals, t.expect, gc.Commentf("%s", t.string))
	}
}

func (s *userSuite) TestParseUserTag(c *gc.C) {
	for i, t := range []struct {
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
		tag: "application-dave",
		err: names.InvalidTagError("application-dave", names.UserTagKind),
	}} {
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

func (s *userSuite) TestNewLocalUserTag(c *gc.C) {
	user := names.NewLocalUserTag("bob")
	c.Assert(user.Name(), gc.Equals, "bob")
	c.Assert(user.Domain(), gc.Equals, "")
	c.Assert(user.IsLocal(), gc.Equals, true)
	c.Assert(user.String(), gc.Equals, "user-bob")

	c.Assert(func() { names.NewLocalUserTag("bob@local") }, gc.PanicMatches, `invalid user name "bob@local"`)
	c.Assert(func() { names.NewLocalUserTag("") }, gc.PanicMatches, `invalid user name ""`)
	c.Assert(func() { names.NewLocalUserTag("!@#") }, gc.PanicMatches, `invalid user name "!@#"`)
}
