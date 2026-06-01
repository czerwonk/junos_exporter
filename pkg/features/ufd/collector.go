// SPDX-License-Identifier: MIT

package ufd

import (
	"encoding/xml"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_ufd_"

var (
	groupFailureActive *prometheus.Desc
	groupDebounce      *prometheus.Desc
	groupUplinkCount   *prometheus.Desc
	groupDownlinkCount *prometheus.Desc
	uplinkActive       *prometheus.Desc
	downlinkActive     *prometheus.Desc
)

func init() {
	groupLabels := []string{"target", "group"}
	groupLabelsWithAction := []string{"target", "group", "failure_action"}
	ifaceLabels := []string{"target", "group", "interface"}

	groupFailureActive = prometheus.NewDesc(
		prefix+"group_failure_active",
		"1 if the UFD group's failure-action is currently engaged (i.e. not 'Inactive'), 0 otherwise",
		groupLabelsWithAction, nil,
	)
	groupDebounce = prometheus.NewDesc(
		prefix+"group_debounce_seconds",
		"UFD group debounce interval in seconds",
		groupLabels, nil,
	)
	groupUplinkCount = prometheus.NewDesc(
		prefix+"group_uplink_count",
		"Number of uplinks in the UFD group's link-to-monitor list",
		groupLabels, nil,
	)
	groupDownlinkCount = prometheus.NewDesc(
		prefix+"group_downlink_count",
		"Number of downlinks in the UFD group's link-to-disable list",
		groupLabels, nil,
	)
	uplinkActive = prometheus.NewDesc(
		prefix+"uplink_active",
		"1 if the monitored uplink is currently marked active by UFD (asterisk in show output), 0 otherwise",
		ifaceLabels, nil,
	)
	downlinkActive = prometheus.NewDesc(
		prefix+"downlink_active",
		"1 if the downlink is currently marked active (asterisk in show output, i.e. not disabled by UFD), 0 otherwise",
		ifaceLabels, nil,
	)
}

type ufdCollector struct{}

// Name returns the name of the collector
func (*ufdCollector) Name() string { return "ufd" }

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector { return &ufdCollector{} }

// Describe describes the metrics
func (*ufdCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- groupFailureActive
	ch <- groupDebounce
	ch <- groupUplinkCount
	ch <- groupDownlinkCount
	ch <- uplinkActive
	ch <- downlinkActive
}

// Collect collects metrics from JunOS
func (c *ufdCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var groups []ufdGroup
	err := client.RunCommandAndParseWithParser("show uplink-failure-detection", func(b []byte) error {
		var perr error
		groups, perr = parseGroups(b)
		return perr
	})
	if err != nil {
		return err
	}

	for _, g := range groups {
		groupLabels := append(labelValues, g.Name)

		active := 0.0
		if g.FailureAction != "" && !strings.EqualFold(strings.TrimSpace(g.FailureAction), "Inactive") {
			active = 1.0
		}
		ch <- prometheus.MustNewConstMetric(
			groupFailureActive, prometheus.GaugeValue, active,
			append(groupLabels, strings.TrimSpace(g.FailureAction))...,
		)

		ch <- prometheus.MustNewConstMetric(groupDebounce, prometheus.GaugeValue, g.DebounceInterval, groupLabels...)
		ch <- prometheus.MustNewConstMetric(groupUplinkCount, prometheus.GaugeValue, float64(len(g.Uplinks)), groupLabels...)
		ch <- prometheus.MustNewConstMetric(groupDownlinkCount, prometheus.GaugeValue, float64(len(g.Downlinks)), groupLabels...)

		for _, raw := range g.Uplinks {
			name, marked := splitMark(raw)
			ifaceLabels := append(append([]string{}, groupLabels...), name)
			ch <- prometheus.MustNewConstMetric(uplinkActive, prometheus.GaugeValue, marked, ifaceLabels...)
		}

		for _, raw := range g.Downlinks {
			name, marked := splitMark(raw)
			ifaceLabels := append(append([]string{}, groupLabels...), name)
			ch <- prometheus.MustNewConstMetric(downlinkActive, prometheus.GaugeValue, marked, ifaceLabels...)
		}
	}

	return nil
}

// parseGroups handles both response shapes documented in the QFX YANG: the
// direct <uplink-failure-detection-information> body and the multi-RE / VC
// wrapper. Same dispatch pattern as the alarm collector. When multiple REs
// report groups, they are merged - the re-name is not exposed as a label
// today (matches the alarm collector's behaviour).
func parseGroups(b []byte) ([]ufdGroup, error) {
	if strings.Contains(string(b), "<multi-routing-engine-results") {
		var m multiEngineResult
		if err := xml.Unmarshal(b, &m); err != nil {
			return nil, err
		}
		var groups []ufdGroup
		for _, re := range m.Engines.Items {
			groups = append(groups, re.Groups...)
		}
		return groups, nil
	}

	var s singleEngineResult
	if err := xml.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return s.Groups, nil
}

// splitMark separates the trailing "*" marker from the interface name and
// returns the clean name plus 1.0 if the marker was present, 0.0 otherwise.
func splitMark(s string) (string, float64) {
	s = strings.TrimSpace(s)
	if before, ok := strings.CutSuffix(s, "*"); ok {
		return before, 1.0
	}
	return s, 0.0
}
