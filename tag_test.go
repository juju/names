// Copyright 2015 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names_test

import (
	"strconv"

	"github.com/juju/names/v4"
	gc "gopkg.in/check.v1"
)

type tagSuite struct{}

var _ = gc.Suite(&tagSuite{})

var tagKindTests = []struct {
	tag  string
	kind string
	err  string
}{
	{tag: "unit-wordpress-42", kind: names.UnitTagKind},
	{tag: "machine-42", kind: names.MachineTagKind},
	{tag: "application-foo", kind: names.ApplicationTagKind},
	{tag: "applicationoffer-foo", kind: names.ApplicationOfferTagKind},
	{tag: "environment-42", kind: names.EnvironTagKind},
	{tag: "model-42", kind: names.ModelTagKind},
	{tag: "user-admin", kind: names.UserTagKind},
	{tag: "relation-application1.rel1#other-svc.other-rel2", kind: names.RelationTagKind},
	{tag: "relation-application.peerRelation", kind: names.RelationTagKind},
	{tag: "foo", err: `"foo" is not a valid tag`},
	{tag: "unit", err: `"unit" is not a valid tag`},
	{tag: "network", err: `"network" is not a valid tag`},
	{tag: "ab01cd23-0123-4edc-9a8b-fedcba987654", err: `"ab01cd23-0123-4edc-9a8b-fedcba987654" is not a valid tag`},
	{tag: "action-ab01cd23-0123-4edc-9a8b-fedcba987654", kind: names.ActionTagKind},
	{tag: "action-1", kind: names.ActionTagKind},
	{tag: "operation-1", kind: names.OperationTagKind},
	{tag: "volume-0", kind: names.VolumeTagKind},
	{tag: "storage-data-0", kind: names.StorageTagKind},
	{tag: "filesystem-0", kind: names.FilesystemTagKind},
	{tag: "ipaddress", err: `"ipaddress" is not a valid tag`},
	{tag: "ipaddress-42424242-1111-2222-3333-0123456789ab", kind: names.IPAddressTagKind},
	{tag: "subnet", err: `"subnet" is not a valid tag`},
	{tag: "subnet-16", kind: names.SubnetTagKind},
	{tag: "space", err: `"space" is not a valid tag`},
	{tag: "space-42", kind: names.SpaceTagKind},
	{tag: "cloud", err: `"cloud" is not a valid tag`},
	{tag: "cloud-aws", kind: names.CloudTagKind},
	{tag: "cloudcred", err: `"cloudcred" is not a valid tag`},
	{tag: "cloudcred-aws_admin_foo", kind: names.CloudCredentialTagKind},
	{tag: "caasmodel-57", kind: names.CAASModelTagKind},
	{tag: "controller-f47ac10b-58cc-4372-a567-0e02b2c3d479", kind: names.ControllerTagKind},
	{tag: "controller-123", kind: names.ControllerAgentTagKind},
}

func (*tagSuite) TestTagKind(c *gc.C) {
	for i, test := range tagKindTests {
		c.Logf("test %d: %q -> %q", i, test.tag, test.kind)
		kind, err := names.TagKind(test.tag)
		if test.err == "" {
			c.Assert(test.kind, gc.Equals, kind)
			c.Assert(err, gc.IsNil)
		} else {
			c.Assert(kind, gc.Equals, "")
			c.Assert(err, gc.ErrorMatches, test.err)
		}
	}
}

