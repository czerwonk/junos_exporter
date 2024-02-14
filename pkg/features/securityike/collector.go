// SPDX-License-Identifier: MIT

package securityike

import (
	"encoding/xml"
	"strconv"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_security_ike_"

var (
	connectedActiveUsers *prometheus.Desc
)

func init() {
	l := []string{"target", "re_name"}

	connectedActiveUsers = prometheus.NewDesc(prefix+"connected_active_users", "Number of connected active users", append(l, "remote_address", "remote_port", "ike_id", "x_auth_username", "x_auth_user_assigned_ip"), nil)
}

type securityIKECollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &securityIKECollector{}
}

// Name returns the name of the collector
func (*securityIKECollector) Name() string {
	return "Security IKE"
}

// Describe describes the metrics
func (*securityIKECollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- connectedActiveUsers
}

// Collect collects metrics from JunOS
func (c *securityIKECollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = multiEngineResult{}
	err := client.RunCommandAndParseWithParser("show security ike active-peer", func(b []byte) error {
		return parseXML(b, &x)
	})
	if err != nil {
		return err
	}

	for _, re := range x.Results.RoutingEngines {
		ls := append(labelValues, re.Name)
		activePeersCounters := make(map[string]int)
		for _, ap := range re.IKEActivePeersInformation.IKEActivePeers {
			saRemotePort := strconv.Itoa(ap.IKESARemotePort)
			key := ap.IKESARemoteAddress + saRemotePort + ap.IKEIKEID + ap.IKEXAuthUsername + ap.IKEXAuthUserAssignedIP
			if _, exists := activePeersCounters[key]; !exists {
				activePeersCounters[key] = 0
			}
			activePeersCounters[key] += 1
			ch <- prometheus.MustNewConstMetric(connectedActiveUsers, prometheus.GaugeValue, float64(activePeersCounters[key]), append(ls, ap.IKESARemoteAddress, saRemotePort, ap.IKEIKEID, ap.IKEXAuthUsername, ap.IKEXAuthUserAssignedIP)...)
		}
	}

	return err
}

func parseXML(b []byte, res *multiEngineResult) error {
	if strings.Contains(string(b), "multi-routing-engine-results") {
		return xml.Unmarshal(b, res)
	}

	fi := singleEngineResult{}

	err := xml.Unmarshal(b, &fi)
	if err != nil {
		return err
	}

	res.Results.RoutingEngines = []routingEngine{
		{
			Name:                      "N/A",
			IKEActivePeersInformation: fi.IKEActivePeersInformation,
		},
	}
	return nil
}
