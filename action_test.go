// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/names/v6"
)

type actionSuite struct{}

var _ = gc.Suite(&actionSuite{})

var parseActionTagTests = []struct {
	tag      string
	expected names.Tag
	err      error
}{
	{tag: "", err: names.InvalidTagError("", "")},
	{tag: "action-f47ac10b-58cc-4372-a567-0e02b2c3d479", expected: names.NewActionTag("f47ac10b-58cc-4372-a567-0e02b2c3d479")},
	{tag: "action-1", expected: names.NewActionTag("1")},
	{tag: "action-foo", err: names.InvalidTagError("action-foo", "action")},
	{tag: "bob", err: names.InvalidTagError("bob", "")},
	{tag: "application-ned", err: names.InvalidTagError("application-ned", names.ActionTagKind)}}

func (s *actionSuite) TestParseActionTag(c *gc.C) {
	for i, t := range parseActionTagTests {
		c.Logf("test %d: %s", i, t.tag)
		got, err := names.ParseActionTag(t.tag)
		if t.err != nil {
			c.Check(err, gc.DeepEquals, t.err)
			continue
		}
		c.Check(err, jc.ErrorIsNil)
		c.Check(got, gc.FitsTypeOf, t.expected)
		c.Check(got, gc.Equals, t.expected)
	}
}

func (s *actionSuite) TestActionReceiverTag(c *gc.C) {
	testCases := []struct {
		name     string
		expected names.Tag
		err      string
	}{
		{name: "mysql", err: `invalid actionreceiver name "mysql"`},
		{name: "mysql/3", expected: names.NewUnitTag("mysql/3")},
		{name: "3", expected: names.NewMachineTag("3")},
	}

	for _, tcase := range testCases {
		tag, err := names.ActionReceiverTag(tcase.name)
		if tcase.err != "" {
			c.Check(err, gc.ErrorMatches, tcase.err)
			c.Check(tag, gc.IsNil)
			continue
		}
		c.Check(err, jc.ErrorIsNil)
		c.Check(tag, gc.FitsTypeOf, tcase.expected)
		c.Check(tag, gc.Equals, tcase.expected)
	}

}

func (s *actionSuite) TestActionReceiverFromTag(c *gc.C) {
	for i, test := range []struct {
		name     string
		expected names.Tag
		err      string
	}{
		{name: "rambleon", err: `invalid actionreceiver tag "rambleon"`},
		{name: "unit-mysql-2", expected: names.NewUnitTag("mysql/2")},
		{name: "machine-13", expected: names.NewMachineTag("13")},
	} {
		c.Logf("test %d", i)
		tag, err := names.ActionReceiverFromTag(test.name)
		if test.err != "" {
			c.Check(err, gc.ErrorMatches, test.err)
			c.Check(tag, gc.IsNil)
			continue
		}
		c.Check(tag, gc.Equals, test.expected)
		c.Check(err, jc.ErrorIsNil)
	}
}
