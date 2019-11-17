package ipsec

import (
	"fmt"

	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_ipsec_security_associations_"

var (
	blockState    *prometheus.Desc
	activeTunnels *prometheus.Desc
)

func init() {
	l := []string{"target", "description", "name"}

	blockState = prometheus.NewDesc(prefix+"state", "State of the Security Association", l, nil)
	activeTunnels = prometheus.NewDesc(prefix+"active_tunnels", "Total active tunnels", l, nil)
}

type ipsecCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &ipsecCollector{}
}

// Name returns the name of the collector
func (*ipsecCollector) Name() string {
	return "IPSec"
}

// Describe describes the metrics
func (*ipsecCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- blockState
	ch <- activeTunnels
}

// Collect collects metrics from JunOS
func (c *ipsecCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = IpsecRpc{}
	err := client.RunCommandAndParse("show security ipsec security-associations", &x)
	if err != nil {
		return err
	}

	ls := append(labelValues, "active tunnels", "")
	ch <- prometheus.MustNewConstMetric(activeTunnels, prometheus.GaugeValue, float64(x.Information.ActiveTunnels), ls...)

	for _, block := range x.Information.SecurityAssociations {
		c.collectForSecurityAssociation(block, ch, labelValues)
	}

	return nil
}

func (c *ipsecCollector) collectForSecurityAssociation(block IpsecSecurityAssociationBlock, ch chan<- prometheus.Metric, labelValues []string) {
	// build SA name
	var saName string
	var saDesc string
	for _, sa := range block.SecurityAssociations {
		saName = sa.RemoteGateway
		saDesc = fmt.Sprintf("security association for remote gateway %s", sa.RemoteGateway)
	}
	lp := append(labelValues, saDesc, saName)
	stateVal := stateToInt(&block.State)
	ch <- prometheus.MustNewConstMetric(blockState, prometheus.GaugeValue, float64(stateVal), lp...)
}

func stateToInt(state *string) int {
	retval := 0

	if *state == "up" {
		retval = 1
	}

	return retval
}
