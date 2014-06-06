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
	{UnitTag("mysql/1"), unitTag{name: "mysql-1"}},
	{ServiceTag("ceph"), serviceTag{name: "ceph"}},
	{RelationTag("wordpress:haproxy"), relationTag{key: "wordpress.haproxy"}},
	{NewEnvironTag("local"), EnvironTag{uuid: "local"}},
	{UserTag("admin"), userTag{name: "admin"}},
	{NetworkTag("eth0"), networkTag{name: "eth0"}},
}

type equalitySuite struct{}

var _ = gc.Suite(&equalitySuite{})

func (s *equalitySuite) TestTagEquality(c *gc.C) {
	for _, tt := range tagEqualityTests {
		c.Check(tt.want, gc.Equals, tt.expected)
	}
}
