// SPDX-License-Identifier: MIT

package bgp

import (
	"math"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"

	"strings"
)

const prefix string = "junos_bgp_session_"

var (
	upDesc                      *prometheus.Desc
	stateDesc                   *prometheus.Desc
	receivedPrefixesDesc        *prometheus.Desc
	acceptedPrefixesDesc        *prometheus.Desc
	rejectedPrefixesDesc        *prometheus.Desc
	activePrefixesDesc          *prometheus.Desc
	advertisedPrefixesDesc      *prometheus.Desc
	inputMessagesDesc           *prometheus.Desc
	outputMessagesDesc          *prometheus.Desc
	flapsDesc                   *prometheus.Desc
	prefixesLimitCountDesc      *prometheus.Desc
	prefixesLimitPercentageDesc *prometheus.Desc
)

func init() {
	l := []string{"target", "asn", "ip", "description", "group"}
	upDesc = prometheus.NewDesc(prefix+"up", "Session is up (1 = Established)", l, nil)
	stateDesc = prometheus.NewDesc(prefix+"state", "State of the bgp Session (1 = Active, 2 = Connect, 3 = Established, 4 = Idle, 5 = OpenConfirm, 6 = OpenSent, 7 = route reflector client, 0 = Other)", l, nil)
	inputMessagesDesc = prometheus.NewDesc(prefix+"messages_input_count", "Number of received messages", l, nil)
	outputMessagesDesc = prometheus.NewDesc(prefix+"messages_output_count", "Number of transmitted messages", l, nil)
	flapsDesc = prometheus.NewDesc(prefix+"flap_count", "Number of session flaps", l, nil)

	l = append(l, "table")

	receivedPrefixesDesc = prometheus.NewDesc(prefix+"prefixes_received_count", "Number of received prefixes", l, nil)
	acceptedPrefixesDesc = prometheus.NewDesc(prefix+"prefixes_accepted_count", "Number of accepted prefixes", l, nil)
	rejectedPrefixesDesc = prometheus.NewDesc(prefix+"prefixes_rejected_count", "Number of rejected prefixes", l, nil)
	activePrefixesDesc = prometheus.NewDesc(prefix+"prefixes_active_count", "Number of active prefixes (best route in RIB)", l, nil)
	advertisedPrefixesDesc = prometheus.NewDesc(prefix+"prefixes_advertised_count", "Number of prefixes announced to peer", l, nil)
	prefixesLimitPercentageDesc = prometheus.NewDesc(prefix+"prefixes_limit_percentage", "percentage of received prefixes against prefix-limit", l, nil)
	prefixesLimitCountDesc = prometheus.NewDesc(prefix+"prefixes_limit_count", "prefix-count variable set in prefix-limit", l, nil)
}

type bgpCollector struct {
	LogicalSystem string
}

// NewCollector creates a new collector
func NewCollector(logicalSystem string) collector.RPCCollector {
	return &bgpCollector{LogicalSystem: logicalSystem}
}

// Name returns the name of the collector
func (*bgpCollector) Name() string {
	return "BGP"
}

// Describe describes the metrics
func (*bgpCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- upDesc
	ch <- receivedPrefixesDesc
	ch <- acceptedPrefixesDesc
	ch <- rejectedPrefixesDesc
	ch <- activePrefixesDesc
	ch <- advertisedPrefixesDesc
	ch <- inputMessagesDesc
	ch <- outputMessagesDesc
	ch <- flapsDesc
	ch <- prefixesLimitCountDesc
	ch <- prefixesLimitPercentageDesc
}

// Collect collects metrics from JunOS
func (c *bgpCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	err := c.collect(client, ch, labelValues)
	if err != nil {
		return err
	}

	return nil
}

func (c *bgpCollector) collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = information{}

	var cmd string

	if c.LogicalSystem == "" {
		cmd = "show bgp neighbor"
	} else {
		cmd = "show bgp neighbor logical-system" + c.LogicalSystem
	}

	err := client.RunCommandAndParse(cmd, &x)
	if err != nil {
		return err
	}

	for _, peer := range x.Peers {
		c.collectForPeer(peer, ch, labelValues)
	}

	return nil
}

