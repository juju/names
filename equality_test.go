// Copyright 2013 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	gc "launchpad.net/gocheck"
)

var tagEqualityTests = []struct {
	expected Tag
	want     Tag
}{
	{NewMachineTag("0"), MachineTag{id: "0"}},
	{NewMachineTag("10/lxc/1"), MachineTag{id: "10-lxc-1"}},
	{NewUnitTag("mysql/1"), UnitTag{name: "mysql-1"}},
	{NewServiceTag("ceph"), ServiceTag{name: "ceph"}},
	{NewRelationTag("wordpress:haproxy"), RelationTag{key: "wordpress.haproxy"}},
	{NewEnvironTag("local"), EnvironTag{uuid: "local"}},
	{NewUserTag("admin"), UserTag{name: "admin"}},
	{NewUserTag("admin@local"), UserTag{name: "admin", provider: "local"}},
	{NewUserTag("admin@foobar"), UserTag{name: "admin", provider: "foobar"}},
	{NewNetworkTag("eth0"), NetworkTag{name: "eth0"}},
	{NewActionTag("foo" + actionMarker + "321"), makeActionTag("foo", "321")},
	{NewActionTag("foo/0" + actionMarker + "321"), makeActionTag("foo/0", "321")},
	{NewActionResultTag("foo" + actionResultMarker + "321"), makeActionResultTag("foo", "321")},
	{NewActionResultTag("foo/0" + actionResultMarker + "321"), makeActionResultTag("foo/0", "321")},
}

type equalitySuite struct{}

var _ = gc.Suite(&equalitySuite{})

func (s *equalitySuite) TestTagEquality(c *gc.C) {
	for _, tt := range tagEqualityTests {
		c.Check(tt.want, gc.Equals, tt.expected)
	}
}

func makeActionTag(prefix, suffix string) ActionTag {
	id := prefix + ActionMarker + suffix
	return ActionTag{idPrefixer: makePrefixer(id, ActionTagKind, ActionMarker)}
}
func makeActionResultTag(prefix, suffix string) ActionResultTag {
	id := prefix + ActionResultMarker + suffix
	return ActionResultTag{idPrefixer: makePrefixer(id, ActionResultTagKind, ActionResultMarker)}
}
func makePrefixer(id, kind, marker string) idPrefixer {
	return idPrefixer{id: id, kind: kind, marker: marker}
}
