// Copyright 2015 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	gc "gopkg.in/check.v1"

	"github.com/juju/names/v6"
)

type subnetSuite struct{}

var _ = gc.Suite(&subnetSuite{})

func (s *subnetSuite) TestNewSubnetTag(c *gc.C) {
	id := "0195847b-95bb-7ca1-a7ee-2211d802d5b3"
	tag := names.NewSubnetTag(id)
	parsed, err := names.ParseSubnetTag(tag.String())
	c.Assert(err, gc.IsNil)
	c.Assert(parsed.Kind(), gc.Equals, names.SubnetTagKind)
	c.Assert(parsed.Id(), gc.Equals, id)
	c.Assert(parsed.String(), gc.Equals, names.SubnetTagKind+"-"+id)

	f := func() {
		tag = names.NewSubnetTag("foo")
	}
	c.Assert(f, gc.PanicMatches, "foo is not a valid subnet ID")
}

var parseSubnetTagTests = []struct {
	tag      string
	expected names.Tag
	err      error
}{{
	tag: "",
	err: names.InvalidTagError("", ""),
}, {
	tag:      "subnet-16",
	expected: names.NewSubnetTag("16"),
}, {
	tag:      "subnet-0195847b-95bb-7ca1-a7ee-2211d802d5b3",
	expected: names.NewSubnetTag("0195847b-95bb-7ca1-a7ee-2211d802d5b3"),
}, {
	tag: "subnet-foo",
	err: names.InvalidTagError("subnet-foo", names.SubnetTagKind),
}, {
	tag: "subnet-",
	err: names.InvalidTagError("subnet-", names.SubnetTagKind),
}, {
	tag: "foobar",
	err: names.InvalidTagError("foobar", ""),
}, {
	tag: "unit-foo-0",
	err: names.InvalidTagError("unit-foo-0", names.SubnetTagKind),
}}

func (s *subnetSuite) TestParseSubnetTag(c *gc.C) {
	for i, t := range parseSubnetTagTests {
		c.Logf("test %d: %s", i, t.tag)
		got, err := names.ParseSubnetTag(t.tag)
		if err != nil || t.err != nil {
			c.Check(err, gc.DeepEquals, t.err)
			continue
		}
		c.Check(got, gc.FitsTypeOf, t.expected)
		c.Check(got, gc.Equals, t.expected)
	}
}
