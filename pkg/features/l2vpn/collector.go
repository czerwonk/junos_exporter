package l2vpn

import (
	"strconv"
	"strings"
	"time"

	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	l2vpnConnectionStateDesc *prometheus.Desc
	l2vpnConnectionsDesc     *prometheus.Desc
	l2vpnMap                 = map[string]int{
		"EI":    0,
		"EM":    1,
		"VC-Dn": 2,
		"CM":    3,
		"CN":    4,
		"OR":    5,
		"OL":    6,
		"LD":    7,
		"RD":    8,
		"LN":    9,
		"RN":    10,
		"XX":    11,
		"MM":    12,
		"BK":    13,
		"PF":    14,
		"RS":    15,
		"LB":    16,
		"VM":    17,
		"NC":    18,
		"WE":    19,
		"NP":    20,
		"->":    21,
		"<-":    22,
		"Up":    23,
		"Dn":    24,
		"CF":    25,
		"SC":    26,
		"LM":    27,
		"RM":    28,
		"IL":    29,
		"MI":    20,
		"ST":    21,
		"PB":    22,
		"SN":    23,
		"RB":    24,
		"HS":    25,
	}
)

func init() {
	l2vpnPrefix := "junos_l2vpn_"

	lcount := []string{"target", "routing_instance"}
	lstate := []string{"target", "routing_instance", "connection_id", "remote_pe", "last_change", "up_transitions", "local_interface_name"}
	l2StateDescription := "A l2vpn can have one of the following state-mappings EI: 0, EM: 1, VC-Dn: 2, CM: 3, CN: 4, OR: 5, OL: 6, LD: 7, RD: 8, LN: 9, RN: 10, XX: 11, MM: 12, BK: 13, PF: 14, RS: 15, LB: 16, VM: 17, NC: 18, WE: 19, NP: 20, ->: 21, <-: 22, Up: 23, Dn: 24, CF: 25, SC: 26, LM: 27, RM: 28, IL: 39, MI: 30, ST: 31, PB: 32, SN: 33, RB: 34, HS: 35"
	l2vpnConnectionsDesc = prometheus.NewDesc(l2vpnPrefix+"connection_count", "Number of l2vpn connections", lcount, nil)
	l2vpnConnectionStateDesc = prometheus.NewDesc(l2vpnPrefix+"connection_status", l2StateDescription, lstate, nil)
}

// Collector collects l2vpn metrics
type l2vpnCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &l2vpnCollector{}
}

// Name returns the name of the collector
func (*l2vpnCollector) Name() string {
	return "L2 Circuit"
}

// Describe describes the metrics
func (*l2vpnCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- l2vpnConnectionStateDesc
	ch <- l2vpnConnectionsDesc
}

// Collect collects metrics from JunOS
func (c *l2vpnCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	return c.collectl2vpnMetrics(client, ch, labelValues)
}

func (c *l2vpnCollector) collectl2vpnMetrics(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = l2vpnRpc{}
	err := client.RunCommandAndParse("show l2vpn connections", &x)
	if err != nil {
		return err
	}

	instances := x.Information.RoutingInstances

	for i := 0; i < len(instances); i++ {
		connCount := 0
		for s := 0; s < len(instances[i].ReferenceSite); s++ {
			connCount += +len(instances[i].ReferenceSite[s].Connections)
		}
		l := append(labelValues, instances[i].RoutingInstanceName)
		ch <- prometheus.MustNewConstMetric(l2vpnConnectionsDesc, prometheus.GaugeValue, float64(connCount), l...)

	}

	for _, a := range instances {
		for _, site := range a.ReferenceSite {
			// l = append(l, site.ID)
			for _, conn := range site.Connections {
				l := append(labelValues, a.RoutingInstanceName)
				// l = append(l, site.ID)
				c.collectForConnection(ch, conn, l)
			}
		}
	}

	return nil
}

func (c *l2vpnCollector) collectForConnection(ch chan<- prometheus.Metric,
	conn l2vpnConnection, labelValues []string) {
	id := conn.ID
	remotePe := conn.RemotePe
	lastChange := string_to_date(conn.LastChange)
	upTransitions := conn.UpTransitions
	localInterface := ""
	if len(conn.LocalInterface) == 1 {
		localInterface = conn.LocalInterface[0].Name
	}

	l := append(labelValues, id, remotePe, lastChange, upTransitions, localInterface)
	state := l2vpnMap[conn.StatusString]

	ch <- prometheus.MustNewConstMetric(l2vpnConnectionStateDesc, prometheus.GaugeValue, float64(state), l...)
}

func string_to_date(date string) string {
	layout := "Jan 2 15:04:05 2006"
	t, err := time.Parse(layout, strings.TrimRight(date, " \n"))
	if err != nil {
		return ""
	}
	return strconv.FormatInt(t.Unix(), 10)
}
