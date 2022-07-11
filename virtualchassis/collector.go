package virtualchassis

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_virtualchassis_"

var (
	virtualchassismemberstatus *prometheus.Desc
)

func init() {
	l := []string{"target", "status", "serial", "model", "id", "fpcslot", "role"}
	virtualchassismemberstatus = prometheus.NewDesc(prefix+"member_status", "virtualchassis member-status (1: Prsnt, 0: NotPrsnt)", l, nil)
}

type virtualchassisCollector struct {
}

// Name returns the name of the collector
func (*virtualchassisCollector) Name() string {
	return "virtualchassis"
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &virtualchassisCollector{}
}

// Describe describes the metrics
func (*virtualchassisCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- virtualchassismemberstatus
}

// Collect collects metrics from JunOS
func (c *virtualchassisCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	statusValues := map[string]int{
		"NotPrsnt": 0,
		"Prsnt":   1,
	}

	var x = virtualChassisRpc{}
	if client.Netconf {
		err := client.RunCommandAndParse("<get-virtual-chassis-information/>", &x)
		if err != nil {
			return nil
		}
	} else {
		err := client.RunCommandAndParse("show virtual-chassis", &x)
		if err != nil {
			return err
		}

	for _, m := range x.VirtualChassisInformation.MemberList.Member {
		l := labelValues
		l = append(l, m.Status, m.SerialNumber, m.Model, m.Id, m.FpcSlot, m.Role )
		ch <- prometheus.MustNewConstMetric(virtualchassismemberstatus, prometheus.GaugeValue, float64(statusValues[m.Status]), l...)
	}

	return nil
}
