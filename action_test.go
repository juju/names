// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	"fmt"

	gc "launchpad.net/gocheck"

	"github.com/juju/names"
)

type actionSuite struct{}

var _ = gc.Suite(&actionSuite{})

func (s *actionSuite) TestActionNameFormats(c *gc.C) {
	marker := names.ActionMarker
	actionNameTests := []struct {
		pattern string
		valid   bool
	}{
		{pattern: "", valid: false},
		{pattern: "service", valid: false},
		{pattern: "service" + marker, valid: false},
		{pattern: "service" + marker + "0", valid: true},
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

	assertAction := func(s string, expect bool) {
		c.Assert(names.IsValidAction(s), gc.Equals, expect)
	}

	for i, test := range actionNameTests {
		c.Logf("test %d: %q", i, test.pattern)
		assertAction(test.pattern, test.valid)
	}
}

func (s *actionSuite) TestInvalidActionNamesPanic(c *gc.C) {
	invalidActionNameTests := []string{
		"",      // blank is not a valid action id
		"admin", // probably a user name, which isn't a valid action id
	}

	for _, name := range invalidActionNameTests {
		expect := fmt.Sprintf("%q is not a valid action id", name)
		testFunc := func() { names.NewActionTag(name) }
		c.Assert(testFunc, gc.PanicMatches, expect)
	}
}

func (s *actionSuite) TestParseActionTag(c *gc.C) {
	parseActionTagTests := []struct {
		tag      string
		expected names.Tag
		err      error
	}{
		{
			tag:      "",
			expected: nil,
			err:      names.InvalidTagError("", ""),
		}, {
			tag:      "action-good" + names.ActionMarker + "123",
			expected: names.NewActionTag("good" + names.ActionMarker + "123"),
			err:      nil,
		}, {
			tag:      "action-good/0" + names.ActionMarker + "123",
			expected: names.NewActionTag("good/0" + names.ActionMarker + "123"),
			err:      nil,
		}, {
			tag:      "action-bad/00" + names.ActionMarker + "123",
			expected: nil,
			err:      names.InvalidTagError("action-bad/00"+names.ActionMarker+"123", names.ActionTagKind),
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

func (s *actionSuite) TestActionResultNameFormats(c *gc.C) {
	marker := names.ActionResultMarker
	actionResultNameTests := []struct {
		pattern string
		valid   bool
	}{
		{pattern: "", valid: false},
		{pattern: "service", valid: false},
		{pattern: "service" + marker, valid: false},
		{pattern: "service" + marker + "0", valid: true},
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

	assertActionResult := func(s string, expect bool) {
		c.Assert(names.IsValidActionResult(s), gc.Equals, expect)
	}

	for i, test := range actionResultNameTests {
		c.Logf("test %d: %q", i, test.pattern)
		assertActionResult(test.pattern, test.valid)
	}
}

func (s *actionSuite) TestInvalidActionResultNamesPanic(c *gc.C) {
	invalidActionResultNameTests := []string{
		"",         // blank is not a valid action id
		"admin",    // probably a user name, which isn't a valid action id
		"blah_a_0", // accidentally use an actionId instead of actionResultId
	}

	for _, name := range invalidActionResultNameTests {
		expect := fmt.Sprintf("%q is not a valid action result id", name)
		testFunc := func() { names.NewActionResultTag(name) }
		c.Assert(testFunc, gc.PanicMatches, expect)
	}
}

func (s *actionSuite) TestParseActionResultTag(c *gc.C) {
	parseActionResultTagTests := []struct {
		tag      string
		expected names.Tag
		err      error
	}{
		{
			tag:      "",
			expected: nil,
			err:      names.InvalidTagError("", ""),
		}, {
			tag:      "actionresult-good" + names.ActionResultMarker + "123",
			expected: names.NewActionResultTag("good" + names.ActionResultMarker + "123"),
			err:      nil,
		}, {
			tag:      "actionresult-good/0" + names.ActionResultMarker + "123",
			expected: names.NewActionResultTag("good/0" + names.ActionResultMarker + "123"),
			err:      nil,
		}, {
			tag:      "actionresult-bad/00" + names.ActionResultMarker + "123",
			expected: nil,
			err:      names.InvalidTagError("actionresult-bad/00"+names.ActionResultMarker+"123", names.ActionResultTagKind),
		}, {
			tag:      "dave",
			expected: nil,
			err:      names.InvalidTagError("dave", ""),
		}, {
			tag:      "actionresult-dave/0",
			expected: nil,
			err:      names.InvalidTagError("actionresult-dave/0", names.ActionResultTagKind),
		}, {
			tag:      "action",
			expected: nil,
			err:      names.InvalidTagError("action", ""),
		}, {
			tag:      "user-dave",
			expected: nil,
			err:      names.InvalidTagError("user-dave", names.ActionResultTagKind),
		}}

	for i, t := range parseActionResultTagTests {
		c.Logf("test %d: %s", i, t.tag)
		got, err := names.ParseActionResultTag(t.tag)
		if err != nil || t.err != nil {
			c.Check(err, gc.DeepEquals, t.err)
			continue
		}
		c.Check(got, gc.FitsTypeOf, t.expected)
		c.Check(got, gc.Equals, t.expected)
	}
}

func (s *actionSuite) TestPrefixSuffix(c *gc.C) {
	var tests = []struct {
		prefix string
		suffix int
	}{
		{prefix: "asdf", suffix: 0},
		{prefix: "qwer/0", suffix: 10},
		{prefix: "zxcv/3", suffix: 11},
	}

	for _, test := range tests {
		suf := fmt.Sprintf("%d", test.suffix)

		action := names.NewActionTag(test.prefix + names.ActionMarker + suf)
		c.Assert(action.Prefix(), gc.Equals, test.prefix)
		c.Assert(action.Sequence(), gc.Equals, test.suffix)

		result := names.NewActionResultTag(test.prefix + names.ActionResultMarker + suf)
		c.Assert(result.Prefix(), gc.Equals, test.prefix)
		c.Assert(result.Sequence(), gc.Equals, test.suffix)

		c.Assert(action.PrefixTag(), gc.Not(gc.IsNil))
		c.Assert(action.PrefixTag(), gc.DeepEquals, result.PrefixTag())
	}
}
