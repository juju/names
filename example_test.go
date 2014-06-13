// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package names

import "fmt"

func ExampleParseTag() {
	tag, _ := ParseTag("user-100")
	switch tag := tag.(type) {
	case UserTag:
		fmt.Printf("User tag, id: %s\n", tag.Id())
	default:
		fmt.Printf("Unknown tag, type %T", tag)
	}
}
