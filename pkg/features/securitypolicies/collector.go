// SPDX-License-Identifier: MIT

package securitypolicies

import (
	"encoding/xml"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_security_policies_"

var (
	hitCountDesc         *prometheus.Desc
	sessionCreationsDesc *prometheus.Desc
	sessionDeletionsDesc *prometheus.Desc
	inputBytesDesc       *prometheus.Desc
	outputBytesDesc      *prometheus.Desc
	inputPacketsDesc     *prometheus.Desc
	outputPacketsDesc    *prometheus.Desc
)

func init() {
	la := []string{"target", "from_zone", "to_zone", "policy_name"}
	lb := []string{"target", "from_zone", "to_zone", "policy_name", "direction"}

	hitCountDesc = prometheus.NewDesc(prefix+"hit_count", "Policy hit count", la, nil)

	sessionCreationsDesc = prometheus.NewDesc(prefix+"session_creations", "Policy session creations", la, nil)
	sessionDeletionsDesc = prometheus.NewDesc(prefix+"session_deletions", "Policy session deletions", la, nil)

	inputBytesDesc = prometheus.NewDesc(prefix+"input_bytes", "Policy input bytes", lb, nil)
	outputBytesDesc = prometheus.NewDesc(prefix+"output_bytes", "Policy output bytes", lb, nil)
	inputPacketsDesc = prometheus.NewDesc(prefix+"input_packets", "Policy input packets", lb, nil)
	outputPacketsDesc = prometheus.NewDesc(prefix+"output_packets", "Policy output packets", lb, nil)
}

type securityPolicyCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &securityPolicyCollector{}
}

// Name returns the name of the collector
func (*securityPolicyCollector) Name() string {
	return "Security Policies"
}

// Describe describes the metrics
func (*securityPolicyCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- hitCountDesc
	ch <- sessionCreationsDesc
	ch <- sessionDeletionsDesc
	ch <- inputBytesDesc
	ch <- outputBytesDesc
	ch <- inputBytesDesc
	ch <- outputBytesDesc
}

// Collect collects metrics from JunOS
func (c *securityPolicyCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	err := c.CollectStats(client, ch, labelValues)
	if err != nil {
		return err
	}

	err = c.CollectHits(client, ch, labelValues)
	if err != nil {
		return err
	}

	return nil
}

func (c *securityPolicyCollector) CollectStats(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = statsMultiEngineResult{}
	err := client.RunCommandAndParseWithParser("show security policies detail", func(b []byte) error {
		return parseStatsXML(b, &x)
	})
	if err != nil {
		return err
	}

	for _, ctx := range x.Results.RoutingEngines[0].Policies.Contexts {
		for _, pol := range ctx.Policies {
			if pol.PolicyInformation.StatisticsInformation == nil {
				continue
			}

			var ln, li, lr []string
			if ctx.ContextInformation.GlobalContext == nil {
				ln = append(labelValues, ctx.ContextInformation.FromZone, ctx.ContextInformation.ToZone)
			} else {
				ln = append(labelValues, "junos-global", "junos-global")
			}
			ln = append(ln, pol.PolicyInformation.Name)

			li = append(li, ln...)
			li = append(li, "initial")
			lr = append(lr, ln...)
			lr = append(lr, "reply")

			ch <- prometheus.MustNewConstMetric(sessionCreationsDesc, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.SessionCreations, ln...)
			ch <- prometheus.MustNewConstMetric(sessionDeletionsDesc, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.SessionDeletions, ln...)
			ch <- prometheus.MustNewConstMetric(inputBytesDesc, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.InputBytesInit, li...)
			ch <- prometheus.MustNewConstMetric(inputBytesDesc, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.InputBytesReply, lr...)
			ch <- prometheus.MustNewConstMetric(outputBytesDesc, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.OutputBytesInit, li...)
			ch <- prometheus.MustNewConstMetric(outputBytesDesc, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.OutputBytesReply, lr...)
			ch <- prometheus.MustNewConstMetric(inputPacketsDesc, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.InputPacketsInit, li...)
			ch <- prometheus.MustNewConstMetric(inputPacketsDesc, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.InputPacketsReply, lr...)
			ch <- prometheus.MustNewConstMetric(outputPacketsDesc, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.OutputPacketsInit, li...)
			ch <- prometheus.MustNewConstMetric(outputPacketsDesc, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.OutputPacketsReply, lr...)
		}
	}

	return nil
}

func (c *securityPolicyCollector) CollectHits(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = hitsMultiEngineResult{}
	err := client.RunCommandAndParseWithParser("show security policies hit-count", func(b []byte) error {
		return parseHitsXML(b, &x)
	})
	if err != nil {
		return err
	}

	for _, pol := range x.MultiRoutingEngineResults.RoutingEngine[0].HitCount.Policies {
		ls := append(labelValues, pol.FromZone, pol.ToZone, pol.PolicyName)
		ch <- prometheus.MustNewConstMetric(hitCountDesc, prometheus.CounterValue, pol.Count, ls...)
	}
	return nil
}

func parseStatsXML(b []byte, res *statsMultiEngineResult) error {
	if strings.Contains(string(b), "multi-routing-engine-results") {
		return xml.Unmarshal(b, res)
	}

	fi := statsSingleEngineResult{}

	err := xml.Unmarshal(b, &fi)
	if err != nil {
		return err
	}

	res.Results.RoutingEngines = []statsRoutingEngine{
		{
			Name:     "N/A",
			Policies: fi.Policies,
		},
	}

	return nil
}

func parseHitsXML(b []byte, res *hitsMultiEngineResult) error {
	if strings.Contains(string(b), "multi-routing-engine-results") {
		return xml.Unmarshal(b, res)
	}

	fi := hitsSingleEngineResult{}

	err := xml.Unmarshal(b, &fi)
	if err != nil {
		return err
	}

	res.MultiRoutingEngineResults.RoutingEngine = []hitsRoutingEngine{
		{
			Name:     "N/A",
			HitCount: fi.HitCount,
		},
	}

	return nil
}
