// SPDX-License-Identifier: MIT

package twamp

import (
	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_twamp_probe_results_"

var (
	currLossPercentDesc      *prometheus.Desc
	currRTTDesc              *prometheus.Desc
	currRTTJitterDesc        *prometheus.Desc
	currRTTIngressJitterDesc *prometheus.Desc
	currRTTEgressJitterDesc  *prometheus.Desc
	totalSentDesc            *prometheus.Desc
	totalReceivedDesc        *prometheus.Desc
)

func init() {
	l := []string{"target", "owner", "test", "target_address", "source_address", "type"}
	totalSentDesc = prometheus.NewDesc(prefix+"sent_total", "Number of probes sent within the current test", l, nil)
	totalReceivedDesc = prometheus.NewDesc(prefix+"received_total", "Number of probe responses received within the current test", l, nil)
	currLossPercentDesc = prometheus.NewDesc(prefix+"loss_percent_current", "Percentage of probes lost during the most recently completed test", l, nil)
	currRTTDesc = prometheus.NewDesc(prefix+"rtt_current", "RTT for the most recently completed test, in microseconds", l, nil)
	currRTTJitterDesc = prometheus.NewDesc(prefix+"rtt_jitter_current", "Jitter for the most recently completed test, in microseconds", l, nil)
	currRTTIngressJitterDesc = prometheus.NewDesc(prefix+"ingress_jitter_current", "Ingress Jitter for the most recently completed test, in microseconds", l, nil)
	currRTTEgressJitterDesc = prometheus.NewDesc(prefix+"egress_jitter_current", "Egress Jitter for the most recently completed test, in microseconds", l, nil)
}

type twampCollector struct{}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &twampCollector{}
}

// Name returns the name of the collector
func (*twampCollector) Name() string {
	return "twamp"
}

// Describe describes the metrics
func (*twampCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- totalSentDesc
	ch <- totalReceivedDesc
	ch <- currLossPercentDesc
	ch <- currRTTDesc
	ch <- currRTTJitterDesc
	ch <- currRTTIngressJitterDesc
	ch <- currRTTEgressJitterDesc
}

// Collect collects metrics from JunOS
func (c *twampCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	err := c.collect(client, ch, labelValues)
	if err != nil {
		return err
	}

	return nil
}

func (c *twampCollector) collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = result{}

	err := client.RunCommandAndParse("show services monitoring twamp client probe-results", &x)
	if err != nil {
		return err
	}

	for _, probe := range x.Results.Probes {
		c.collectForProbe(probe, ch, labelValues)
	}

	return nil
}

func (c *twampCollector) collectForProbe(p probe, ch chan<- prometheus.Metric, labelValues []string) {
	l := append(labelValues, []string{p.Owner, p.Test, p.TargetAddress, p.SourceAddress, p.Type}...)
	ch <- prometheus.MustNewConstMetric(currRTTDesc, prometheus.GaugeValue, float64(p.GenericSampleResults.RTT), l...)
	ch <- prometheus.MustNewConstMetric(currRTTJitterDesc, prometheus.GaugeValue, float64(p.GenericSampleResults.RTTJitter), l...)
	ch <- prometheus.MustNewConstMetric(currRTTIngressJitterDesc, prometheus.GaugeValue, float64(p.GenericSampleResults.IngressJitter), l...)
	ch <- prometheus.MustNewConstMetric(currRTTEgressJitterDesc, prometheus.GaugeValue, float64(p.GenericSampleResults.EgressJitter), l...)

	for _, aggResult := range p.GenericAggregateResults {
		if aggResult.AggregateType == "last test" {
			ch <- prometheus.MustNewConstMetric(totalSentDesc, prometheus.GaugeValue, float64(aggResult.NumSamplesTx), l...)
			ch <- prometheus.MustNewConstMetric(totalReceivedDesc, prometheus.GaugeValue, float64(aggResult.NumSamplesRx), l...)
			ch <- prometheus.MustNewConstMetric(currLossPercentDesc, prometheus.GaugeValue, float64(aggResult.LossPercentage), l...)

			// Loop through measurements within this aggregate result
			//			if len(aggResult.GenericAggregateMeasurement) > 0 {
			//				// Example: Print details of the first measurement type (e.g., Round trip time)
			//				// You can add more logic here to find specific measurement types if needed
			//				for k, measurement := range aggResult.GenericAggregateMeasurement {
			//					fmt.Printf("      Measurement #%d (%s):\n", k+1, measurement.MeasurementType)
			//					fmt.Printf("        Avg: %d, Min: %d, Max: %d, StdDev: %d, Samples: %d\n",
			//						measurement.MeasurementAvg,
			//						measurement.MeasurementMin,
			//						measurement.MeasurementMax,
			//						measurement.MeasurementStddev,
			//						measurement.MeasurementSamples)
			//				}
			//			}
		}
	}
}
