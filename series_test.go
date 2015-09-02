// Copyright 2015 Canonical Ltd.
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
}{
	{
		pattern: ".bad-wolf",
		err:     "invalid series format: .*",
	},
	{
		pattern: "badseries",
		err:     "unknown series: .*",
	},
}

func (s *seriesSuite) TestInvalidVerifySeries(c *gc.C) {
	for i, test := range invalidSeriesNameTests {
		c.Logf("test %d: %q", i, test.pattern)
		err := names.VerifySeries(test.pattern)
		c.Assert(err, gc.ErrorMatches, test.err)
	}
}

func (s *seriesSuite) TestValidVerifySeries(c *gc.C) {
	for series := range names.KnownSeries {
		c.Logf("test %q", series)
		err := names.VerifySeries(series)
		c.Assert(err, gc.IsNil)
	}
}
