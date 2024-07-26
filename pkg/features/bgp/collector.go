// SPDX-License-Identifier: MIT

package bgp

import (
	"fmt"
	"math"
	"regexp"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/czerwonk/junos_exporter/pkg/dynamiclabels"
	"github.com/prometheus/client_golang/prometheus"

	"strings"
)

const prefix string = "junos_bgp_session_"

type description struct {
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
}

func newDescriptions(dynLabels dynamiclabels.Labels) *description {
	d := &description{}

	l := []string{"target", "asn", "ip", "description", "group"}
	l = append(l, dynLabels.Keys()...)
	d.upDesc = prometheus.NewDesc(prefix+"up", "Session is up (1 = Established)", l, nil)
	d.stateDesc = prometheus.NewDesc(prefix+"state", "State of the bgp Session (1 = Active, 2 = Connect, 3 = Established, 4 = Idle, 5 = OpenConfirm, 6 = OpenSent, 7 = route reflector client, 0 = Other)", l, nil)
	d.inputMessagesDesc = prometheus.NewDesc(prefix+"messages_input_count", "Number of received messages", l, nil)
	d.outputMessagesDesc = prometheus.NewDesc(prefix+"messages_output_count", "Number of transmitted messages", l, nil)
	d.flapsDesc = prometheus.NewDesc(prefix+"flap_count", "Number of session flaps", l, nil)
	d.medDesc = prometheus.NewDesc(prefix+"metric_out", "MED configured for the session", l, nil)
	d.preferenceDesc = prometheus.NewDesc(prefix+"preference", "Preference configured for the session", l, nil)
	d.holdTimeDesc = prometheus.NewDesc(prefix+"hold_time_seconds", "Hold time configured for the session", l, nil)

	infoLabels := append(l, "local_as", "import_policy", "export_policy", "options")
	d.infoDesc = prometheus.NewDesc(prefix+"info", "Information about the session (e.g. configuration)", infoLabels, nil)

	l = append(l, "table")

	d.receivedPrefixesDesc = prometheus.NewDesc(prefix+"prefixes_received_count", "Number of received prefixes", l, nil)
	d.acceptedPrefixesDesc = prometheus.NewDesc(prefix+"prefixes_accepted_count", "Number of accepted prefixes", l, nil)
	d.rejectedPrefixesDesc = prometheus.NewDesc(prefix+"prefixes_rejected_count", "Number of rejected prefixes", l, nil)
	d.activePrefixesDesc = prometheus.NewDesc(prefix+"prefixes_active_count", "Number of active prefixes (best route in RIB)", l, nil)
	d.advertisedPrefixesDesc = prometheus.NewDesc(prefix+"prefixes_advertised_count", "Number of prefixes announced to peer", l, nil)
	d.prefixesLimitPercentageDesc = prometheus.NewDesc(prefix+"prefixes_limit_percentage", "percentage of received prefixes against prefix-limit", l, nil)
	d.prefixesLimitCountDesc = prometheus.NewDesc(prefix+"prefixes_limit_count", "prefix-count variable set in prefix-limit", l, nil)

	return d
}

type bgpCollector struct {
	LogicalSystem string
	descriptionRe *regexp.Regexp
}

type groupMap map[int64]group

// NewCollector creates a new collector
func NewCollector(logicalSystem string, descRe *regexp.Regexp) collector.RPCCollector {
	return &bgpCollector{
		LogicalSystem: logicalSystem,
		descriptionRe: descRe,
	}
}

// Name returns the name of the collector
func (*bgpCollector) Name() string {
	return "BGP"
}

