// Copyright 2015 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	"fmt"

	gc "gopkg.in/check.v1"

	"github.com/juju/names"
)

type filesystemSuite struct{}

var _ = gc.Suite(&filesystemSuite{})

func (s *filesystemSuite) TestFilesystemTag(c *gc.C) {
	c.Assert(names.NewFilesystemTag("abc").String(), gc.Equals, "filesystem-abc")
}

func (s *filesystemSuite) TestFilesystemNameValidity(c *gc.C) {
	assertFilesystemNameValid(c, "abc")
	assertFilesystemNameValid(c, "abc-def.123")
	assertFilesystemNameInvalid(c, "-1")
	assertFilesystemNameInvalid(c, "")
	assertFilesystemNameInvalid(c, "#")
}

func (s *filesystemSuite) TestParseFilesystemTag(c *gc.C) {
	assertParseFilesystemTag(c, "filesystem-abc", names.NewFilesystemTag("abc"))
	assertParseFilesystemTag(c, "filesystem-88", names.NewFilesystemTag("88"))
	assertParseFilesystemTagInvalid(c, "", names.InvalidTagError("", ""))
	assertParseFilesystemTagInvalid(c, "one", names.InvalidTagError("one", ""))
	assertParseFilesystemTagInvalid(c, "filesystem-", names.InvalidTagError("filesystem-", names.FilesystemTagKind))
	assertParseFilesystemTagInvalid(c, "machine-0", names.InvalidTagError("machine-0", names.FilesystemTagKind))
}

func assertFilesystemNameValid(c *gc.C, name string) {
	c.Assert(names.IsValidFilesystem(name), gc.Equals, true)
	names.NewFilesystemTag(name)
}

func assertFilesystemNameInvalid(c *gc.C, name string) {
	c.Assert(names.IsValidFilesystem(name), gc.Equals, false)
	testFilesystemTag := func() { names.NewFilesystemTag(name) }
	expect := fmt.Sprintf("%q is not a valid filesystem name", name)
	c.Assert(testFilesystemTag, gc.PanicMatches, expect)
}

func assertParseFilesystemTag(c *gc.C, tag string, expect names.FilesystemTag) {
	t, err := names.ParseFilesystemTag(tag)
	c.Assert(err, gc.IsNil)
	c.Assert(t, gc.Equals, expect)
}

func assertParseFilesystemTagInvalid(c *gc.C, tag string, expect error) {
	_, err := names.ParseFilesystemTag(tag)
	c.Assert(err, gc.ErrorMatches, expect.Error())
}
