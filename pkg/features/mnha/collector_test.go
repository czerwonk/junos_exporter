// SPDX-License-Identifier: MIT

package mnha

import (
	"encoding/xml"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
)

// findByDesc returns the metrics matching a given Desc, relying on Desc
// pointer identity (descs are package-level vars created once in init()).
func findByDesc(metrics []prometheus.Metric, desc *prometheus.Desc) []prometheus.Metric {
	var out []prometheus.Metric
	for _, m := range metrics {
		if m.Desc() == desc {
			out = append(out, m)
		}
	}
	return out
}

func labelMap(t *testing.T, m prometheus.Metric) map[string]string {
	t.Helper()

	var dm dto.Metric
	if err := m.Write(&dm); err != nil {
		t.Fatalf("writing metric: %v", err)
	}

	out := make(map[string]string, len(dm.Label))
	for _, lp := range dm.Label {
		out[lp.GetName()] = lp.GetValue()
	}

	return out
}

func metricValue(t *testing.T, m prometheus.Metric) float64 {
	t.Helper()

	var dm dto.Metric
	if err := m.Write(&dm); err != nil {
		t.Fatalf("writing metric: %v", err)
	}

	if dm.Gauge != nil {
		return dm.Gauge.GetValue()
	}

	return dm.Counter.GetValue()
}

func TestCollectDetail(t *testing.T) {
	var res detailResult
	if err := xml.Unmarshal([]byte(detailSample), &res); err != nil {
		t.Fatal(err)
	}

	ch := make(chan prometheus.Metric, 64)
	collectDetail(&res, ch, []string{"srx1"})
	close(ch)

	var metrics []prometheus.Metric
	for m := range ch {
		metrics = append(metrics, m)
	}
	assert.Len(t, metrics, 30, "total metrics emitted")

	nodeMetrics := findByDesc(metrics, nodeStatusDesc)
	assert.Len(t, nodeMetrics, 1)
	assert.Equal(t, 1.0, metricValue(t, nodeMetrics[0]))
	nodeLabels := labelMap(t, nodeMetrics[0])
	assert.Equal(t, "srx1", nodeLabels["target"])
	assert.Equal(t, "0", nodeLabels["grid_id"])
	assert.Equal(t, "1", nodeLabels["local_id"])
	assert.Equal(t, "ONLINE", nodeLabels["status"])

	bfdMetrics := findByDesc(metrics, peerBFDUpDesc)
	assert.Len(t, bfdMetrics, 1)
	assert.Equal(t, 1.0, metricValue(t, bfdMetrics[0]), "peer BFD is UP in the sample")

	detectMetrics := findByDesc(metrics, peerDetectTimeMsDesc)
	assert.Len(t, detectMetrics, 1)
	assert.Equal(t, 3000.0, metricValue(t, detectMetrics[0]))

	coldSyncMetrics := findByDesc(metrics, coldSyncCompleteDesc)
	assert.Len(t, coldSyncMetrics, 1)
	assert.Equal(t, 1.0, metricValue(t, coldSyncMetrics[0]))

	packetMetrics := findByDesc(metrics, peerPacketsTotalDesc)
	assert.Len(t, packetMetrics, 8, "4 packet types x send/receive")

	loopbackMetrics := findByDesc(metrics, loopbackCheckSuccessDesc)
	assert.Len(t, loopbackMetrics, 3, "loopback/nexthop/mbuf for the single PFE in the sample")
	for _, m := range loopbackMetrics {
		assert.Equal(t, 1.0, metricValue(t, m), "all loopback checks succeed in the sample")
	}

	spuUpMetrics := findByDesc(metrics, spuUpDesc)
	assert.Len(t, spuUpMetrics, 1)
	assert.Equal(t, 1.0, metricValue(t, spuUpMetrics[0]))
}

func TestParseDetectTimeMs(t *testing.T) {
	ms, ok := parseDetectTimeMs("3 * 1000ms")
	assert.True(t, ok)
	assert.Equal(t, 3000.0, ms)

	_, ok = parseDetectTimeMs("N/A")
	assert.False(t, ok)
}

func TestOnlineOfflineValue(t *testing.T) {
	assert.Equal(t, 1.0, onlineOfflineValue("ONLINE"))
	assert.Equal(t, 0.0, onlineOfflineValue("OFFLINE"))
	assert.Equal(t, -1.0, onlineOfflineValue("SOMETHING-ELSE"))
}

func TestUpDownValue(t *testing.T) {
	assert.Equal(t, 1.0, upDownValue("UP"))
	assert.Equal(t, 0.0, upDownValue("DOWN"))
}

func TestYesNoValue(t *testing.T) {
	assert.Equal(t, 1.0, yesNoValue("YES"))
	assert.Equal(t, 0.0, yesNoValue("NO"))
}

func TestEnabledDisabledValue(t *testing.T) {
	assert.Equal(t, 1.0, enabledDisabledValue("Enabled"))
	assert.Equal(t, 0.0, enabledDisabledValue("Disabled"))
}
