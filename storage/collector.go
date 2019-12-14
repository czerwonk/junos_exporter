package storage

import (
	"encoding/xml"
	"strconv"
	"strings"

	"github.com/czerwonk/junos_exporter/collector"
	"github.com/czerwonk/junos_exporter/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix string = "junos_storage_"

var (
	totalBlocksDesc     *prometheus.Desc
	usedBlocksDesc      *prometheus.Desc
	availableBlocksDesc *prometheus.Desc
	usedPercentDesc     *prometheus.Desc
)

func init() {
	l := []string{"target", "device", "re_name", "mountpoint"}
	totalBlocksDesc = prometheus.NewDesc(prefix+"total_blocks_count", "Total number of blocks", l, nil)
	usedBlocksDesc = prometheus.NewDesc(prefix+"used_blocks_count", "Number of used blocks", l, nil)
	availableBlocksDesc = prometheus.NewDesc(prefix+"available_blocks_count", "Number of available blocks", l, nil)
	usedPercentDesc = prometheus.NewDesc(prefix+"used_percent", "Percent of used storage", l, nil)
}

type storageCollector struct {
}

// NewCollector creates a new collector
func NewCollector() collector.RPCCollector {
	return &storageCollector{}
}

// Name returns the name of the collector
func (*storageCollector) Name() string {
	return "Storage"
}

// Describe describes the metrics
func (*storageCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- totalBlocksDesc
	ch <- usedBlocksDesc
	ch <- availableBlocksDesc
	ch <- usedPercentDesc
}

// Collect collects metrics from JunOS
func (c *storageCollector) Collect(client *rpc.Client, ch chan<- prometheus.Metric, labelValues []string) error {
	var x = MultiRoutingEngineResults{}
	err := client.RunCommandAndParseWithParser("show system storage", func(b []byte) error {
		return parseXML(b, &x)
	})
	if err != nil {
		return err
	}

	for _, re := range x.Results {
		for _, f := range re.Storage.Information.Filesystems {
			l := append(labelValues, f.FilesystemName, re.Name, f.MountedOn)

			ch <- prometheus.MustNewConstMetric(totalBlocksDesc, prometheus.GaugeValue, float64(f.TotalBlocks), l...)
			ch <- prometheus.MustNewConstMetric(usedBlocksDesc, prometheus.GaugeValue, float64(f.UsedBlocks), l...)
			ch <- prometheus.MustNewConstMetric(availableBlocksDesc, prometheus.GaugeValue, float64(f.AvailableBlocks), l...)
			percent := strings.TrimSpace(f.UsedPercent)
			value, err := strconv.ParseFloat(percent, 64)
			if err != nil {
				value = 0
			}
			ch <- prometheus.MustNewConstMetric(usedPercentDesc, prometheus.GaugeValue, value, l...)
		}
	}

	return nil
}

func parseXML(b []byte, res *MultiRoutingEngineResults) error {
	if strings.Contains(string(b), "multi-routing-engine-results") {
		return xml.Unmarshal(b, res)
	}

	si := StorageInformation{}
	err := xml.Unmarshal(b, &si)
	if err != nil {
		return err
	}

	res.Results = []MultiRoutingEngineItem{
		MultiRoutingEngineItem{
			Name:    "",
			Storage: si,
		},
	}
	return nil
}
