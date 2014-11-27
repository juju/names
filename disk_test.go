// Copyright 2013 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	"fmt"

	gc "gopkg.in/check.v1"

	"github.com/juju/names"
)

type diskSuite struct{}

var _ = gc.Suite(&diskSuite{})

func (s *diskSuite) TestDiskTag(c *gc.C) {
	c.Assert(names.NewDiskTag("1").String(), gc.Equals, "disk-1")
}

func (s *diskSuite) TestDiskNameValidity(c *gc.C) {
	assertDiskNameValid(c, "0")
	assertDiskNameValid(c, "1000")
	assertDiskNameInvalid(c, "-1")
	assertDiskNameInvalid(c, "")
	assertDiskNameInvalid(c, "one")
	assertDiskNameInvalid(c, "#")
}

func (s *diskSuite) TestParseDiskTag(c *gc.C) {
	assertParseDiskTag(c, "disk-0", names.NewDiskTag("0"))
	assertParseDiskTag(c, "disk-88", names.NewDiskTag("88"))
	assertParseDiskTagInvalid(c, "", names.InvalidTagError("", ""))
	assertParseDiskTagInvalid(c, "one", names.InvalidTagError("one", ""))
	assertParseDiskTagInvalid(c, "disk-", names.InvalidTagError("disk-", names.DiskTagKind))
	assertParseDiskTagInvalid(c, "machine-0", names.InvalidTagError("machine-0", names.DiskTagKind))
}

func assertDiskNameValid(c *gc.C, name string) {
	c.Assert(names.IsValidDisk(name), gc.Equals, true)
	names.NewDiskTag(name)
}

func assertDiskNameInvalid(c *gc.C, name string) {
	c.Assert(names.IsValidDisk(name), gc.Equals, false)
	testDiskTag := func() { names.NewDiskTag(name) }
	expect := fmt.Sprintf("%q is not a valid disk name", name)
	c.Assert(testDiskTag, gc.PanicMatches, expect)
}

func assertParseDiskTag(c *gc.C, tag string, expect names.DiskTag) {
	t, err := names.ParseDiskTag(tag)
	c.Assert(err, gc.IsNil)
	c.Assert(t, gc.Equals, expect)
}

func assertParseDiskTagInvalid(c *gc.C, tag string, expect error) {
	_, err := names.ParseDiskTag(tag)
	c.Assert(err, gc.ErrorMatches, expect.Error())
}
