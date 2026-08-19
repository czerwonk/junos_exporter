// SPDX-License-Identifier: MIT

package securitynat

import (
	"errors"
	"strconv"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_security_nat_"

const (
	poolUsageRPC = "show security nat resource-usage source-pool all"
	ruleRPC      = "show security nat source rule all"
)

var (
	poolResourceUsedDesc              *prometheus.Desc
	poolResourceAvailableDesc         *prometheus.Desc
	poolResourceTotalDesc             *prometheus.Desc
	poolUsagePercentDesc              *prometheus.Desc
	poolPeakUsagePercentDesc          *prometheus.Desc
	poolPeakUsageTimestampSecondsDesc *prometheus.Desc
	poolAddressCountDesc              *prometheus.Desc
	poolPortOverloadFactorDesc        *prometheus.Desc

	rulesCountDesc                      *prometheus.Desc
	ruleReferencedAddressCountDesc      *prometheus.Desc
	ruleTranslationHitsTotalDesc        *prometheus.Desc
	ruleTranslationSuccessHitsTotalDesc *prometheus.Desc
	ruleConcurrentHitsDesc              *prometheus.Desc
	ruleInfoDesc                        *prometheus.Desc
)

func init() {
	// The used/available/total triple is counted in ports for port-translating
	// pools and in addresses otherwise, which is what the style label conveys.
	lp := []string{"target", "pool_name", "style"}
	poolResourceUsedDesc = prometheus.NewDesc(prefix+"pool_resource_used", "Pool resources (ports for PAT pools, addresses otherwise) currently in use", lp, nil)
	poolResourceAvailableDesc = prometheus.NewDesc(prefix+"pool_resource_available", "Pool resources currently available", lp, nil)
	poolResourceTotalDesc = prometheus.NewDesc(prefix+"pool_resource_total", "Pool resources available in total", lp, nil)
	poolUsagePercentDesc = prometheus.NewDesc(prefix+"pool_usage_percent", "Current pool usage in percent as reported by Junos", lp, nil)
	poolPeakUsagePercentDesc = prometheus.NewDesc(prefix+"pool_peak_usage_percent", "Highest pool usage in percent observed by the device so far", lp, nil)
	poolPeakUsageTimestampSecondsDesc = prometheus.NewDesc(prefix+"pool_peak_usage_timestamp_seconds", "Time the peak pool usage was observed (not reported while no peak was recorded)", lp, nil)
	poolAddressCountDesc = prometheus.NewDesc(prefix+"pool_address_count", "Number of addresses in the pool", lp, nil)
	poolPortOverloadFactorDesc = prometheus.NewDesc(prefix+"pool_port_overload_factor", "Port overloading factor configured for the pool", lp, nil)

	l := []string{"target"}
	rulesCountDesc = prometheus.NewDesc(prefix+"rules_count", "Number of source NAT rules configured", l, nil)

	lv := []string{"target", "ip_version"}
	ruleReferencedAddressCountDesc = prometheus.NewDesc(prefix+"rule_referenced_address_count", "Number of addresses referenced by source NAT rules", lv, nil)

	// Only labels identifying the rule and the pool it translates to, so that
	// rate() over the counters survives changes to the zones a rule matches.
	lr := []string{"target", "rule_name", "rule_set_name", "rule_id", "pool_name"}
	ruleTranslationHitsTotalDesc = prometheus.NewDesc(prefix+"rule_translation_hits_total", "Translations attempted by the rule", lr, nil)
	ruleTranslationSuccessHitsTotalDesc = prometheus.NewDesc(prefix+"rule_translation_success_hits_total", "Translations performed successfully by the rule", lr, nil)
	ruleConcurrentHitsDesc = prometheus.NewDesc(prefix+"rule_concurrent_hits", "Sessions currently using the rule", lr, nil)

	// The zones a rule matches churn on configuration changes, so they are
	// exposed as an info metric instead of on the counters above.
	li := []string{"target", "rule_name", "rule_set_name", "rule_id", "from_zone", "to_zone"}
	ruleInfoDesc = prometheus.NewDesc(prefix+"rule_info", "Zones matched by the source NAT rule, one series per zone pair (always 1)", li, nil)
}

type securityNATCollector struct{}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &securityNATCollector{}
}

// Name returns the name of the collector
func (*securityNATCollector) Name() string {
	return "Security NAT"
}

// Describe describes the metrics
func (*securityNATCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- poolResourceUsedDesc
	ch <- poolResourceAvailableDesc
	ch <- poolResourceTotalDesc
	ch <- poolUsagePercentDesc
	ch <- poolPeakUsagePercentDesc
	ch <- poolPeakUsageTimestampSecondsDesc
	ch <- poolAddressCountDesc
	ch <- poolPortOverloadFactorDesc
	ch <- rulesCountDesc
	ch <- ruleReferencedAddressCountDesc
	ch <- ruleTranslationHitsTotalDesc
	ch <- ruleTranslationSuccessHitsTotalDesc
	ch <- ruleConcurrentHitsDesc
	ch <- ruleInfoDesc
}

