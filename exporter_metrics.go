// SPDX-License-Identifier: MIT

package main

import (
	"net/http"
	"path"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	promcollectors "github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	log "github.com/sirupsen/logrus"
)

// exporterMetrics holds the metrics describing the exporter process itself.
// They live in a registry created once at startup, separate from the registry
// built per scrape for device metrics, because Go runtime counters and the
// reload gauges describe state that spans requests.
type exporterMetrics struct {
	registry               *prometheus.Registry
	reloadSuccessful       prometheus.Gauge
	reloadSuccessTimestamp prometheus.Gauge
}

func newExporterMetrics() *exporterMetrics {
	m := &exporterMetrics{
		registry: prometheus.NewRegistry(),
		reloadSuccessful: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "junos_exporter_config_last_reload_successful",
			Help: "Whether the last configuration reload attempt was successful.",
		}),
		reloadSuccessTimestamp: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "junos_exporter_config_last_reload_success_timestamp_seconds",
			Help: "Timestamp of the last successful configuration reload.",
		}),
	}

	m.registry.MustRegister(
		promcollectors.NewGoCollector(),
		promcollectors.NewProcessCollector(promcollectors.ProcessCollectorOpts{}),
		m.reloadSuccessful,
		m.reloadSuccessTimestamp,
	)

	return m
}

// recordReload reports the outcome of a configuration load. The timestamp is
// only advanced on success, so it keeps pointing at the configuration the
// exporter is actually serving.
func (m *exporterMetrics) recordReload(err error) {
	if err != nil {
		m.reloadSuccessful.Set(0)
		return
	}

	m.reloadSuccessful.Set(1)
	m.reloadSuccessTimestamp.Set(float64(time.Now().Unix()))
}

func (m *exporterMetrics) handler() http.Handler {
	l := log.New()
	l.Level = log.ErrorLevel

	return promhttp.HandlerFor(m.registry, promhttp.HandlerOpts{
		ErrorLog:      l,
		ErrorHandling: promhttp.ContinueOnError,
	})
}

// exporterMetricsPath derives the path serving the exporter's own metrics from
// the path serving device metrics. The latter takes a target parameter, so
// serving both from it would repeat every runtime metric once per device.
func exporterMetricsPath(telemetryPath string) string {
	return path.Join(telemetryPath, "exporter")
}
