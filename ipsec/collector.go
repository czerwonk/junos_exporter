package ipsec

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_ipsec_security_associations_"

var (
	blockState        *prometheus.Desc
	activeTunnels     *prometheus.Desc
	configuredTunnels *prometheus.Desc
)

func init() {
	l := []string{"target", "re_name", "description", "name"}

	blockState = prometheus.NewDesc(prefix+"state", "State of the Security Association", l, nil)
	activeTunnels = prometheus.NewDesc(prefix+"active_tunnels", "Total active tunnels", l, nil)
	configuredTunnels = prometheus.NewDesc("junos_ipsec_configured_tunnels", "Total configured tunnels", l, nil)
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
	ch <- configuredTunnels
}

// Collect collects metrics from JunOS
func (c *ipsecCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = RpcReply{}
	err := client.RunCommandAndParse("show security ipsec security-associations", &x)
	if err != nil {
		return err
	}

	for _, re := range x.MultiRoutingEngineResults.RoutingEngine {
		ls := append(labelValues, re.Name, "active tunnels", "")
		ch <- prometheus.MustNewConstMetric(activeTunnels, prometheus.GaugeValue, float64(re.IpSec.ActiveTunnels), ls...)

		for _, block := range re.IpSec.SecurityAssociations {
			c.collectForSecurityAssociation(block, ch, append(labelValues, re.Name))
		}
	}

	var conf = ConfigurationSecurityIpsec{}
	err = client.RunCommandAndParse("show configuration security ipsec", &conf)
	if err != nil {
		return err
	}

	cls := append(labelValues, "N/A", "configured tunnels", "")
	ch <- prometheus.MustNewConstMetric(configuredTunnels, prometheus.GaugeValue, float64(len(conf.Configuration.Security.Ipsec.Vpn)), cls...)

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

func parseXML(b []byte, res *RpcReply) error {
	if strings.Contains(string(b), "multi-routing-engine-results") {
		return xml.Unmarshal(b, res)
	}

	fi := RpcReplyNoRE{}

	err := xml.Unmarshal(b, &fi)
	if err != nil {
		return err
	}

	res.MultiRoutingEngineResults.RoutingEngine = []RoutingEngine{
		{
			Name:  "N/A",
			IpSec: fi.IpSec,
		},
	}
	return nil
}
