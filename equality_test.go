// Copyright 2013 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package names

import (
	gc "gopkg.in/check.v1"
)

var tagEqualityTests = []struct {
	expected Tag
	want     Tag
}{
	{NewMachineTag("0"), MachineTag{id: "0"}},
	{NewMachineTag("10/lxc/1"), MachineTag{id: "10-lxc-1"}},
	{NewUnitTag("mysql/1"), UnitTag{name: "mysql-1"}},
	{NewServiceTag("ceph"), ServiceTag{Name: "ceph"}},
	{NewRelationTag("wordpress:haproxy"), RelationTag{key: "wordpress.haproxy"}},
	{NewEnvironTag("local"), EnvironTag{uuid: "local"}},
	{NewUserTag("admin"), UserTag{name: "admin"}},
	{NewUserTag("admin@local"), UserTag{name: "admin", provider: "local"}},
	{NewUserTag("admin@foobar"), UserTag{name: "admin", provider: "foobar"}},
	{NewNetworkTag("eth0"), NetworkTag{name: "eth0"}},
	{NewActionTag("01234567-aaaa-bbbb-cccc-012345678901"), ActionTag{"01234567-aaaa-bbbb-cccc-012345678901"}},
}

type equalitySuite struct{}

var _ = gc.Suite(&equalitySuite{})

func (s *equalitySuite) TestTagEquality(c *gc.C) {
	for _, tt := range tagEqualityTests {
		c.Check(tt.want, gc.Equals, tt.expected)
	}
}
