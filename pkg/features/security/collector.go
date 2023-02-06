package security

import (
	"encoding/xml"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_security_"

var (
	fpcNumber          *prometheus.Desc
	picNumber          *prometheus.Desc
	cpuUtilization     *prometheus.Desc
	memoryUtilization  *prometheus.Desc
	currentFlowSession *prometheus.Desc
	maxFlowSession     *prometheus.Desc
	currentCpSession   *prometheus.Desc
	maxCpSession       *prometheus.Desc
)

func init() {
	l := []string{"target", "re_name"}

	fpcNumber = prometheus.NewDesc(prefix+"fpc_number", "FPC number", l, nil)
	picNumber = prometheus.NewDesc(prefix+"pic_number", "PIC number", l, nil)
	cpuUtilization = prometheus.NewDesc(prefix+"cpu_utilization", "CPU utilization", l, nil)
	memoryUtilization = prometheus.NewDesc(prefix+"memory_utilization", "Memory utilization", l, nil)
	currentFlowSession = prometheus.NewDesc(prefix+"current_flow_session", "Current flow of session", l, nil)
	maxFlowSession = prometheus.NewDesc(prefix+"maximum_flow_session", "Maximum flow of session", l, nil)
	currentCpSession = prometheus.NewDesc(prefix+"current_cp_session", "Current central point session", l, nil)
	maxCpSession = prometheus.NewDesc(prefix+"max_cp_session", "Maximum central point session", l, nil)
}

type securityCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &securityCollector{}
}

// Name returns the name of the collector
func (*securityCollector) Name() string {
	return "Security"
}

// Describe describes the metrics
func (*securityCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- fpcNumber
	ch <- picNumber
	ch <- cpuUtilization
	ch <- memoryUtilization
	ch <- currentFlowSession
	ch <- maxFlowSession
	ch <- currentCpSession
	ch <- maxCpSession
}

// Collect collects metrics from JunOS
func (c *securityCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = multiEngineResult{}
	err := client.RunCommandAndParseWithParser("show security monitoring", func(b []byte) error {
		return parseXML(b, &x)
	})
	if err != nil {
		return err
	}

	for _, re := range x.Results.RoutingEngines {
		ls := append(labelValues, re.Name)
		for _, ps := range re.PerformanceSummary.PerformanceStatistics {
			ch <- prometheus.MustNewConstMetric(fpcNumber, prometheus.GaugeValue, float64(ps.FPCNumber), ls...)
			ch <- prometheus.MustNewConstMetric(picNumber, prometheus.GaugeValue, float64(ps.PICNumber), ls...)
			ch <- prometheus.MustNewConstMetric(cpuUtilization, prometheus.GaugeValue, float64(ps.CPUUtil), ls...)
			ch <- prometheus.MustNewConstMetric(memoryUtilization, prometheus.GaugeValue, float64(ps.MemoryUtil), ls...)
			ch <- prometheus.MustNewConstMetric(currentFlowSession, prometheus.GaugeValue, float64(ps.CurrentFlow), ls...)
			ch <- prometheus.MustNewConstMetric(maxFlowSession, prometheus.GaugeValue, float64(ps.MaxFlow), ls...)
			ch <- prometheus.MustNewConstMetric(currentCpSession, prometheus.GaugeValue, float64(ps.CurrentCP), ls...)
			ch <- prometheus.MustNewConstMetric(maxCpSession, prometheus.GaugeValue, float64(ps.MaxCP), ls...)
		}
	}

	return nil
}

func parseXML(b []byte, res *multiEngineResult) error {
	if strings.Contains(string(b), "multi-routing-engine-results") {
		return xml.Unmarshal(b, res)
	}

	fi := singleEngineResult{}

	err := xml.Unmarshal(b, &fi)
	if err != nil {
		return err
	}

	res.Results.RoutingEngines = []routingEngine{
		{
			Name:               "N/A",
			PerformanceSummary: fi.PerformanceSummary,
		},
	}
	return nil
}
