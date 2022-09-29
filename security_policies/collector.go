package security_policies

import (
	"encoding/xml"
	"strings"

	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_security_policies_"

var (
	hitCount         *prometheus.Desc
	sessionCreations *prometheus.Desc
	sessionDeletions *prometheus.Desc
	inputBytes       *prometheus.Desc
	outputBytes      *prometheus.Desc
	inputPackets     *prometheus.Desc
	outputPackets    *prometheus.Desc
)

func init() {
	la := []string{"target", "from_zone", "to_zone", "policy_name"}
	lb := []string{"target", "from_zone", "to_zone", "policy_name", "direction"}

	hitCount = prometheus.NewDesc(prefix+"hit_count", "Policy hit count", la, nil)

	sessionCreations = prometheus.NewDesc(prefix+"session_creations", "Policy session creations", la, nil)
	sessionDeletions = prometheus.NewDesc(prefix+"session_deletions", "Policy session deletions", la, nil)

	inputBytes = prometheus.NewDesc(prefix+"input_bytes", "Policy input bytes", lb, nil)
	outputBytes = prometheus.NewDesc(prefix+"output_bytes", "Policy output bytes", lb, nil)
	inputPackets = prometheus.NewDesc(prefix+"input_packets", "Policy input packets", lb, nil)
	outputPackets = prometheus.NewDesc(prefix+"output_packets", "Policy output packets", lb, nil)
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
	ch <- hitCount
	ch <- sessionCreations
	ch <- sessionDeletions
	ch <- inputBytes
	ch <- outputBytes
	ch <- inputBytes
	ch <- outputBytes
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
	var x = StatsRpcReply{}
	err := client.RunCommandAndParseWithParser("show security policies detail", func(b []byte) error {
		return parseStatsXML(b, &x)
	})
	if err != nil {
		return err
	}

	for _, ctx := range x.MultiRoutingEngineResults.RoutingEngine[0].Policies.Contexts {
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

			ch <- prometheus.MustNewConstMetric(sessionCreations, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.SessionCreations, ln...)
			ch <- prometheus.MustNewConstMetric(sessionDeletions, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.SessionDeletions, ln...)
			ch <- prometheus.MustNewConstMetric(inputBytes, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.InputBytesInit, li...)
			ch <- prometheus.MustNewConstMetric(inputBytes, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.InputBytesReply, lr...)
			ch <- prometheus.MustNewConstMetric(outputBytes, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.OutputBytesInit, li...)
			ch <- prometheus.MustNewConstMetric(outputBytes, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.OutputBytesReply, lr...)
			ch <- prometheus.MustNewConstMetric(inputPackets, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.InputPacketsInit, li...)
			ch <- prometheus.MustNewConstMetric(inputPackets, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.InputPacketsReply, lr...)
			ch <- prometheus.MustNewConstMetric(outputPackets, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.OutputPacketsInit, li...)
			ch <- prometheus.MustNewConstMetric(outputPackets, prometheus.CounterValue, pol.PolicyInformation.StatisticsInformation.OutputPacketsReply, lr...)
		}
	}

	return nil
}

func (c *securityPolicyCollector) CollectHits(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = HitsRpcReply{}
	err := client.RunCommandAndParseWithParser("show security policies hit-count", func(b []byte) error {
		return parseHitsXML(b, &x)
	})
	if err != nil {
		return err
	}

	for _, pol := range x.MultiRoutingEngineResults.RoutingEngine[0].HitCount.Policies {
		ls := append(labelValues, pol.FromZone, pol.ToZone, pol.PolicyName)
		ch <- prometheus.MustNewConstMetric(hitCount, prometheus.CounterValue, pol.Count, ls...)
	}
	return nil
}

func parseStatsXML(b []byte, res *StatsRpcReply) error {
	if strings.Contains(string(b), "multi-routing-engine-results") {
		return xml.Unmarshal(b, res)
	}

	fi := StatsRpcReplyNoRE{}

	err := xml.Unmarshal(b, &fi)
	if err != nil {
		return err
	}

	res.MultiRoutingEngineResults.RoutingEngine = []StatsRoutingEngine{
		{
			Name:     "N/A",
			Policies: fi.Policies,
		},
	}
	return nil
}

func parseHitsXML(b []byte, res *HitsRpcReply) error {
	if strings.Contains(string(b), "multi-routing-engine-results") {
		return xml.Unmarshal(b, res)
	}

	fi := HitsRpcReplyNoRE{}

	err := xml.Unmarshal(b, &fi)
	if err != nil {
		return err
	}

	res.MultiRoutingEngineResults.RoutingEngine = []HitsRoutingEngine{
		{
			Name:     "N/A",
			HitCount: fi.HitCount,
		},
	}
	return nil
}