var parseTagTests = []struct {
	tag        string
	expectKind string
	expectType interface{}
	resultId   string
	resultErr  string
}{{
	tag:        "machine-10",
	expectKind: names.MachineTagKind,
	expectType: names.MachineTag{},
	resultId:   "10",
}, {
	tag:        "machine-10-lxc-1",
	expectKind: names.MachineTagKind,
	expectType: names.MachineTag{},
	resultId:   "10/lxc/1",
}, {
	tag:        "machine-#",
	expectKind: names.MachineTagKind,
	expectType: names.MachineTag{},
	resultErr:  `"machine-#" is not a valid tag of kind machine`,
}, {
	tag:        "unit-wordpress-0",
	expectKind: names.UnitTagKind,
	expectType: names.UnitTag{},
	resultId:   "wordpress/0",
}, {
	tag:        "unit-rabbitmq-server-0",
	expectKind: names.UnitTagKind,
	expectType: names.UnitTag{},
	resultId:   "rabbitmq-server/0",
}, {
	tag:        "unit-#",
	expectKind: names.UnitTagKind,
	expectType: names.UnitTag{},
	resultErr:  `"unit-#" is not a valid tag of kind unit`,
}, {
	tag:        "application-wordpress",
	expectKind: names.ApplicationTagKind,
	expectType: names.ApplicationTag{},
	resultId:   "wordpress",
}, {
	tag:        "application-#",
	expectKind: names.ApplicationTagKind,
	expectType: names.ApplicationTag{},
	resultErr:  `"application-#" is not a valid tag of kind application`,
}, {
	tag:        "applicationoffer-hosted-mysql",
	expectKind: names.ApplicationOfferTagKind,
	expectType: names.ApplicationOfferTag{},
	resultId:   "hosted-mysql",
}, {
	tag:        "applicationoffer-#",
	expectKind: names.ApplicationOfferTagKind,
	expectType: names.ApplicationOfferTag{},
	resultErr:  `"applicationoffer-#" is not a valid tag of kind applicationoffer`,
}, {
	tag:        "environment-f47ac10b-58cc-4372-a567-0e02b2c3d479",
	expectKind: names.EnvironTagKind,
	expectType: names.EnvironTag{},
	resultId:   "f47ac10b-58cc-4372-a567-0e02b2c3d479",
}, {
	tag:        "model-f47ac10b-58cc-4372-a567-0e02b2c3d479",
	expectKind: names.ModelTagKind,
	expectType: names.ModelTag{},
	resultId:   "f47ac10b-58cc-4372-a567-0e02b2c3d479",
}, {
	tag:        "relation-my-svc1.myrel1#other-svc.other-rel2",
	expectKind: names.RelationTagKind,
	expectType: names.RelationTag{},
	resultId:   "my-svc1:myrel1 other-svc:other-rel2",
}, {
	tag:        "relation-riak.ring",
	expectKind: names.RelationTagKind,
	expectType: names.RelationTag{},
	resultId:   "riak:ring",
}, {
	tag:        "environment-/",
	expectKind: names.EnvironTagKind,
	expectType: names.EnvironTag{},
	resultErr:  `"environment-/" is not a valid tag of kind environment`,
}, {
	tag:        "model-/",
	expectKind: names.ModelTagKind,
	expectType: names.ModelTag{},
	resultErr:  `"model-/" is not a valid tag of kind model`,
}, {
	tag:        "user-foo",
	expectKind: names.UserTagKind,
	expectType: names.UserTag{},
	resultId:   "foo",
}, {
	tag:        "user-foo@remote",
	expectKind: names.UserTagKind,
	expectType: names.UserTag{},
	resultId:   "foo@remote",
}, {
	tag:        "user-/",
	expectKind: names.UserTagKind,
	expectType: names.UserTag{},
	resultErr:  `"user-/" is not a valid tag of kind user`,
}, {
	tag:        "action-00000000-abcd",
	expectKind: names.ActionTagKind,
	expectType: names.ActionTag{},
	resultErr:  `"action-00000000-abcd" is not a valid tag of kind action`,
}, {
	tag:        "action-00000033",
	expectKind: names.ActionTagKind,
	expectType: names.ActionTag{},
	resultErr:  `"action-00000033" is not a valid tag of kind action`,
}, {
	tag:        "action-abedaf33-3212-4fde-aeca-87356432deca",
	expectKind: names.ActionTagKind,
	expectType: names.ActionTag{},
	resultId:   "abedaf33-3212-4fde-aeca-87356432deca",
}, {
	tag:        "action-1",
	expectKind: names.ActionTagKind,
	expectType: names.ActionTag{},
	resultId:   "1",
}, {
	tag:        "operation-1",
	expectKind: names.OperationTagKind,
	expectType: names.OperationTag{},
	resultId:   "1",
}, {
	tag:        "volume-2",
	expectKind: names.VolumeTagKind,
	expectType: names.VolumeTag{},
	resultId:   "2",
}, {
	tag:        "filesystem-3",
	expectKind: names.FilesystemTagKind,
	expectType: names.FilesystemTag{},
	resultId:   "3",
}, {
	tag:        "storage-block-storage-0",
	expectKind: names.StorageTagKind,
	expectType: names.StorageTag{},
	resultId:   "block-storage/0",
}, {
	tag:       "foo",
	resultErr: `"foo" is not a valid tag`,
}, {
	tag:       "ipaddress-",
	resultErr: `"ipaddress-" is not a valid tag of kind ipaddress`,
}, {
	tag:        "ipaddress-42424242-1111-2222-3333-0123456789ab",
	expectKind: names.IPAddressTagKind,
	expectType: names.IPAddressTag{},
	resultId:   "42424242-1111-2222-3333-0123456789ab",
}, {
	tag:       "subnet-",
	resultErr: `"subnet-" is not a valid tag of kind subnet`,
}, {
	tag:        "subnet-16",
	expectKind: names.SubnetTagKind,
	expectType: names.SubnetTag{},
	resultId:   "16",
}, {
	tag:       "space-",
	resultErr: `"space-" is not a valid tag of kind space`,
}, {
	tag:        "space-myspace1",
	expectKind: names.SpaceTagKind,
	expectType: names.SpaceTag{},
	resultId:   "myspace1",
}, {
	tag:        "cloud-aws",
	expectKind: names.CloudTagKind,
	expectType: names.CloudTag{},
	resultId:   "aws",
}, {
	tag:        "cloudcred-aws_admin_foo%5fbar",
	expectKind: names.CloudCredentialTagKind,
	expectType: names.CloudCredentialTag{},
	resultId:   "aws/admin/foo_bar",
}, {
	tag:        "caasmodel-f47ac10b-58cc-4372-a567-0e02b2c3d479",
	expectKind: names.CAASModelTagKind,
	expectType: names.CAASModelTag{},
	resultId:   "f47ac10b-58cc-4372-a567-0e02b2c3d479",
}, {
	tag:        "caasmodel-/",
	expectKind: names.CAASModelTagKind,
	expectType: names.CAASModelTag{},
	resultErr:  `"caasmodel-/" is not a valid tag of kind caasmodel`,
}, {
	tag:        "controller-f47ac10b-58cc-4372-a567-0e02b2c3d479",
	expectKind: names.ControllerTagKind,
	expectType: names.ControllerTag{},
	resultId:   "f47ac10b-58cc-4372-a567-0e02b2c3d479",
}, {
	tag:        "controller-123",
	expectKind: names.ControllerAgentTagKind,
	expectType: names.ControllerAgentTag{},
	resultId:   "123",
}, {
	tag:        "controller-invalid",
	expectKind: names.ControllerTagKind,
	expectType: names.ControllerTag{},
	resultErr:  `"controller-invalid" is not a valid tag of kind controller`,
}}

