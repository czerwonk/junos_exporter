// SPDX-License-Identifier: MIT

package evpnipprefix

import (
	"encoding/xml"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_evpn_ip_prefix_"

var (
	localCount         *prometheus.Desc
	remoteCount        *prometheus.Desc
	advertisementCount *prometheus.Desc
)

func init() {
	countLabels := []string{"target", "context", "afi"}
	advLabels := []string{"target", "context", "afi", "status"}

	localCount = prometheus.NewDesc(prefix+"local_count",
		"Locally originated EVPN Type-5 IP-prefix routes per L3 context and AFI",
		countLabels, nil)
	remoteCount = prometheus.NewDesc(prefix+"remote_count",
		"Remotely advertised EVPN Type-5 IP-prefix routes per L3 context and AFI",
		countLabels, nil)
	advertisementCount = prometheus.NewDesc(prefix+"advertisement_count",
		"Per-PE EVPN Type-5 advertisements grouped by status (accepted, rejected, ...) per L3 context and AFI",
		advLabels, nil)
}

type evpnIPPrefixCollector struct{}

// Name returns the name of the collector.
func (*evpnIPPrefixCollector) Name() string { return "evpn_ip_prefix" }

// NewCollector creates a new collector.
func NewCollector() collector.RPCCollector { return new(evpnIPPrefixCollector) }

// Describe describes the metrics.
func (*evpnIPPrefixCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- localCount
	ch <- remoteCount
	ch <- advertisementCount
}

// Collect issues `show evpn ip-prefix-database` and emits per-context per-AFI
// counters. Per-prefix series are intentionally NOT emitted — cardinality on
// a busy fabric (thousands of prefixes × multiple advertising PEs) would be
// untenable for Prometheus. Operators who need per-prefix detail should query
// the CLI directly.
func (c *evpnIPPrefixCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var contexts []pfxL3Context
	err := client.RunCommandAndParseWithParser("show evpn ip-prefix-database", func(b []byte) error {
		var perr error
		contexts, perr = parseContexts(b)
		return perr
	})
	if err != nil {
		return err
	}

	for _, ctx := range contexts {
		// Locally originated prefixes — one count per AFI.
		for _, t := range ctx.LocalTables {
			afi := afiFromDescription(t.Description)
			l := append(append([]string{}, labelValues...), ctx.Name, afi)
			ch <- prometheus.MustNewConstMetric(localCount, prometheus.GaugeValue, float64(len(t.Entries)), l...)
		}

		// Remotely received prefixes — count of prefix entries per AFI,
		// plus a count of advertisements per status per AFI.
		for _, t := range ctx.RemoteTables {
			afi := afiFromDescription(t.Description)
			l := append(append([]string{}, labelValues...), ctx.Name, afi)
			ch <- prometheus.MustNewConstMetric(remoteCount, prometheus.GaugeValue, float64(len(t.Entries)), l...)

			byStatus := map[string]int{}
			for _, e := range t.Entries {
				for _, adv := range e.Advertisements {
					byStatus[normalizeStatus(adv.Status)]++
				}
			}
			for status, n := range byStatus {
				sl := append(append([]string{}, labelValues...), ctx.Name, afi, status)
				ch <- prometheus.MustNewConstMetric(advertisementCount, prometheus.GaugeValue, float64(n), sl...)
			}
		}
	}

	return nil
}

// afiFromDescription extracts the AFI from <table-description> text such as
// "IPv4->EVPN" or "EVPN->IPv6", returning "v4" or "v6". The arrow direction
// is irrelevant — the AFI is the same for both ends of the table.
func afiFromDescription(d string) string {
	switch {
	case strings.Contains(d, "IPv4"):
		return "v4"
	case strings.Contains(d, "IPv6"):
		return "v6"
	default:
		return strings.ToLower(strings.TrimSpace(d))
	}
}

// normalizeStatus lowercases the status string so the label values are stable
// regardless of the device's exact capitalisation.
func normalizeStatus(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" {
		return "unknown"
	}
	return s
}

// parseContexts handles both single-RE and multi-RE response shapes. Same
// substring-dispatch pattern as the other recent collectors.
func parseContexts(b []byte) ([]pfxL3Context, error) {
	if strings.Contains(string(b), "<multi-routing-engine-results") {
		var m multiEngineResult
		if err := xml.Unmarshal(b, &m); err != nil {
			return nil, err
		}
		var out []pfxL3Context
		for _, re := range m.Engines.Items {
			out = append(out, re.Contexts...)
		}
		return out, nil
	}

	var s singleEngineResult
	if err := xml.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return s.Contexts, nil
}