// Describe describes the metrics
func (*bgpCollector) Describe(ch chan<- *prometheus.Desc) {
	d := newDescriptions(nil)
	ch <- d.upDesc
	ch <- d.receivedPrefixesDesc
	ch <- d.acceptedPrefixesDesc
	ch <- d.rejectedPrefixesDesc
	ch <- d.activePrefixesDesc
	ch <- d.advertisedPrefixesDesc
	ch <- d.inputMessagesDesc
	ch <- d.outputMessagesDesc
	ch <- d.flapsDesc
	ch <- d.prefixesLimitCountDesc
	ch <- d.prefixesLimitPercentageDesc
	ch <- d.infoDesc
	ch <- d.medDesc
	ch <- d.preferenceDesc
	ch <- d.holdTimeDesc
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
	lv := append(labelValues, []string{
		p.ASN,
		ip[0],
		p.Description,
		groupForPeer(p, groups)}...)

	up := 0
	if p.State == "Established" {
		up = 1
	}

	dynLabels := dynamiclabels.ParseDescription(p.Description, c.descriptionRe)
	lv = append(lv, dynLabels.Values()...)

	d := newDescriptions(dynLabels)

	ch <- prometheus.MustNewConstMetric(d.upDesc, prometheus.GaugeValue, float64(up), lv...)
	ch <- prometheus.MustNewConstMetric(d.stateDesc, prometheus.GaugeValue, bgpStateToNumber(p.State), lv...)
	ch <- prometheus.MustNewConstMetric(d.inputMessagesDesc, prometheus.GaugeValue, float64(p.InputMessages), lv...)
	ch <- prometheus.MustNewConstMetric(d.outputMessagesDesc, prometheus.GaugeValue, float64(p.OutputMessages), lv...)
	ch <- prometheus.MustNewConstMetric(d.flapsDesc, prometheus.GaugeValue, float64(p.Flaps), lv...)
	ch <- prometheus.MustNewConstMetric(d.preferenceDesc, prometheus.GaugeValue, float64(p.OptionInformation.Preference), lv...)
	ch <- prometheus.MustNewConstMetric(d.medDesc, prometheus.GaugeValue, float64(p.OptionInformation.MetricOut), lv...)
	ch <- prometheus.MustNewConstMetric(d.holdTimeDesc, prometheus.GaugeValue, float64(p.OptionInformation.Holdtime), lv...)

	infoValues := append(lv,
		localASNForPeer(p),
		formatPolicy(p.OptionInformation.ImportPolicy),
		formatPolicy(p.OptionInformation.ExportPolicy),
		p.OptionInformation.Options)
	ch <- prometheus.MustNewConstMetric(d.infoDesc, prometheus.GaugeValue, 1, infoValues...)

	c.collectRIBForPeer(p, ch, lv, d)
}

func (*bgpCollector) collectRIBForPeer(p peer, ch chan<- prometheus.Metric, labelValues []string, d *description) {
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
		ch <- prometheus.MustNewConstMetric(d.prefixesLimitCountDesc, prometheus.GaugeValue, float64(p.OptionInformation.PrefixLimit.PrefixCount), append(labelValues, rib_name)...)
	}

	for _, rib := range p.RIBs {
		l := append(labelValues, rib.Name)
		ch <- prometheus.MustNewConstMetric(d.receivedPrefixesDesc, prometheus.GaugeValue, float64(rib.ReceivedPrefixes), l...)
		ch <- prometheus.MustNewConstMetric(d.acceptedPrefixesDesc, prometheus.GaugeValue, float64(rib.AcceptedPrefixes), l...)
		ch <- prometheus.MustNewConstMetric(d.rejectedPrefixesDesc, prometheus.GaugeValue, float64(rib.RejectedPrefixes), l...)
		ch <- prometheus.MustNewConstMetric(d.activePrefixesDesc, prometheus.GaugeValue, float64(rib.ActivePrefixes), l...)
		ch <- prometheus.MustNewConstMetric(d.advertisedPrefixesDesc, prometheus.GaugeValue, float64(rib.AdvertisedPrefixes), l...)

		if rib.Name == rib_name {
			if p.OptionInformation.PrefixLimit.PrefixCount > 0 {
				prefixesLimitPercent := float64(rib.ReceivedPrefixes) / float64(p.OptionInformation.PrefixLimit.PrefixCount)
				ch <- prometheus.MustNewConstMetric(d.prefixesLimitPercentageDesc, prometheus.GaugeValue, math.Round(prefixesLimitPercent*100)/100, l...)
			}
		}
	}
}
