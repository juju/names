// Copyright 2015 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"fmt"
	"regexp"
)

// ValidSeries matches any valid series string.
var validSeries = regexp.MustCompile("^" + SeriesSnippet + "$")

// KnownSeries is a map of all the different series' that juju knows about.
var KnownSeries = map[string]bool{
	"bundle":      true,
	"oneiric":     true,
	"precise":     true,
	"quantal":     true,
	"raring":      true,
	"saucy":       true,
	"trusty":      true,
	"utopic":      true,
	"vivid":       true,
	"wily":        true,
	"win2012hvr2": true,
	"win2012hv":   true,
	"win2012r2":   true,
	"win2012":     true,
	"win7":        true,
	"win8":        true,
	"win81":       true,
	"centos7":     true,
}

// VerifySeries checks to see if a given string is a valid series string and if
// it is a known series.
func VerifySeries(series string) error {
	if !validSeries.MatchString(series) {
		return fmt.Errorf("invalid series format: %q", series)
	}
	if !KnownSeries[series] {
		return fmt.Errorf("unknown series: %q", series)
	}
	return nil
}
