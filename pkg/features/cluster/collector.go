// SPDX-License-Identifier: MIT

package cluster

import (
	"strconv"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_chassis_cluster_"

var (
	nodeStatusDesc      *prometheus.Desc
	nodePriorityDesc    *prometheus.Desc
	failoverCountDesc   *prometheus.Desc
)

func init() {
	l := []string{"target", "cluster_id", "redundancy_group", "node", "status"}
	nodeStatusDesc = prometheus.NewDesc(prefix+"node_status", "Chassis cluster redundancy group node status (1 primary, 2 secondary, 3 secondary-hold, 4 disabled, 5 lost, 6 not-configured, 7 ineligible)", l, nil)

	l = []string{"target", "cluster_id", "redundancy_group", "node"}
	nodePriorityDesc = prometheus.NewDesc(prefix+"node_priority", "Chassis cluster redundancy group node priority", l, nil)

	l = []string{"target", "cluster_id", "redundancy_group"}
	failoverCountDesc = prometheus.NewDesc(prefix+"failover_count", "Chassis cluster redundancy group failover count", l, nil)
}

type chassisClusterCollector struct{}

func NewCollector() collector.RPCCollector {
	return &chassisClusterCollector{}
}

func (*chassisClusterCollector) Name() string {
	return "Chassis Cluster"
}

func (*chassisClusterCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- nodeStatusDesc
	ch <- nodePriorityDesc
	ch <- failoverCountDesc
}

func (c *chassisClusterCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	statusValues := map[string]int{
		"primary":        1,
		"secondary":      2,
		"secondary-hold": 3,
		"disabled":       4,
		"lost":           5,
		"not-configured": 6,
		"ineligible":     7,
	}

	var x chassisClusterResult
	err := client.RunCommandAndParse("show chassis cluster status", &x)
	if err != nil {
		return nil
	}

	if len(x.Status.RedundancyGroups) == 0 {
		return nil
	}

	clusterID := strconv.Itoa(x.Status.ClusterID)

	for _, rg := range x.Status.RedundancyGroups {
		rgID := strconv.Itoa(rg.RedundancyGroupID)

		ch <- prometheus.MustNewConstMetric(failoverCountDesc, prometheus.CounterValue, float64(rg.FailoverCount), append(labelValues, clusterID, rgID)...)

		for i, name := range rg.DeviceStats.DeviceNames {
			if i >= len(rg.DeviceStats.Statuses) || i >= len(rg.DeviceStats.Priorities) {
				break
			}
			status := rg.DeviceStats.Statuses[i]
			l := append(labelValues, clusterID, rgID, name)
			ch <- prometheus.MustNewConstMetric(nodeStatusDesc, prometheus.GaugeValue, float64(statusValues[status]), append(l, status)...)
			ch <- prometheus.MustNewConstMetric(nodePriorityDesc, prometheus.GaugeValue, float64(rg.DeviceStats.Priorities[i]), l...)
		}
	}

	return nil
}
