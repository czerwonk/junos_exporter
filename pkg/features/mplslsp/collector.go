package mplslsp

import (
	"github.com/czerwonk/junos_exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "junos_mpls_lsp_"

var (
	lspState         *prometheus.Desc
	lspPathState     *prometheus.Desc
	lspPathFlapCount *prometheus.Desc

	lspStateMap = map[string]int{
		"Dn": 0,
		"Up": 1,
	}
)

func init() {
	ls := []string{"target", "lspname", "lspsrc", "lspdst"}
	lspState = prometheus.NewDesc(prefix+"state", "mpls_lsp state (0: down, 1:up)", ls, nil)

	lps := []string{"target", "lspname", "lspsrc", "lspdst", "title", "name"}
	lspPathState = prometheus.NewDesc(prefix+"path_state", "mpls_lsp pathstate (0: down, 1:up)", lps, nil)
	lspPathFlapCount = prometheus.NewDesc(prefix+"path_flapcount", "mpls_lsp path flap count", lps, nil)
}

type mplsLSPCollector struct {
}

// Name returns the name of the collector
func (*mplsLSPCollector) Name() string {
	return "mpls_lsp"
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &mplsLSPCollector{}
}

// Describe describes the metrics
func (*mplsLSPCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- lspState
}

// Collect collects metrics from JunOS
func (c *mplsLSPCollector) Collect(client collector.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = result{}
	err := client.RunCommandAndParse("show mpls lsp ingress extensive", &x) //ingress:Display LSPs originating at this router
	if err != nil {
		return err
	}

	for _, lsp := range x.Information.Sessions {
		l := append(labelValues, lsp.Name, lsp.SrcIP, lsp.DstIP)
		ch <- prometheus.MustNewConstMetric(lspState, prometheus.GaugeValue, float64(lspStateMap[lsp.LSPState]), l...)

		for _, path := range lsp.Path {
			l := append(labelValues, lsp.Name, lsp.SrcIP, lsp.DstIP, path.Title, path.Name)
			ch <- prometheus.MustNewConstMetric(lspPathState, prometheus.GaugeValue, float64(lspStateMap[path.State]), l...)
			ch <- prometheus.MustNewConstMetric(lspPathFlapCount, prometheus.GaugeValue, float64(path.FlapCount), l...)
		}
	}

	return nil
}
