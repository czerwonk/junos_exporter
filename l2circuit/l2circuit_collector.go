package l2circuit

import (
	"regexp"

	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	l2circuitConnectionStateDesc *prometheus.Desc
	l2circuitConnectionsDesc     *prometheus.Desc
	l2circuitMap                 = map[string]int{
		"EI":    0,
		"MM":    1,
		"EM":    2,
		"CM":    3,
		"VM":    4,
		"OL":    5,
		"NC":    6,
		"BK":    7,
		"CB":    8,
		"LD":    9,
		"RD":    10,
		"XX":    11,
		"NP":    12,
		"Dn":    13,
		"VC-Dn": 14,
		"Up":    15,
		"CF":    16,
		"IB":    17,
		"TM":    18,
		"ST":    19,
		"SP":    20,
		"RS":    21,
		"HS":    22,
	}
)

func init() {
	l2circuitPrefix := "junos_l2circuit_"

	l := []string{"target", "address", "vcid"}

	l2StateDescription := "A l2circuit can have one of the following state-mappings EI: 0,MM: 1,EM: 2,CM: 3,VM: 4,OL: 5,NC: 6,BK: 7,CB: 8,LD: 9,RD: 10,XX: 11,NP: 12,Dn: 13,VC-Dn: 14, Up: 15, CF: 16,IB: 17,TM: 18,ST: 19,SP: 20,RS: 21,HS: 22"

	l2circuitConnectionsDesc = prometheus.NewDesc(l2circuitPrefix+"connection_count", "Number of L2Circuits", l, nil)

	l2circuitConnectionStateDesc = prometheus.NewDesc(l2circuitPrefix+"connection_status", l2StateDescription, l, nil)
}

// Collector collects L2CIRCUIT metrics
type l2circuitCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &l2circuitCollector{}
}

// Describe describes the metrics
func (*l2circuitCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- l2circuitConnectionStateDesc
	ch <- l2circuitConnectionsDesc
}

// Collect collects metrics from JunOS
func (c *l2circuitCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	return c.collectL2circuitMetrics(client, ch, labelValues)
}

func (c *l2circuitCollector) collectL2circuitMetrics(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = L2circuitRpc{}
	err := client.RunCommandAndParse("show l2circuit connections brief", &x)
	if err != nil {
		return err
	}

	neighbors := x.Information.Neighbors

	conncount := 0
	for i := 0; i < len(neighbors); i++ {
		conncount = conncount + len(neighbors[i].Connections)
	}

	for _, a := range neighbors {
		l := append(labelValues, a.Address)
		for _, b := range a.Connections {
			idstr := b.ID
			re := regexp.MustCompile(`\(vc ([0-9]+)\)`)
			idint := re.FindStringSubmatch(idstr)
			id := idint[len(idint)-1]
			l := append(l, id)
			state := l2circuitMap[b.StatusString]
			ch <- prometheus.MustNewConstMetric(l2circuitConnectionsDesc, prometheus.GaugeValue, float64(conncount), l...)
			ch <- prometheus.MustNewConstMetric(l2circuitConnectionStateDesc, prometheus.GaugeValue, float64(state), l...)
		}
	}

	return nil
}
