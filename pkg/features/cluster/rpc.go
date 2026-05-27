// SPDX-License-Identifier: MIT

package cluster

import "encoding/xml"

type chassisClusterResult struct {
	XMLName xml.Name             `xml:"rpc-reply"`
	Status  chassisClusterStatus `xml:"chassis-cluster-status"`
}

type chassisClusterStatus struct {
	ClusterID        int               `xml:"cluster-id"`
	RedundancyGroups []redundancyGroup `xml:"redundancy-group"`
}

type redundancyGroup struct {
	RedundancyGroupID int         `xml:"redundancy-group-id"`
	FailoverCount     int         `xml:"redundancy-group-failover-count"`
	DeviceStats       deviceStats `xml:"device-stats"`
}

type deviceStats struct {
	DeviceNames []string `xml:"device-name"`
	Priorities  []int    `xml:"device-priority"`
	Statuses    []string `xml:"redundancy-group-status"`
}
