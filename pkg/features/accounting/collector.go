// SPDX-License-Identifier: MIT

package accounting

import (
	"fmt"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_accounting_inline_"

var (
	inlineActiveFlowsDesc     *prometheus.Desc
	inlineIpv4ActiveFlowsDesc *prometheus.Desc
	inlineIpv6ActiveFlowsDesc *prometheus.Desc

	inlineFlowsDesc          *prometheus.Desc
	inlineIpv4TotalFlowsDesc *prometheus.Desc
	inlineIpv6TotalFlowsDesc *prometheus.Desc

	inlineFlowCreationFailuresDesc     *prometheus.Desc
	inlineIpv4FlowCreationFailuresDesc *prometheus.Desc
	inlineIpv6FlowCreationFailuresDesc *prometheus.Desc
)

func init() {
	l := []string{"target", "fpc"}
	inlineActiveFlowsDesc = prometheus.NewDesc(prefix+"active_flow_count", "Number of active flows", l, nil)
	inlineIpv4ActiveFlowsDesc = prometheus.NewDesc(prefix+"ipv4_active_flow_count", "Number of active ipv4 flows", l, nil)
	inlineIpv6ActiveFlowsDesc = prometheus.NewDesc(prefix+"ipv6_active_flow_count", "Number of active ipv6 flows", l, nil)

	inlineFlowsDesc = prometheus.NewDesc(prefix+"flow_count", "Number of flows", l, nil)
	inlineIpv4TotalFlowsDesc = prometheus.NewDesc(prefix+"ipv4_flow_count", "Number of ipv4 flows", l, nil)
	inlineIpv6TotalFlowsDesc = prometheus.NewDesc(prefix+"ipv6_flow_count", "Number of ipv6 flows", l, nil)

	inlineFlowCreationFailuresDesc = prometheus.NewDesc(prefix+"creation_failure_count", "Number of flow creation failures", l, nil)
	inlineIpv4FlowCreationFailuresDesc = prometheus.NewDesc(prefix+"ipv4_creation_failure_count", "Number of ipv4 flow creation failures", l, nil)
	inlineIpv6FlowCreationFailuresDesc = prometheus.NewDesc(prefix+"ipv6_creation_failure_count", "Number of ipv6 flow creation failures", l, nil)
}

type accountingCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &accountingCollector{}
}

// Name returns the name of the collector
func (*accountingCollector) Name() string {
	return "Accounting"
}

// Describe describes the metrics
func (*accountingCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- inlineActiveFlowsDesc
	ch <- inlineIpv4ActiveFlowsDesc
	ch <- inlineIpv6ActiveFlowsDesc

	ch <- inlineFlowsDesc
	ch <- inlineIpv4TotalFlowsDesc
	ch <- inlineIpv6TotalFlowsDesc

	ch <- inlineFlowCreationFailuresDesc
	ch <- inlineIpv4FlowCreationFailuresDesc
	ch <- inlineIpv6FlowCreationFailuresDesc
}

// Collect collects metrics from JunOS
func (c *accountingCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	flow, err := c.accountingFlows(client)
	if err != nil {
		return err
	}

	failure, err := c.accountingFailures(client)
	if err != nil {
		return err
	}

	l := append(labelValues, []string{flow.FpcSlot}...)

	ch <- prometheus.MustNewConstMetric(inlineActiveFlowsDesc, prometheus.GaugeValue, float64(flow.InlineActiveFlows), l...)
	ch <- prometheus.MustNewConstMetric(inlineIpv4ActiveFlowsDesc, prometheus.GaugeValue, float64(flow.InlineIPv4ActiveFlows), l...)
	ch <- prometheus.MustNewConstMetric(inlineIpv6ActiveFlowsDesc, prometheus.GaugeValue, float64(flow.InlineIPv6ActiveFlows), l...)
	ch <- prometheus.MustNewConstMetric(inlineFlowsDesc, prometheus.GaugeValue, float64(flow.InlineFlows), l...)
	ch <- prometheus.MustNewConstMetric(inlineIpv4TotalFlowsDesc, prometheus.GaugeValue, float64(flow.InlineIPv4TotalFlows), l...)
	ch <- prometheus.MustNewConstMetric(inlineIpv6TotalFlowsDesc, prometheus.GaugeValue, float64(flow.InlineIPv6TotalFlows), l...)
	ch <- prometheus.MustNewConstMetric(inlineFlowCreationFailuresDesc, prometheus.GaugeValue, float64(failure.InlineFlowCreationFailures), l...)
	ch <- prometheus.MustNewConstMetric(inlineIpv4FlowCreationFailuresDesc, prometheus.GaugeValue, float64(failure.InlineIPv4FlowCreationFailures), l...)
	ch <- prometheus.MustNewConstMetric(inlineIpv6FlowCreationFailuresDesc, prometheus.GaugeValue, float64(failure.InlineIPv6FlowCreationFailures), l...)

	return nil
}

func (c *accountingCollector) accountingFlows(client collector.Client) (*accountingFlow, error) {
	var x = result{}
	err := client.RunCommandAndParse("show services accounting flow inline-jflow", &x)
	if err != nil {
		return nil, err
	}

	if x.Error.Message != "" {
		return nil, fmt.Errorf("accounting command not supported")
	}

	return &accountingFlow{
		FpcSlot:               x.Information.InlineFlow.FpcSlot,
		InlineActiveFlows:     float64(x.Information.InlineFlow.InlineActiveFlows),
		InlineIPv4ActiveFlows: float64(x.Information.InlineFlow.InlineIPv4ActiveFlows),
		InlineIPv6ActiveFlows: float64(x.Information.InlineFlow.InlineIPv6ActiveFlows),

		InlineFlows:          float64(x.Information.InlineFlow.InlineFlows),
		InlineIPv4TotalFlows: float64(x.Information.InlineFlow.InlineIPv4TotalFlows),
		InlineIPv6TotalFlows: float64(x.Information.InlineFlow.InlineIPv6TotalFlows),
	}, nil
}

func (c *accountingCollector) accountingFailures(client collector.Client) (*accountingError, error) {
	var x = accountingFlowError{}
	err := client.RunCommandAndParse("show services accounting errors inline-jflow fpc-slot 0", &x)
	if err != nil {
		return nil, err
	}

	return &accountingError{
		FpcSlot: x.Information.InlineFlow.FpcSlot,

		InlineFlowCreationFailures:     float64(x.Information.InlineFlow.InlineFlowCreationFailures),
		InlineIPv4FlowCreationFailures: float64(x.Information.InlineFlow.InlineIPv4FlowCreationFailures),
		InlineIPv6FlowCreationFailures: float64(x.Information.InlineFlow.InlineIPv6FlowCreationFailures),
	}, nil
}
