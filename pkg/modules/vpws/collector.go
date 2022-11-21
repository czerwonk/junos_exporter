package vpws

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_vpws_"

var (
	vpwsStatus *prometheus.Desc
	vpwsSid    *prometheus.Desc

	vpwsStatusMap = map[string]int{
		"Down": 0,
		"Up":   1,
	}

	vpwsSidMap = map[string]int{
		"Unresolved": 0,
		"Resolved":   1,
	}
)

func init() {
	l := []string{"target", "vpwsinstance", "rd", "interface", "esi", "mode", "role"}
	vpwsStatus = prometheus.NewDesc(prefix+"status", "vpws status (0: down, 1:up)", l, nil)

	ls := []string{"target", "vpwsinstance", "rd", "interface", "sidorigin", "sid", "ip", "esi", "mode", "role"}
	vpwsSid = prometheus.NewDesc(prefix+"sid", "vpws sid (0: Unresolved, 1:Resolved)", ls, nil)
}

type vpwsCollector struct {
}

// Name returns the name of the collector
func (*vpwsCollector) Name() string {
	return "vpws"
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &vpwsCollector{}
}

// Describe describes the metrics
func (*vpwsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- vpwsStatus
	ch <- vpwsSid
}

// Collect collects metrics from JunOS
func (c *vpwsCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = result{}
	err := client.RunCommandAndParse("show evpn vpws-instance", &x)
	if err != nil {
		return err
	}

	for _, vInst := range x.Information.VpwsInstances {
		for _, vIf := range vInst.Interfaces {
			l := append(labelValues, vInst.Name, vInst.RD, vIf.Name, vIf.Esi, vIf.Mode, vIf.Role)
			ch <- prometheus.MustNewConstMetric(vpwsStatus, prometheus.GaugeValue, float64(vpwsStatusMap[vIf.Status]), l...)

			for _, vSid := range vIf.LocalStatus.SidPeInfo {
				l := append(labelValues, vInst.Name, vInst.RD, vIf.Name, "local", vIf.LocalStatus.Sid, vSid.IP, vSid.Esi, vSid.Mode, vSid.Role)
				ch <- prometheus.MustNewConstMetric(vpwsSid, prometheus.GaugeValue, float64(vpwsSidMap[vSid.Status]), l...)
			}

			for _, vSid := range vIf.RemoteStatus.SidPeInfo {
				l := append(labelValues, vInst.Name, vInst.RD, vIf.Name, "remote", vIf.RemoteStatus.Sid, vSid.IP, vSid.Esi, vSid.Mode, vSid.Role)
				ch <- prometheus.MustNewConstMetric(vpwsSid, prometheus.GaugeValue, float64(vpwsSidMap[vSid.Status]), l...)
			}

		}
	}

	return nil
}