var makeTag = map[string]func(string) names.Tag{
	names.MachineTagKind:          func(tag string) names.Tag { return names.NewMachineTag(tag) },
	names.UnitTagKind:             func(tag string) names.Tag { return names.NewUnitTag(tag) },
	names.ApplicationTagKind:      func(tag string) names.Tag { return names.NewApplicationTag(tag) },
	names.ApplicationOfferTagKind: func(tag string) names.Tag { return names.NewApplicationOfferTag(tag) },
	names.RelationTagKind:         func(tag string) names.Tag { return names.NewRelationTag(tag) },
	names.EnvironTagKind:          func(tag string) names.Tag { return names.NewEnvironTag(tag) },
	names.ModelTagKind:            func(tag string) names.Tag { return names.NewModelTag(tag) },
	names.UserTagKind:             func(tag string) names.Tag { return names.NewUserTag(tag) },
	names.ActionTagKind:           func(tag string) names.Tag { return names.NewActionTag(tag) },
	names.OperationTagKind:        func(tag string) names.Tag { return names.NewOperationTag(tag) },
	names.VolumeTagKind:           func(tag string) names.Tag { return names.NewVolumeTag(tag) },
	names.FilesystemTagKind:       func(tag string) names.Tag { return names.NewFilesystemTag(tag) },
	names.StorageTagKind:          func(tag string) names.Tag { return names.NewStorageTag(tag) },
	names.IPAddressTagKind:        func(tag string) names.Tag { return names.NewIPAddressTag(tag) },
	names.SubnetTagKind:           func(tag string) names.Tag { return names.NewSubnetTag(tag) },
	names.SpaceTagKind:            func(tag string) names.Tag { return names.NewSpaceTag(tag) },
	names.CloudTagKind:            func(tag string) names.Tag { return names.NewCloudTag(tag) },
	names.CloudCredentialTagKind:  func(tag string) names.Tag { return names.NewCloudCredentialTag(tag) },
	names.CAASModelTagKind:        func(tag string) names.Tag { return names.NewCAASModelTag(tag) },
	names.ControllerTagKind: func(tag string) names.Tag {
		_, err := strconv.Atoi(tag)
		if err == nil {
			return names.NewControllerAgentTag(tag)
		}
		return names.NewControllerTag(tag)
	},
}

