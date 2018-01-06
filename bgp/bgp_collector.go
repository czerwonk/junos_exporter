package bgp

import "github.com/prometheus/client_golang/prometheus"

type BgpCollector struct {
}

func (*BgpCollector) Describe(ch chan<- *prometheus.Desc) {

}

func (c *BgpCollector) Collect(datasource BgpDatasource, ch chan<- prometheus.Metric) error {
	sessions, err := datasource.BgpSessions()
	if err != nil {
		return err
	}

	for _, s := range sessions {
		c.collectForSession(s, ch)
	}

	return nil
}

func (*BgpCollector) collectForSession(s *BgpSession, ch chan<- prometheus.Metric) {

}
