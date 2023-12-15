// Copyright 2019 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/names/v5"
)

type ControllerAgentSuite struct{}

var _ = gc.Suite(&ControllerAgentSuite{})

func (s *ControllerAgentSuite) TestControllerAgentTag(c *gc.C) {
	c.Assert(names.NewControllerAgentTag("123").String(), gc.Equals, "controller-123")
}

func (s *ControllerAgentSuite) TestIdFormats(c *gc.C) {
	c.Assert(names.IsValidControllerAgent("123"), jc.IsTrue)
	c.Assert(names.IsValidControllerAgent("-123"), jc.IsFalse)
	c.Assert(names.IsValidControllerAgent("invalid"), jc.IsFalse)
}

func (s *ControllerAgentSuite) TestNumber(c *gc.C) {
	ca := names.ControllerAgentTag{}
	c.Assert(ca.Number(), gc.Equals, 0)
	ca = names.NewControllerAgentTag("5")
	c.Assert(ca.Number(), gc.Equals, 5)
}

var parseControllerAgentTagTests = []struct {
	tag      string
	expected names.Tag
	err      error
}{{
	tag: "",
	err: names.InvalidTagError("", ""),
}, {
	tag:      "controller-0",
	expected: names.NewControllerAgentTag("0"),
}, {
	tag: "dave",
	err: names.InvalidTagError("dave", ""),
}, {
	tag: "controller-dave",
	err: names.InvalidTagError("controller-dave", names.ControllerAgentTagKind),
}}

func (s *ControllerAgentSuite) TestParseControllerAgentTag(c *gc.C) {
	for i, t := range parseControllerAgentTagTests {
		c.Logf("test %d: %s", i, t.tag)
		got, err := names.ParseControllerAgentTag(t.tag)
		if err != nil || t.err != nil {
			c.Check(err, gc.DeepEquals, t.err)
			continue
		}
		c.Check(got, gc.FitsTypeOf, t.expected)
		c.Check(got, gc.Equals, t.expected)
	}
}
