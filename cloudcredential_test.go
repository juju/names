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
		input  string
		string string
		cloud  names.CloudTag
		owner  names.UserTag
		name   string
	}{
		{
			input:  "aws/bob/foo",
			string: "cloudcred-aws_bob_foo",
			cloud:  names.NewCloudTag("aws"),
			owner:  names.NewUserTag("bob"),
			name:   "foo",
		}, {
			input:  "manual_cloud/bob/foo",
			string: "cloudcred-manual%5fcloud_bob_foo",
			cloud:  names.NewCloudTag("manual_cloud"),
			owner:  names.NewUserTag("bob"),
			name:   "foo",
		}, {
			input:  "aws/bob@remote/foo",
			string: "cloudcred-aws_bob@remote_foo",
			cloud:  names.NewCloudTag("aws"),
			owner:  names.NewUserTag("bob@remote"),
			name:   "foo",
		}, {
			input:  "aws/bob@remote/foo@somewhere.com",
			string: "cloudcred-aws_bob@remote_foo@somewhere.com",
			cloud:  names.NewCloudTag("aws"),
			owner:  names.NewUserTag("bob@remote"),
			name:   "foo@somewhere.com",
		}, {
			input:  "aws/bob@remote/foo_bar",
			string: `cloudcred-aws_bob@remote_foo%5fbar`,
			cloud:  names.NewCloudTag("aws"),
			owner:  names.NewUserTag("bob@remote"),
			name:   "foo_bar",
		}, {
			input:  "google/bob+bob@remote/foo_bar",
			string: `cloudcred-google_bob+bob@remote_foo%5fbar`,
			cloud:  names.NewCloudTag("google"),
			owner:  names.NewUserTag("bob+bob@remote"),
			name:   "foo_bar",
		},
	} {
		c.Logf("test %d: %s", i, t.input)
		cloudTag := names.NewCloudCredentialTag(t.input)
		c.Check(cloudTag.String(), gc.Equals, t.string)
		c.Check(cloudTag.Id(), gc.Equals, t.input)
		c.Check(cloudTag.Cloud(), gc.Equals, t.cloud)
		c.Check(cloudTag.Owner(), gc.Equals, t.owner)
		c.Check(cloudTag.Name(), gc.Equals, t.name)
	}
}

func (s *cloudCredentialSuite) TestIsValidCloudCredential(c *gc.C) {
	for i, t := range []struct {
		string string
		expect bool
	}{
		{"", false},
		{"aws/bob/foo", true},
		{"manual_cloud/bob/foo", true},
		{"aws/bob@local/foo", true},
		{"google/bob+bob@local/foo", true},
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
		{"foo@bar", true},
		{"foo+foo@bar", true},
		{"foo_bar", true},
		{"123", false},
		{"0foo", false},
	} {
		c.Logf("test %d: %s", i, t.string)
		c.Check(names.IsValidCloudCredentialName(t.string), gc.Equals, t.expect, gc.Commentf("%s", t.string))
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
		tag:      "cloudcred-manual%5fcloud_bob_foo",
		expected: names.NewCloudCredentialTag("manual_cloud/bob/foo"),
	}, {
		tag:      "cloudcred-aws-china_bob_foo-manchu",
		expected: names.NewCloudCredentialTag("aws-china/bob/foo-manchu"),
	}, {
		tag:      "cloudcred-aws-china_bob_foo@somewhere.com",
		expected: names.NewCloudCredentialTag("aws-china/bob/foo@somewhere.com"),
	}, {
		tag:      `cloudcred-aws-china_bob_foo%5fbar`,
		expected: names.NewCloudCredentialTag("aws-china/bob/foo_bar"),
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

func (s *cloudCredentialSuite) TestIsZero(c *gc.C) {
	c.Assert(names.CloudCredentialTag{}.IsZero(), gc.Equals, true)
	c.Assert(names.NewCloudCredentialTag("aws/bob/foo").IsZero(), gc.Equals, false)
}

func (s *cloudCredentialSuite) TestZeroString(c *gc.C) {
	c.Assert(names.CloudCredentialTag{}.String(), gc.Equals, "")
}

func (s *cloudCredentialSuite) TestZeroId(c *gc.C) {
	c.Assert(names.CloudCredentialTag{}.Id(), gc.Equals, "")
}
