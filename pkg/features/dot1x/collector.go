// SPDX-License-Identifier: MIT

package dot1x

import (
	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_dot1x_"

var (
	currAuthStateDesc    *prometheus.Desc
	currAuthMethodeDesc  *prometheus.Desc
	currAuthVlanDesc     *prometheus.Desc
	currAuthVoipVlanDesc *prometheus.Desc
)

func init() {
	l := []string{"target", "interface_name", "user_mac_address", "user_name"}
	currAuthStateDesc = prometheus.NewDesc(prefix+"auth_state", "Interface dot1x Authentication State 1: Authenticated, 2: Initialize, 3: Connecting, 4: Held", l, nil)
	currAuthMethodeDesc = prometheus.NewDesc(prefix+"auth_method", "Interface dot1x Authentication Method 1: Radius, 2: Mac Radius, 3: None, 4: Fail", l, nil)
	currAuthVlanDesc = prometheus.NewDesc(prefix+"authenticated_vlan", "Interface dot1x Authentication ", l, nil)
	currAuthVoipVlanDesc = prometheus.NewDesc(prefix+"authenticated_voip_vlan", "Interface dot1x Authenticated Voip Vlan", l, nil)
}

type dot1xCollector struct{}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &dot1xCollector{}
}

// Name returns the name of the collector
func (*dot1xCollector) Name() string {
	return "dot1x"
}

// Describe describes the metrics
func (*dot1xCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- currAuthStateDesc
	ch <- currAuthMethodeDesc
	ch <- currAuthVlanDesc
	ch <- currAuthVoipVlanDesc
}

func dot1xInterfaceState(State string) float64 {
	switch State {
	case "Authenticated":
		return 1
	case "Initialize":
		return 2
	case "Connecting":
		return 3
	case "Held":
		return 4
	default:
		return 0
	}
}

func dot1xInterfaceAuthMethhod(State string) float64 {
	switch State {
	case "Radius":
		return 1
	case "Mac Radius":
		return 2
	case "None":
		return 3
	case "Fail":
		return 4
	default:
		return 0
	}
}

// Collect collects metrics from JunOS
func (c *dot1xCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	err := c.collect(client, ch, labelValues)
	if err != nil {
		return err
	}

	return nil
}

func (c *dot1xCollector) collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = result{}

	err := client.RunCommandAndParse("show dot1x interface extensive", &x)
	if err != nil {
		return err
	}
	for _, dot1xInterface := range x.Results.Interfaces {
		c.collectForInterface(dot1xInterface, ch, labelValues)
	}

	return nil
}

func (c *dot1xCollector) collectForInterface(p dot1xInterface, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, []string{p.InterfaceName, p.UserMacAddress, p.UserName}...)
	ch <- prometheus.MustNewConstMetric(currAuthStateDesc, prometheus.GaugeValue, dot1xInterfaceState(p.State), l...)
	ch <- prometheus.MustNewConstMetric(currAuthMethodeDesc, prometheus.GaugeValue, dot1xInterfaceAuthMethhod(p.AuthenticatedMethod), l...)
	ch <- prometheus.MustNewConstMetric(currAuthVlanDesc, prometheus.GaugeValue, float64(p.AuthenticatedVlan), l...)
	ch <- prometheus.MustNewConstMetric(currAuthVoipVlanDesc, prometheus.GaugeValue, float64(p.AuthenticatedVoipVlan), l...)
}
