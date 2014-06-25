// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package names_test

import (
	gc "launchpad.net/gocheck"

	"github.com/juju/names"
)

type actionSuite struct{}

var _ = gc.Suite(&actionSuite{})

var marker = names.ActionMarker
var actionNameTests = []struct {
	pattern string
	valid   bool
}{
	{pattern: "", valid: false},
	{pattern: "service", valid: false},
	{pattern: "service" + marker, valid: false},
	{pattern: "service" + marker + "0", valid: false},
	{pattern: "service" + marker + "00", valid: false},
	{pattern: "service" + marker + "0" + marker + "0", valid: false},

	{pattern: "service-name/0" + marker, valid: false},
	{pattern: "service-name-0" + marker, valid: false},
	{pattern: "service-name/0" + marker + "0", valid: true},
	{pattern: "service-name-0" + marker + "0", valid: false},

	{pattern: "service-name/0" + marker + "00", valid: false},
	{pattern: "service-name-0" + marker + "00", valid: false},
	{pattern: "service-name/0" + marker + "01", valid: false},
	{pattern: "service-name-0" + marker + "01", valid: false},
	{pattern: "service-name/0" + marker + "11", valid: true},
	{pattern: "service-name-0" + marker + "11", valid: false},
}

func (s *actionSuite) TestActionNameFormats(c *gc.C) {
	assertAction := func(s string, expect bool) {
		c.Assert(names.IsAction(s), gc.Equals, expect)
	}

	for i, test := range actionNameTests {
		c.Logf("test %d: %q", i, test.pattern)
		assertAction(test.pattern, test.valid)
	}
}

var parseActionTagTests = []struct {
	tag      string
	expected names.Tag
	err      error
}{{
	tag:      "",
	expected: nil,
	err:      names.InvalidTagError("", ""),
}, {
	tag:      "action-dave" + names.ActionMarker + "123",
	expected: makeActionTag("action-dave" + names.ActionMarker + "123"),
	err:      names.InvalidTagError("action-dave"+names.ActionMarker+"123", names.ActionTagKind),
}, {
	tag:      "dave",
	expected: nil,
	err:      names.InvalidTagError("dave", ""),
}, {
	tag:      "action-dave/0",
	expected: nil,
	err:      names.InvalidTagError("action-dave/0", names.ActionTagKind),
}, {
	tag:      "action",
	expected: nil,
	err:      names.InvalidTagError("action", ""),
}, {
	tag:      "user-dave",
	expected: nil,
	err:      names.InvalidTagError("user-dave", names.ActionTagKind),
}}

func (s *actionSuite) TestParseActionTag(c *gc.C) {
	for i, t := range parseActionTagTests {
		c.Logf("test %d: %s", i, t.tag)
		got, err := names.ParseActionTag(t.tag)
		if err != nil || t.err != nil {
			c.Check(err, gc.DeepEquals, t.err)
			continue
		}
		c.Check(got, gc.FitsTypeOf, t.expected)
		c.Check(got, gc.Equals, t.expected)
	}
}

func makeActionTag(actionId string) names.ActionTag {
	if tag, ok := names.ParseActionId(actionId); ok {
		return tag
	}
	return names.ActionTag{}
}
