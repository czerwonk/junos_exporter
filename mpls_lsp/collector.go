package mpls_lsp

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_mpls_lsp_"

var (
	mpls_lspState         *prometheus.Desc
	mpls_lspPathState     *prometheus.Desc
	mpls_lspPathFlapCount *prometheus.Desc

	mpls_lspStateMap = map[string]int{
		"Dn": 0,
		"Up": 1,
	}
)

func init() {
	ls := []string{"target", "lspname", "lspsrc", "lspdst"}
	mpls_lspState = prometheus.NewDesc(prefix+"state", "mpls_lsp state (0: down, 1:up)", ls, nil)

	lps := []string{"target", "lspname", "lspsrc", "lspdst", "title", "name"}
	mpls_lspPathState = prometheus.NewDesc(prefix+"path_state", "mpls_lsp pathstate (0: down, 1:up)", lps, nil)
	mpls_lspPathFlapCount = prometheus.NewDesc(prefix+"path_flapcount", "mpls_lsp path flap count", lps, nil)
}

type mpls_lspCollector struct {
}

// Name returns the name of the collector
func (*mpls_lspCollector) Name() string {
	return "mpls_lsp"
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &mpls_lspCollector{}
}

// Describe describes the metrics
func (*mpls_lspCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- mpls_lspState
}

// Collect collects metrics from JunOS
func (c *mpls_lspCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
        var x = mpls_lspRpc{}
//        err := client.RunCommandAndParse("show mpls lsp ingress extensive", &x) //ingress:Display LSPs originating at this router
        err := client.RunCommandAndParse("<get-mpls-lsp-information><ingress/><extensive/></get-mpls-lsp-information>", &x) //ingress:Display LSPs originating at this router
	if err != nil {
		return err
	}

	for _, lsp := range x.Information.Sessions {
		l := append(labelValues, lsp.Name, lsp.SrcIP, lsp.DstIP)
		ch <- prometheus.MustNewConstMetric(mpls_lspState, prometheus.GaugeValue, float64(mpls_lspStateMap[lsp.LSPState]), l...)

		for _, path := range lsp.Path {
			l := append(labelValues, lsp.Name, lsp.SrcIP, lsp.DstIP, path.Title, path.Name)
			ch <- prometheus.MustNewConstMetric(mpls_lspPathState, prometheus.GaugeValue, float64(mpls_lspStateMap[path.State]), l...)
			ch <- prometheus.MustNewConstMetric(mpls_lspPathFlapCount, prometheus.GaugeValue, float64(path.FlapCount), l...)
		}
	}

	return nil
}