// Collect collects metrics from JunOS
func (c *securityNATCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var errs []error

	if err := c.collectPoolUsage(client, ch, labelValues); err != nil {
		errs = append(errs, err)
	}

	if err := c.collectRules(client, ch, labelValues); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func (c *securityNATCollector) collectPoolUsage(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var res poolUsageResult
	if err := client.RunCommandAndParse(poolUsageRPC, &res); err != nil {
		return err
	}

	for _, e := range res.Info.Entries {
		l := labelsFor(labelValues, trim(e.PoolName), trim(e.Style))

		ch <- prometheus.MustNewConstMetric(poolResourceUsedDesc, prometheus.GaugeValue, e.TotalUsed, l...)
		ch <- prometheus.MustNewConstMetric(poolResourceAvailableDesc, prometheus.GaugeValue, e.TotalAvail, l...)
		ch <- prometheus.MustNewConstMetric(poolResourceTotalDesc, prometheus.GaugeValue, e.TotalTotal, l...)
		ch <- prometheus.MustNewConstMetric(poolAddressCountDesc, prometheus.GaugeValue, e.TotalAddress, l...)
		ch <- prometheus.MustNewConstMetric(poolPortOverloadFactorDesc, prometheus.GaugeValue, e.PortOlFactor, l...)

		if v, ok := parsePercent(e.TotalUsage); ok {
			ch <- prometheus.MustNewConstMetric(poolUsagePercentDesc, prometheus.GaugeValue, v, l...)
		}

		if v, ok := parsePercent(e.PeakUsage); ok {
			ch <- prometheus.MustNewConstMetric(poolPeakUsagePercentDesc, prometheus.GaugeValue, v, l...)
		}

		// A pool that never saw traffic carries no peak timestamp; reporting it
		// as 0 would claim a peak at the start of the unix epoch.
		if e.PeakDateTime.Seconds > 0 {
			ch <- prometheus.MustNewConstMetric(poolPeakUsageTimestampSecondsDesc, prometheus.GaugeValue, float64(e.PeakDateTime.Seconds), l...)
		}
	}

	return nil
}

func (c *securityNATCollector) collectRules(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var res ruleResult
	if err := client.RunCommandAndParse(ruleRPC, &res); err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(rulesCountDesc, prometheus.GaugeValue, res.Info.TotalRules.Total, labelValues...)
	ch <- prometheus.MustNewConstMetric(ruleReferencedAddressCountDesc, prometheus.GaugeValue, res.Info.TotalRefAddr.V4, labelsFor(labelValues, "ipv4")...)
	ch <- prometheus.MustNewConstMetric(ruleReferencedAddressCountDesc, prometheus.GaugeValue, res.Info.TotalRefAddr.V6, labelsFor(labelValues, "ipv6")...)

	for _, r := range res.Info.Rules {
		name, setName, id := trim(r.Name), trim(r.SetName), trim(r.ID)

		for _, from := range orEmpty(r.FromContextName) {
			for _, to := range orEmpty(r.ToContextName) {
				li := labelsFor(labelValues, name, setName, id, trim(from), trim(to))
				ch <- prometheus.MustNewConstMetric(ruleInfoDesc, prometheus.GaugeValue, 1, li...)
			}
		}

		if r.Hits == nil {
			continue
		}

		lr := labelsFor(labelValues, name, setName, id, trim(r.Action.Pool))
		ch <- prometheus.MustNewConstMetric(ruleTranslationHitsTotalDesc, prometheus.CounterValue, r.Hits.TranslationHits, lr...)
		ch <- prometheus.MustNewConstMetric(ruleTranslationSuccessHitsTotalDesc, prometheus.CounterValue, r.Hits.SuccessHits, lr...)
		ch <- prometheus.MustNewConstMetric(ruleConcurrentHitsDesc, prometheus.GaugeValue, r.Hits.ConcurrentHits, lr...)
	}

	return nil
}

// labelsFor appends values to the label values common to all metrics of this
// collector without ever writing into the caller's backing array.
func labelsFor(labelValues []string, values ...string) []string {
	l := make([]string, 0, len(labelValues)+len(values))
	l = append(l, labelValues...)

	return append(l, values...)
}

// orEmpty makes a single iteration over an absent repeated element possible, so
// a rule without an explicit context still produces an info series.
func orEmpty(values []string) []string {
	if len(values) == 0 {
		return []string{""}
	}

	return values
}

// trim removes the padding Junos adds to some string elements.
func trim(s string) string {
	return strings.TrimSpace(s)
}

// parsePercent parses Junos' "<n>%" percentage format, e.g. "31%" -> 31.
func parsePercent(s string) (float64, bool) {
	v, err := strconv.ParseFloat(strings.TrimSuffix(trim(s), "%"), 64)
	if err != nil {
		return 0, false
	}

	return v, true
}
