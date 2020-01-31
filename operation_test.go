// Copyright 2020 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	gc "gopkg.in/check.v1"

	jc "github.com/juju/testing/checkers"
	"gopkg.in/juju/names.v3"
)

type operationSuite struct{}

var _ = gc.Suite(&operationSuite{})

var parseOperationTagTests = []struct {
	tag      string
	expected names.Tag
	err      error
}{
	{tag: "", err: names.InvalidTagError("", "")},
	{tag: "operation-1", expected: names.NewOperationTag("1")},
	{tag: "operation-foo", err: names.InvalidTagError("operation-foo", "operation")},
	{tag: "bob", err: names.InvalidTagError("bob", "")},
	{tag: "application-ned", err: names.InvalidTagError("application-ned", names.OperationTagKind)}}

func (s *operationSuite) TestParseOperationTag(c *gc.C) {
	for i, t := range parseOperationTagTests {
		c.Logf("test %d: %s", i, t.tag)
		got, err := names.ParseOperationTag(t.tag)
		if t.err != nil {
			c.Check(err, gc.DeepEquals, t.err)
			continue
		}
		c.Check(err, jc.ErrorIsNil)
		c.Check(got, gc.FitsTypeOf, t.expected)
		c.Check(got, gc.Equals, t.expected)
	}
}

func (s *operationSuite) TestString(c *gc.C) {
	tag := names.NewOperationTag("666")
	c.Assert(tag.String(), gc.Equals, "operation-666")
}
