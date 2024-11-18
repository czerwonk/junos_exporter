package poe

import (
	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"strings"
)

const prefix string = "junos_poe_"

var (
	poeEnabledDesc    *prometheus.Desc
	poeStatusDesc     *prometheus.Desc
	poePowerLimitDesc *prometheus.Desc
	poePowerDesc      *prometheus.Desc
	poeClassDesc      *prometheus.Desc
)

// Initialize metrics descriptions
func init() {
	labels := []string{"interface"}
	poeEnabledDesc = prometheus.NewDesc(prefix+"enabled", "Information about interface status", labels, nil)
	poeStatusDesc = prometheus.NewDesc(prefix+"status", "Information about interface PoE status", labels, nil)
	poePowerLimitDesc = prometheus.NewDesc(prefix+"power_limit", "Information about interface PoE power limit", labels, nil)
	poePowerDesc = prometheus.NewDesc(prefix+"power", "Information about interface PoE power usage", labels, nil)
	poeClassDesc = prometheus.NewDesc(prefix+"class", "Information about interface PoE class", labels, nil)
}

type poeCollector struct{}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &poeCollector{}
}

// Name returns the name of the collector
func (p poeCollector) Name() string {
	return "poe"
}

// Describe describes the metrics
func (p poeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- poeEnabledDesc
	ch <- poeStatusDesc
	ch <- poePowerLimitDesc
	ch <- poePowerDesc
	ch <- poeClassDesc
}

func (p poeCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var result poeInterfaceResult
	err := client.RunCommandAndParse("show poe interface", &result)
	if err != nil {
		return errors.Wrap(err, "failed to run command 'show poe interface'")
	}
	for _, i := range result.Poe.InterfaceInformation {
		p.CollectForInterface(i, ch, labelValues)
	}
	return nil
}

func (p *poeCollector) CollectForInterface(iface InterfaceInformation, ch chan<- prometheus.Metric, labelValues []string) {
	lv := append(labelValues, []string{iface.Name}...)

	enabled := 0
	if iface.Enabled == "Enabled" {
		enabled = 1
	}
	status := 0
	if iface.Status == "ON" {
		status = 1
	}
	powerLimit := parsePower(iface.PowerLimit)
	powerUsage := parsePower(iface.Power)

	class := -1.0
	if iface.Class != "not-applicable" {
		pClass, _ := strconv.ParseFloat(iface.Class, 64)
		class = pClass
	}

	ch <- prometheus.MustNewConstMetric(poeEnabledDesc, prometheus.GaugeValue, float64(enabled), lv...)
	ch <- prometheus.MustNewConstMetric(poeStatusDesc, prometheus.GaugeValue, float64(status), lv...)
	ch <- prometheus.MustNewConstMetric(poePowerLimitDesc, prometheus.GaugeValue, powerLimit, lv...)
	ch <- prometheus.MustNewConstMetric(poePowerDesc, prometheus.GaugeValue, powerUsage, lv...)
	ch <- prometheus.MustNewConstMetric(poeClassDesc, prometheus.GaugeValue, class, lv...)
}

func parsePower(s string) float64 {
	powerWithoutSuffix := strings.ReplaceAll(s, "W", "")
	parsedPower, _ := strconv.ParseFloat(powerWithoutSuffix, 64)
	return parsedPower
}
