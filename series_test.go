// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	gc "gopkg.in/check.v1"

	"github.com/juju/names"
)

type seriesSuite struct{}

var _ = gc.Suite(&seriesSuite{})

var invalidSeriesNameTests = []struct {
	pattern, err string
	valid        bool
}{
	{
		pattern: ".bad-wolf",
		valid:   false,
		err:     "invalid series format: .*",
	},
	{
		pattern: "badseries",
		valid:   false,
		err:     "unknown series: .*",
	},
}

func (s *seriesSuite) TestValidVerifySeries(c *gc.C) {
	for series, _ := range names.KnownSeries {
		c.Logf("test %q", series)
		result, err := names.VerifySeries(series)
		c.Assert(result, gc.Equals, true)
		c.Assert(err, gc.IsNil)
	}
}

func (s *seriesSuite) TestInvalidVerifySeries(c *gc.C) {
	for i, test := range invalidSeriesNameTests {
		c.Logf("test %d: %q", i, test.pattern)
		result, err := names.VerifySeries(test.pattern)
		if test.valid {
			c.Assert(err, gc.IsNil)
			c.Assert(result, gc.Equals, true)
		} else {
			c.Assert(err, gc.ErrorMatches, test.err)
			c.Assert(result, gc.Equals, false)
		}
	}
}
