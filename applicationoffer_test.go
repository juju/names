// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	gc "gopkg.in/check.v1"

	"gopkg.in/juju/names.v2"
)

type applicationOfferSuite struct{}

var _ = gc.Suite(&applicationOfferSuite{})

var parseApplicationOfferTagTests = []struct {
	tag      string
	expected names.Tag
	err      error
}{{
	tag: "",
	err: names.InvalidTagError("", ""),
}, {
	tag:      "applicationoffer-dave",
	expected: names.NewApplicationOfferTag("dave"),
}, {
	tag: "dave",
	err: names.InvalidTagError("dave", ""),
}, {
	tag: "applicationoffer-dave/0",
	err: names.InvalidTagError("applicationoffer-dave/0", names.ApplicationOfferTagKind),
}, {
	tag: "applicationoffer",
	err: names.InvalidTagError("applicationoffer", ""),
}, {
	tag: "user-dave",
	err: names.InvalidTagError("user-dave", names.ApplicationOfferTagKind),
}}

func (s *applicationOfferSuite) TestParseApplicationOfferTag(c *gc.C) {
	for i, t := range parseApplicationOfferTagTests {
		c.Logf("test %d: %s", i, t.tag)
		got, err := names.ParseApplicationOfferTag(t.tag)
		if err != nil || t.err != nil {
			c.Check(err, gc.DeepEquals, t.err)
			continue
		}
		c.Check(got, gc.FitsTypeOf, t.expected)
		c.Check(got, gc.Equals, t.expected)
	}
}
