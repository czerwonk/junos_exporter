package system

import (
	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_system_"

var (
	mbufsCurrentDesc *prometheus.Desc
	mbufsCacheDesc   *prometheus.Desc
	mbufsTotalDesc   *prometheus.Desc
	mbufsDeniedDesc  *prometheus.Desc

	mbufClustersCurrentDesc *prometheus.Desc
	mbufClustersCacheDesc   *prometheus.Desc
	mbufClustersTotalDesc   *prometheus.Desc
	mbufClustersMaxDesc     *prometheus.Desc
	mbufClustersDeniedDesc  *prometheus.Desc

	mbufClustersFromPacketZoneCurrentDesc *prometheus.Desc
	mbufClustersFromPacketZoneCacheDesc   *prometheus.Desc

	jumboClustersCurrentDesc *prometheus.Desc
	jumboClustersCacheDesc   *prometheus.Desc
	jumboClustersTotalDesc   *prometheus.Desc
	jumboClustersMaxDesc     *prometheus.Desc
	jumboClustersDeniedDesc  *prometheus.Desc

	networkAllocCurrentDesc *prometheus.Desc
	networkAllocCacheDesc   *prometheus.Desc
	networkAllocTotalDesc   *prometheus.Desc

	sfbufsDeniedDesc  *prometheus.Desc
	sfbufsDelayedDesc *prometheus.Desc

	ioInitDesc *prometheus.Desc
)

type systemCollector struct {
}

func init() {
	var l []string

	l = []string{"target"}
	mbufsCurrentDesc = prometheus.NewDesc(prefix+"mbufs_bytes_current", "Current number of bytes in mbufs", l, nil)
	mbufsCacheDesc = prometheus.NewDesc(prefix+"mbufs_bytes_cache", "Cached number of bytes in mbufs", l, nil)
	mbufsTotalDesc = prometheus.NewDesc(prefix+"mbufs_bytes_total", "Total nuumber of bytes in mbufs", l, nil)
	mbufsDeniedDesc = prometheus.NewDesc(prefix+"mbufs_denied_count", "Number of mbuf requests denied", l, nil)

	mbufClustersCurrentDesc = prometheus.NewDesc(prefix+"mbuf_cluster_bytes_current", "Current number of bytes in mbuf clusters", l, nil)
	mbufClustersCacheDesc = prometheus.NewDesc(prefix+"mbuf_cluster_bytes_cache", "Cached number of bytes in mbuf clusters", l, nil)
	mbufClustersTotalDesc = prometheus.NewDesc(prefix+"mbuf_cluster_bytes_total", "Total number of bytes in mbuf clusters", l, nil)
	mbufClustersMaxDesc = prometheus.NewDesc(prefix+"mbuf_cluster_bytes_max", "Max number of bytes in mbuf clusters", l, nil)
	mbufClustersDeniedDesc = prometheus.NewDesc(prefix+"mbufs_and_clusters_denied_count", "", l, nil)

	mbufClustersFromPacketZoneCurrentDesc = prometheus.NewDesc(prefix+"mbuf_and_clusters_from_packet_zone_bytes_current", "Current number of bytes used for mbuf+clusters in packet zone", l, nil)
	mbufClustersFromPacketZoneCacheDesc = prometheus.NewDesc(prefix+"mbuf_and_clusters_from_packet_zone_bytes_cache", "Cached number of bytes used for mbuf+clusters in packet zone", l, nil)

	l = append(l, "page_size")
	jumboClustersCurrentDesc = prometheus.NewDesc(prefix+"jumbo_clusters_current", "Current jumbo clusters in use.", l, nil)
	jumboClustersCacheDesc = prometheus.NewDesc(prefix+"jumbo_clusters_cache", "Cached jumbo clusters in use", l, nil)
	jumboClustersTotalDesc = prometheus.NewDesc(prefix+"jumbo_clusters_total", "Total jumbo clusters in use", l, nil)
	jumboClustersMaxDesc = prometheus.NewDesc(prefix+"jumbo_clusters_max", "Max jumbo clusters in use", l, nil)
	jumboClustersDeniedDesc = prometheus.NewDesc(prefix+"jumbo_clusters_denied_count", "Number of jumbo cluster requests denied", l, nil)

	l = []string{"target"}
	networkAllocCurrentDesc = prometheus.NewDesc(prefix+"network_allocated_bytes_current", "Current number of bytes allocated for network", l, nil)
	networkAllocCacheDesc = prometheus.NewDesc(prefix+"network_allocated_bytes_cache", "Cached number of bytes allocated for network", l, nil)
	networkAllocTotalDesc = prometheus.NewDesc(prefix+"network_allocated_bytes_total", "Total number of bytes allocated for network", l, nil)

	sfbufsDeniedDesc = prometheus.NewDesc(prefix+"sfbufs_denied_count", "Number of sfbuf requests denied", l, nil)
	sfbufsDelayedDesc = prometheus.NewDesc(prefix+"sfbufs_delayed_count", "Number of sfbuf requests delayed", l, nil)

	ioInitDesc = prometheus.NewDesc(prefix+"io_requests_count", "Number of I/O requests initiated", l, nil)
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &systemCollector{}
}

// Name returns the name of the collector
func (*systemCollector) Name() string {
	return "System"
}

// Describe describes the metrics
func (*systemCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- mbufsCurrentDesc
	ch <- mbufsCacheDesc
	ch <- mbufsTotalDesc
	ch <- mbufsDeniedDesc
	ch <- mbufClustersCurrentDesc
	ch <- mbufClustersCacheDesc
	ch <- mbufClustersTotalDesc
	ch <- mbufClustersMaxDesc
	ch <- mbufClustersDeniedDesc
	ch <- mbufClustersFromPacketZoneCurrentDesc
	ch <- mbufClustersFromPacketZoneCacheDesc
	ch <- jumboClustersCurrentDesc
	ch <- jumboClustersCacheDesc
	ch <- jumboClustersTotalDesc
	ch <- jumboClustersMaxDesc
	ch <- jumboClustersDeniedDesc
	ch <- networkAllocCurrentDesc
	ch <- networkAllocCacheDesc
	ch <- networkAllocTotalDesc
	ch <- sfbufsDeniedDesc
	ch <- sfbufsDelayedDesc
	ch <- ioInitDesc
}

// Collect collects metrics from JunOS
func (c *systemCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var (
		err error
	)

	err = c.CollectSystem(client, ch, labelValues)
	if err != nil {
		return err
	}

	return nil
}

func (c *systemCollector) CollectSystem(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var (
		r   *BuffersRPC
		l   []string
		err error
	)

	r = new(BuffersRPC)

	err = client.RunCommandAndParse("show system buffers", r)
	if err != nil {
		return err
	}

	ch <- prometheus.MustNewConstMetric(mbufsCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufsCurrent), labelValues...)
	ch <- prometheus.MustNewConstMetric(mbufsCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufsCache), labelValues...)
	ch <- prometheus.MustNewConstMetric(mbufsTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufsTotal), labelValues...)
	ch <- prometheus.MustNewConstMetric(mbufsDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufsDenied), labelValues...)

	ch <- prometheus.MustNewConstMetric(mbufClustersCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersCurrent), labelValues...)
	ch <- prometheus.MustNewConstMetric(mbufClustersCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersCache), labelValues...)
	ch <- prometheus.MustNewConstMetric(mbufClustersTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersTotal), labelValues...)
	ch <- prometheus.MustNewConstMetric(mbufClustersMaxDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersMax), labelValues...)
	ch <- prometheus.MustNewConstMetric(mbufClustersDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersDeniedDesc), labelValues...)

	ch <- prometheus.MustNewConstMetric(mbufClustersFromPacketZoneCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersFromPacketZoneCurrent), labelValues...)
	ch <- prometheus.MustNewConstMetric(mbufClustersFromPacketZoneCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.MbufClustersFromPacketZoneCache), labelValues...)

	l = append(labelValues, "4k")
	ch <- prometheus.MustNewConstMetric(jumboClustersCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCurrent4K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCache4K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersTotal4K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersMaxDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersMax4K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersDenied4K), l...)

	l = append(labelValues, "9k")
	ch <- prometheus.MustNewConstMetric(jumboClustersCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCurrent9K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCache9K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersTotal9K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersMaxDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersMax9K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersDenied9K), l...)

	l = append(labelValues, "16k")
	ch <- prometheus.MustNewConstMetric(jumboClustersCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCurrent16K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersCache16K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersTotal16K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersMaxDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersMax16K), l...)
	ch <- prometheus.MustNewConstMetric(jumboClustersDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.JumboClustersDenied16K), l...)

	ch <- prometheus.MustNewConstMetric(sfbufsDeniedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.SfbufsDenied), labelValues...)
	ch <- prometheus.MustNewConstMetric(sfbufsDelayedDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.SfbufsDelayed), labelValues...)
	ch <- prometheus.MustNewConstMetric(ioInitDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.IoInit), labelValues...)

	// network alloc values seem to be Kb
	ch <- prometheus.MustNewConstMetric(networkAllocCurrentDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.NetworkAllocCurrent*1024), labelValues...)
	ch <- prometheus.MustNewConstMetric(networkAllocCacheDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.NetworkAllocCache*1024), labelValues...)
	ch <- prometheus.MustNewConstMetric(networkAllocTotalDesc, prometheus.GaugeValue, float64(r.MemoryStatistics.NetworkAllocTotal*1024), labelValues...)

	return nil
}
