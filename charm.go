// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// CharmTagKind specifies charm tag kind
const CharmTagKind = "charm"

// CharmTag represents tag for charm
// using charm's URL
type CharmTag struct {
	schema       string
	architecture string
	series       string
	name         string
	revision     int
}

var (
	// schemaSnippet is a regex snippet representing a valid charm URL
	// schemas
	schemaSnippet = "(local:|ch:)?"
	// archSnippet is a regex snippet representing valid architectures
	archSnippet = "([a-z]+([a-z0-9]+)?/)?"
	// supportedArches is a regular expression representing supported
	// architectures
	supportedArches = regexp.MustCompile("^(amd64|arm64|ppc64el|s390x|riscv64)$")
	// seriesSnippet is a regex snippet representing valid series
	seriesSnippet = "([a-z]+([a-z0-9]+)?/)?"
	// charmNameSnippet is a regex snippet representing valid charm
	// names
	charmNameSnippet = "[a-z][a-z0-9]*(-[a-z0-9]*[a-z][a-z0-9]*)*"
	// revisionSnippet is a regex snippet representing valid revisions
	revisionSnippet = "(-(-1|0|[1-9][0-9]*))?"
	// charmURLRegex is a regular expression for a valid charm url
	// (schema:(arch/)?(series/)?name(-revision)?)
	charmURLRegex = regexp.MustCompile("^" +
		schemaSnippet +
		archSnippet +
		seriesSnippet +
		charmNameSnippet +
		revisionSnippet + "$")
)

// String satisfies Tag interface.
// Produces string representation of charm tag.
func (t CharmTag) String() string { return t.Kind() + "-" + t.Id() }

// Kind satisfies Tag interface.
// Returns Charm tag kind.
func (t CharmTag) Kind() string { return CharmTagKind }

// Id satisfies Tag interface.
// NOTE(nvinuesa): Currently, this method returns the URL of the charm as
// it's accepted today. This might change in the future.
func (t CharmTag) Id() string {
	var parts []string
	if t.architecture != "" {
		parts = append(parts, t.architecture)
	}
	if t.series != "" {
		parts = append(parts, t.series)
	}
	if t.revision >= 0 {
		parts = append(parts, fmt.Sprintf("%s-%d", t.name, t.revision))
	} else {
		parts = append(parts, t.name)
	}
	return fmt.Sprintf("%s:%s", t.schema, strings.Join(parts, "/"))
}

// Source returns the source of the charm.
func (t CharmTag) Source() string {
	if t.schema == "local" {
		return "local"
	}
	// By default, assume charmhub.
	return "charmhub"
}

// Architecture returns the architecture of the charm.
func (t CharmTag) Architecture() string {
	return t.architecture
}

// Series returns the series of the charm.
func (t CharmTag) Series() string {
	return t.series
}

// Name returns the name of the charm.
func (t CharmTag) Name() string {
	return t.name
}

// Revision returns the revision of the charm.
func (t CharmTag) Revision() int {
	return t.revision
}

// NewCharmTag returns the tag for the charm with the given name, source,
// revision and architecture.
func NewCharmTag(id string) CharmTag {
	return parseCharmURL(id)
}

var emptyTag = CharmTag{}

// ParseCharmTag parses a charm tag string.
func ParseCharmTag(charmTag string) (CharmTag, error) {
	tag, err := ParseTag(charmTag)
	if err != nil {
		return emptyTag, err
	}
	ct, ok := tag.(CharmTag)
	if !ok {
		return emptyTag, invalidTagError(charmTag, CharmTagKind)
	}
	return ct, nil
}

// IsValidCharm returns whether name is a valid charm url.
func IsValidCharm(url string) bool {
	return charmURLRegex.MatchString(url)
}

func parseCharmURL(url string) CharmTag {
	// URLs without schema are assumed to be charmhub charms.
	schema := "ch:"
	nameRev := ""
	arch := ""
	series := ""
	urlParts := strings.Split(url, ":")
	if len(urlParts) == 2 {
		schema = urlParts[0]
		url = urlParts[1]
	}
	pathParts := strings.Split(url, "/")
	switch len(pathParts) {
	case 3:
		arch, series, nameRev = pathParts[0], pathParts[1], pathParts[2]
	case 2:
		// Since both the architecture and series are optional,
		// the first part can be either architecture or series.
		// To differentiate between them, we go ahead and try to
		// validate the first part as an architecture to decide.

		if isValidArchitecture(pathParts[0]) {
			arch, nameRev = pathParts[0], pathParts[1]
		} else {
			series, nameRev = pathParts[0], pathParts[1]
		}

	default:
		nameRev = pathParts[0]
	}
	name, revision := extractRevision(nameRev)
	return CharmTag{
		schema:       schema,
		architecture: arch,
		series:       series,
		name:         name,
		revision:     revision,
	}
}

func extractRevision(name string) (string, int) {
	revision := -1
	for i := len(name) - 1; i > 0; i-- {
		c := name[i]
		if c >= '0' && c <= '9' {
			continue
		}
		if c == '-' && i != len(name)-1 {
			var err error
			revision, err = strconv.Atoi(name[i+1:])
			if err != nil {
				panic(err) // We just checked it was right.
			}
			name = name[:i]
		}
		break
	}
	return name, revision
}

func isValidArchitecture(architecture string) bool {
	return supportedArches.MatchString(architecture)
}