func (*tagSuite) TestParseTag(c *gc.C) {
	for i, test := range parseTagTests {
		c.Logf("test %d: %q expectKind %q", i, test.tag, test.expectKind)
		tag, err := names.ParseTag(test.tag)
		if test.resultErr != "" {
			c.Assert(err, gc.ErrorMatches, test.resultErr)
			c.Assert(tag, gc.IsNil)

			// If the tag has a valid kind which matches the
			// expected kind, test that using an empty
			// expectKind does not change the error message.
			if tagKind, err := names.TagKind(test.tag); err == nil && tagKind == test.expectKind {
				tag, err := names.ParseTag(test.tag)
				c.Assert(err, gc.ErrorMatches, test.resultErr)
				c.Assert(tag, gc.IsNil)
			}
		} else {
			c.Assert(err, gc.IsNil)
			kind, id := tag.Kind(), tag.Id()
			c.Assert(err, gc.IsNil)
			c.Assert(id, gc.Equals, test.resultId)
			if test.expectKind != "" {
				c.Assert(kind, gc.Equals, test.expectKind)
			} else {
				expectKind, err := names.TagKind(test.tag)
				c.Assert(err, gc.IsNil)
				c.Assert(kind, gc.Equals, expectKind) // will be removed in the next branch
				c.Assert(tag, gc.FitsTypeOf, test.expectType)
			}
			// Check that it's reversible.
			if f := makeTag[kind]; f != nil {
				reversed := f(id).String()
				c.Assert(reversed, gc.Equals, test.tag)
			}
			// Check that it parses ok without an expectKind.
			tag, err := names.ParseTag(test.tag)
			c.Assert(err, gc.IsNil)
			c.Assert(tag, gc.FitsTypeOf, test.expectType)
			c.Assert(tag.Kind(), gc.Equals, test.expectKind) // will be removed in the next branch
			c.Assert(tag.Id(), gc.Equals, id)
		}
	}
}

func (*tagSuite) TestReadableString(c *gc.C) {
	var readableStringTests = []struct {
		tag    names.Tag
		result string
	}{{
		tag:    nil,
		result: "",
	}, {
		tag:    names.NewMachineTag("0"),
		result: "machine 0",
	}, {
		tag:    names.NewUnitTag("wordpress/2"),
		result: "unit wordpress/2",
	}}

	for i, test := range readableStringTests {
		c.Logf("test %d: expected result %q", i, test.result)
		resultStr := names.ReadableString(test.tag)
		c.Assert(resultStr, gc.Equals, test.result)
	}
}
