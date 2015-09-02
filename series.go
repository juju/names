package names

import (
	"fmt"
	"regexp"
)

// ValidSeries matches any valid series string.
var ValidSeries = regexp.MustCompile("^[a-z]+([a-z0-9]+)?$")

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
	"win2012hvr2": true,
	"win2012hv":   true,
	"win2012r2":   true,
	"win2012":     true,
	"win7":        true,
	"win8":        true,
	"win81":       true,
	"centos7":     true,
}

// IsValidSeries checks to see if a given string is a valid series string and if
// it is a known series.
func IsValidSeries(series string) (bool, error) {
	if !ValidSeries.MatchString(series) {
		return false, fmt.Errorf("invalid series format: %q", series)
	}
	if !KnownSeries[series] {
		return false, fmt.Errorf("unknown series: %q", series)
	}
	return true, nil
}
