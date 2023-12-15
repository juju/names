// Copyright 2017 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

const ApplicationOfferTagKind = "applicationoffer"

// IsValidApplicationOffer returns whether name is a valid application offer name.
func IsValidApplicationOffer(uuid string) bool {
	return validUUID.MatchString(uuid)
}

type ApplicationOfferTag struct {
	Name string
}

func (t ApplicationOfferTag) String() string { return t.Kind() + "-" + t.Id() }
func (t ApplicationOfferTag) Kind() string   { return ApplicationOfferTagKind }
func (t ApplicationOfferTag) Id() string     { return t.Name }

// NewApplicationOfferTag returns the tag for the application with the given name.
func NewApplicationOfferTag(applicationOfferName string) ApplicationOfferTag {
	return ApplicationOfferTag{Name: applicationOfferName}
}

// ParseApplicationOfferTag parses a application tag string.
func ParseApplicationOfferTag(applicationOfferTag string) (ApplicationOfferTag, error) {
	tag, err := ParseTag(applicationOfferTag)
	if err != nil {
		return ApplicationOfferTag{}, err
	}
	st, ok := tag.(ApplicationOfferTag)
	if !ok {
		return ApplicationOfferTag{}, invalidTagError(applicationOfferTag, ApplicationOfferTagKind)
	}
	return st, nil
}