func (c *bgpCollector) collectForPeer(p peer, ch chan<- prometheus.Metric, labelValues []string) {
	ip := strings.Split(p.IP, "+")
	l := append(labelValues, []string{p.ASN, ip[0], p.Description, p.Group}...)

	up := 0
	if p.State == "Established" {
		up = 1
	}

	ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, float64(up), l...)
	ch <- prometheus.MustNewConstMetric(stateDesc, prometheus.GaugeValue, bgpStateToNumber(p.State), l...)
	ch <- prometheus.MustNewConstMetric(inputMessagesDesc, prometheus.GaugeValue, float64(p.InputMessages), l...)
	ch <- prometheus.MustNewConstMetric(outputMessagesDesc, prometheus.GaugeValue, float64(p.OutputMessages), l...)
	ch <- prometheus.MustNewConstMetric(flapsDesc, prometheus.GaugeValue, float64(p.Flaps), l...)

	c.collectRIBForPeer(p, ch, l)
}

func (*bgpCollector) collectRIBForPeer(p peer, ch chan<- prometheus.Metric, labelValues []string) {
	var rib_name string

	// derive the name of the rib for which the prefix limit is configured by examining the NLRI type
	switch nlri_type := p.BGPOI.PrefixLimit.NlriType; nlri_type {
	case "inet-unicast":
		rib_name = "inet.0"
	case "inet6-unicast":
		rib_name = "inet6.0"
	default:
		rib_name = ""
	}

	// if the prefix limit is configured inside a routing instance we need to prepend the RTI name to the rib name
	if p.CFGRTI != "" && p.CFGRTI != "master" && rib_name != "" {
		rib_name = p.CFGRTI + "." + rib_name
	}

	if p.BGPOI.PrefixLimit.PrefixCount > 0 {
		ch <- prometheus.MustNewConstMetric(prefixesLimitCountDesc, prometheus.GaugeValue, float64(p.BGPOI.PrefixLimit.PrefixCount), append(labelValues, rib_name)...)
	}

	for _, rib := range p.RIBs {
		l := append(labelValues, rib.Name)
		ch <- prometheus.MustNewConstMetric(receivedPrefixesDesc, prometheus.GaugeValue, float64(rib.ReceivedPrefixes), l...)
		ch <- prometheus.MustNewConstMetric(acceptedPrefixesDesc, prometheus.GaugeValue, float64(rib.AcceptedPrefixes), l...)
		ch <- prometheus.MustNewConstMetric(rejectedPrefixesDesc, prometheus.GaugeValue, float64(rib.RejectedPrefixes), l...)
		ch <- prometheus.MustNewConstMetric(activePrefixesDesc, prometheus.GaugeValue, float64(rib.ActivePrefixes), l...)
		ch <- prometheus.MustNewConstMetric(advertisedPrefixesDesc, prometheus.GaugeValue, float64(rib.AdvertisedPrefixes), l...)

		if rib.Name == rib_name {
			if p.BGPOI.PrefixLimit.PrefixCount > 0 {
				prefixesLimitPercent := float64(rib.ReceivedPrefixes) / float64(p.BGPOI.PrefixLimit.PrefixCount)
				ch <- prometheus.MustNewConstMetric(prefixesLimitPercentageDesc, prometheus.GaugeValue, math.Round(prefixesLimitPercent*100)/100, l...)
			}
		}
	}
}

func bgpStateToNumber(bgpState string) float64 {
	switch bgpState {
	case "Active":
		return 1
	case "Connect":
		return 2
	case "Established":
		return 3
	case "Idle":
		return 4
	case "Openconfirm":
		return 5
	case "OpenSent":
		return 6
	case "route reflector client":
		return 7
	default:
		return 0
	}
}
