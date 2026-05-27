// SPDX-License-Identifier: MIT

package cluster

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseChassisClusterStatus(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/23.4R2-S2.1/junos">
    <chassis-cluster-status>
        <cluster-id>17</cluster-id>
        <redundancy-group>
            <cluster-id>17</cluster-id>
            <redundancy-group-id>0</redundancy-group-id>
            <redundancy-group-failover-count>1</redundancy-group-failover-count>
            <device-stats>
                <device-name>node0</device-name>
                <device-priority>200</device-priority>
                <redundancy-group-status>primary</redundancy-group-status>
                <device-name>node1</device-name>
                <device-priority>150</device-priority>
                <redundancy-group-status>secondary</redundancy-group-status>
            </device-stats>
        </redundancy-group>
        <redundancy-group>
            <cluster-id>17</cluster-id>
            <redundancy-group-id>1</redundancy-group-id>
            <redundancy-group-failover-count>3</redundancy-group-failover-count>
            <device-stats>
                <device-name>node0</device-name>
                <device-priority>200</device-priority>
                <redundancy-group-status>secondary</redundancy-group-status>
                <device-name>node1</device-name>
                <device-priority>150</device-priority>
                <redundancy-group-status>primary</redundancy-group-status>
            </device-stats>
        </redundancy-group>
    </chassis-cluster-status>
</rpc-reply>`

	var rpc chassisClusterResult
	err := xml.Unmarshal([]byte(body), &rpc)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 17, rpc.Status.ClusterID, "cluster-id")
	assert.Len(t, rpc.Status.RedundancyGroups, 2)

	rg0 := rpc.Status.RedundancyGroups[0]
	assert.Equal(t, 0, rg0.RedundancyGroupID, "redundancy-group-id")
	assert.Equal(t, 1, rg0.FailoverCount, "redundancy-group-failover-count")
	assert.Equal(t, []string{"node0", "node1"}, rg0.DeviceStats.DeviceNames, "device-name")
	assert.Equal(t, []int{200, 150}, rg0.DeviceStats.Priorities, "device-priority")
	assert.Equal(t, []string{"primary", "secondary"}, rg0.DeviceStats.Statuses, "redundancy-group-status")

	rg1 := rpc.Status.RedundancyGroups[1]
	assert.Equal(t, 1, rg1.RedundancyGroupID, "redundancy-group-id")
	assert.Equal(t, 3, rg1.FailoverCount, "redundancy-group-failover-count")
	assert.Equal(t, []string{"secondary", "primary"}, rg1.DeviceStats.Statuses, "redundancy-group-status")
}
