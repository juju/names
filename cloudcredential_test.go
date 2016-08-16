// Copyright 2016 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	gc "gopkg.in/check.v1"

	"gopkg.in/juju/names.v2"
)

type cloudCredentialSuite struct{}

var _ = gc.Suite(&cloudCredentialSuite{})

func (s *cloudCredentialSuite) TestCloudCredentialTag(c *gc.C) {
	for i, t := range []struct {
		input     string
		string    string
		cloud     names.CloudTag
		owner     names.UserTag
		name      string
		canonical string
	}{
		{
			input:     "aws/bob/foo",
			canonical: "aws/bob@local/foo",
			string:    "cloudcred-aws_bob_foo",
			cloud:     names.NewCloudTag("aws"),
			owner:     names.NewUserTag("bob"),
			name:      "foo",
		}, {
			input:     "aws/bob@remote/foo",
			canonical: "aws/bob@remote/foo",
			string:    "cloudcred-aws_bob@remote_foo",
			cloud:     names.NewCloudTag("aws"),
			owner:     names.NewUserTag("bob@remote"),
			name:      "foo",
		},
	} {
		c.Logf("test %d: %s", i, t.input)
		cloudTag := names.NewCloudCredentialTag(t.input)
		c.Check(cloudTag.String(), gc.Equals, t.string)
		c.Check(cloudTag.Id(), gc.Equals, t.input)
		c.Check(cloudTag.Cloud(), gc.Equals, t.cloud)
		c.Check(cloudTag.Owner(), gc.Equals, t.owner)
		c.Check(cloudTag.Name(), gc.Equals, t.name)
		c.Check(cloudTag.Canonical(), gc.Equals, t.canonical)
	}
}

func (s *cloudCredentialSuite) TestIsValidCloudCredential(c *gc.C) {
	for i, t := range []struct {
		string string
		expect bool
	}{
		{"", false},
		{"aws/bob/foo", true},
		{"aws/bob@local/foo", true},
		{"/bob/foo", false},
		{"aws//foo", false},
		{"aws/bob/", false},
	} {
		c.Logf("test %d: %s", i, t.string)
		c.Assert(names.IsValidCloudCredential(t.string), gc.Equals, t.expect, gc.Commentf("%s", t.string))
	}
}

func (s *cloudCredentialSuite) TestIsValidCloudCredentialName(c *gc.C) {
	for i, t := range []struct {
		string string
		expect bool
	}{
		{"", false},
		{"foo", true},
		{"f00b4r", true},
		{"foo-bar", true},
		{"123", false},
		{"0foo", false},
	} {
		c.Logf("test %d: %s", i, t.string)
		c.Assert(names.IsValidCloudCredentialName(t.string), gc.Equals, t.expect, gc.Commentf("%s", t.string))
	}
}

func (s *cloudCredentialSuite) TestParseCloudCredentialTag(c *gc.C) {
	for i, t := range []struct {
		tag      string
		expected names.Tag
		err      error
	}{{
		tag: "",
		err: names.InvalidTagError("", ""),
	}, {
		tag:      "cloudcred-aws_bob_foo",
		expected: names.NewCloudCredentialTag("aws/bob/foo"),
	}, {
		tag:      "cloudcred-aws-china_bob_foo-manchu",
		expected: names.NewCloudCredentialTag("aws-china/bob/foo-manchu"),
	}, {
		tag: "foo",
		err: names.InvalidTagError("foo", ""),
	}, {
		tag: "unit-aws",
		err: names.InvalidTagError("unit-aws", names.UnitTagKind), // not a valid unit name either
	}} {
		c.Logf("test %d: %s", i, t.tag)
		got, err := names.ParseCloudCredentialTag(t.tag)
		if err != nil || t.err != nil {
			c.Check(err, gc.DeepEquals, t.err)
			continue
		}
		c.Check(got, gc.FitsTypeOf, t.expected)
		c.Check(got, gc.Equals, t.expected)
	}
}
