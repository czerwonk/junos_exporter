// SPDX-License-Identifier: MIT

package bgp

import (
	"fmt"
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
	infoDesc                    *prometheus.Desc
	medDesc                     *prometheus.Desc
	preferenceDesc              *prometheus.Desc
	holdTimeDesc                *prometheus.Desc
)

func init() {
	l := []string{"target", "asn", "ip", "description", "group"}
	upDesc = prometheus.NewDesc(prefix+"up", "Session is up (1 = Established)", l, nil)
	stateDesc = prometheus.NewDesc(prefix+"state", "State of the bgp Session (1 = Active, 2 = Connect, 3 = Established, 4 = Idle, 5 = OpenConfirm, 6 = OpenSent, 7 = route reflector client, 0 = Other)", l, nil)
	inputMessagesDesc = prometheus.NewDesc(prefix+"messages_input_count", "Number of received messages", l, nil)
	outputMessagesDesc = prometheus.NewDesc(prefix+"messages_output_count", "Number of transmitted messages", l, nil)
	flapsDesc = prometheus.NewDesc(prefix+"flap_count", "Number of session flaps", l, nil)
	medDesc = prometheus.NewDesc(prefix+"metric_out", "MED configured for the session", l, nil)
	preferenceDesc = prometheus.NewDesc(prefix+"preference", "Preference configured for the session", l, nil)
	holdTimeDesc = prometheus.NewDesc(prefix+"hold_time_seconds", "Hold time configured for the session", l, nil)

	infoLabels := append(l, "local_as", "import_policy", "export_policy")
	infoDesc = prometheus.NewDesc(prefix+"info", "Information about the session (e.g. configuration)", infoLabels, nil)

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

type groupMap map[int64]group

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
	ch <- infoDesc
	ch <- medDesc
	ch <- preferenceDesc
	ch <- holdTimeDesc
}

// Collect collects metrics from JunOS
func (c *bgpCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	err := c.collect(client, ch, labelValues)
	if err != nil {
		return err
	}

	return nil
}

func (c *bgpCollector) collectGroups(client collector.Client) (groupMap, error) {
	var x = groupResult{}
	var cmd strings.Builder
	cmd.WriteString("show bgp group")
	if c.LogicalSystem != "" {
		cmd.WriteString(" logical-system " + c.LogicalSystem)
	}

	err := client.RunCommandAndParse(cmd.String(), &x)
	if err != nil {
		return nil, err
	}

	groups := make(groupMap)
	for _, g := range x.Information.Groups {
		groups[g.Index] = g
	}

	return groups, err
}

func (c *bgpCollector) collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	groups, err := c.collectGroups(client)
	if err != nil {
		return fmt.Errorf("could not retrieve BGP group information: %w", err)
	}

	var x = result{}
	var cmd strings.Builder
	cmd.WriteString("show bgp neighbor")
	if c.LogicalSystem != "" {
		cmd.WriteString(" logical-system " + c.LogicalSystem)
	}

	err = client.RunCommandAndParse(cmd.String(), &x)
	if err != nil {
		return err
	}

	for _, peer := range x.Information.Peers {
		c.collectForPeer(peer, groups, ch, labelValues)
	}

	return nil
}

func (c *bgpCollector) collectForPeer(p peer, groups groupMap, ch chan<- prometheus.Metric, labelValues []string) {
	ip := strings.Split(p.IP, "+")
	l := append(labelValues, []string{
		p.ASN,
		ip[0],
		p.Description,
		groupForPeer(p, groups)}...)

	up := 0
	if p.State == "Established" {
		up = 1
	}

	ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, float64(up), l...)
	ch <- prometheus.MustNewConstMetric(stateDesc, prometheus.GaugeValue, bgpStateToNumber(p.State), l...)
	ch <- prometheus.MustNewConstMetric(inputMessagesDesc, prometheus.GaugeValue, float64(p.InputMessages), l...)
	ch <- prometheus.MustNewConstMetric(outputMessagesDesc, prometheus.GaugeValue, float64(p.OutputMessages), l...)
	ch <- prometheus.MustNewConstMetric(flapsDesc, prometheus.GaugeValue, float64(p.Flaps), l...)
	ch <- prometheus.MustNewConstMetric(preferenceDesc, prometheus.GaugeValue, float64(p.OptionInformation.Preference), l...)
	ch <- prometheus.MustNewConstMetric(medDesc, prometheus.GaugeValue, float64(p.OptionInformation.MetricOut), l...)
	ch <- prometheus.MustNewConstMetric(holdTimeDesc, prometheus.GaugeValue, float64(p.OptionInformation.Holdtime), l...)

	infoValues := append(l,
		localASNForPeer(p),
		formatPolicy(p.OptionInformation.ImportPolicy),
		formatPolicy(p.OptionInformation.ExportPolicy))
	ch <- prometheus.MustNewConstMetric(infoDesc, prometheus.GaugeValue, 1, infoValues...)

	c.collectRIBForPeer(p, ch, l)
}

func (*bgpCollector) collectRIBForPeer(p peer, ch chan<- prometheus.Metric, labelValues []string) {
	var rib_name string

	// derive the name of the rib for which the prefix limit is configured by examining the NLRI type
	switch nlri_type := p.OptionInformation.PrefixLimit.NlriType; nlri_type {
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

	if p.OptionInformation.PrefixLimit.PrefixCount > 0 {
		ch <- prometheus.MustNewConstMetric(prefixesLimitCountDesc, prometheus.GaugeValue, float64(p.OptionInformation.PrefixLimit.PrefixCount), append(labelValues, rib_name)...)
	}

	for _, rib := range p.RIBs {
		l := append(labelValues, rib.Name)
		ch <- prometheus.MustNewConstMetric(receivedPrefixesDesc, prometheus.GaugeValue, float64(rib.ReceivedPrefixes), l...)
		ch <- prometheus.MustNewConstMetric(acceptedPrefixesDesc, prometheus.GaugeValue, float64(rib.AcceptedPrefixes), l...)
		ch <- prometheus.MustNewConstMetric(rejectedPrefixesDesc, prometheus.GaugeValue, float64(rib.RejectedPrefixes), l...)
		ch <- prometheus.MustNewConstMetric(activePrefixesDesc, prometheus.GaugeValue, float64(rib.ActivePrefixes), l...)
		ch <- prometheus.MustNewConstMetric(advertisedPrefixesDesc, prometheus.GaugeValue, float64(rib.AdvertisedPrefixes), l...)

		if rib.Name == rib_name {
			if p.OptionInformation.PrefixLimit.PrefixCount > 0 {
				prefixesLimitPercent := float64(rib.ReceivedPrefixes) / float64(p.OptionInformation.PrefixLimit.PrefixCount)
				ch <- prometheus.MustNewConstMetric(prefixesLimitPercentageDesc, prometheus.GaugeValue, math.Round(prefixesLimitPercent*100)/100, l...)
			}
		}
	}
}
