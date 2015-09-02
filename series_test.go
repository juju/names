package names_test

import (
	gc "gopkg.in/check.v1"

	"github.com/juju/names"
)

type seriesSuite struct{}

var _ = gc.Suite(&seriesSuite{})

var seriesNameTests = []struct {
	pattern, err string
	valid        bool
}{
	{
		pattern: "bundle",
		valid:   true,
	},
	{
		pattern: "oneiric",
		valid:   true,
	},
	{
		pattern: "precise",
		valid:   true,
	},
	{
		pattern: "quantal",
		valid:   true,
	},
	{
		pattern: "raring",
		valid:   true,
	},
	{
		pattern: "saucy",
		valid:   true,
	},
	{
		pattern: "trusty",
		valid:   true,
	},
	{
		pattern: "utopic",
		valid:   true,
	},
	{
		pattern: "vivid",
		valid:   true,
	},
	{
		pattern: "win2012hvr2",
		valid:   true,
	},
	{
		pattern: "win2012hv",
		valid:   true,
	},
	{
		pattern: "win2012r2",
		valid:   true,
	},
	{
		pattern: "win2012",
		valid:   true,
	},
	{
		pattern: "win7",
		valid:   true,
	},
	{
		pattern: "win8",
		valid:   true,
	},
	{
		pattern: "win81",
		valid:   true,
	},
	{
		pattern: "centos7",
		valid:   true,
	},
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

func (s *seriesSuite) TestIsValidSeries(c *gc.C) {
	for i, test := range seriesNameTests {
		c.Logf("test %d: %q", i, test.pattern)
		result, err := names.IsValidSeries(test.pattern)
		if test.valid {
			c.Assert(err, gc.IsNil)
			c.Assert(result, gc.Equals, true)
		} else {
			c.Assert(err, gc.ErrorMatches, test.err)
			c.Assert(result, gc.Equals, false)
		}
	}
}
