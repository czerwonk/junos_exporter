// SPDX-License-Identifier: MIT

package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func gaugeValue(t *testing.T, reg *prometheus.Registry, name string) float64 {
	t.Helper()

	families, err := reg.Gather()
	assert.NoError(t, err)

	for _, f := range families {
		if f.GetName() != name {
			continue
		}

		metrics := f.GetMetric()
		assert.Len(t, metrics, 1, "expected exactly one series for %s", name)

		return metrics[0].GetGauge().GetValue()
	}

	t.Fatalf("metric %s not found", name)

	return 0
}

func TestExporterMetricsPath(t *testing.T) {
	tests := []struct {
		name          string
		telemetryPath string
		expected      string
	}{
		{"default", "/metrics", "/metrics/exporter"},
		{"custom", "/junos", "/junos/exporter"},
		{"trailing slash", "/metrics/", "/metrics/exporter"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, exporterMetricsPath(tt.telemetryPath))
		})
	}
}

func TestRecordReloadSuccess(t *testing.T) {
	m := newExporterMetrics()

	before := float64(time.Now().Unix())
	m.recordReload(nil)

	assert.Equal(t, 1.0, gaugeValue(t, m.registry, "junos_exporter_config_last_reload_successful"))

	ts := gaugeValue(t, m.registry, "junos_exporter_config_last_reload_success_timestamp_seconds")
	assert.GreaterOrEqual(t, ts, before)
}

func TestRecordReloadFailureKeepsLastSuccessTimestamp(t *testing.T) {
	m := newExporterMetrics()

	m.recordReload(nil)
	ts := gaugeValue(t, m.registry, "junos_exporter_config_last_reload_success_timestamp_seconds")

	m.recordReload(errors.New("bad config"))

	assert.Equal(t, 0.0, gaugeValue(t, m.registry, "junos_exporter_config_last_reload_successful"))
	assert.Equal(t, ts, gaugeValue(t, m.registry, "junos_exporter_config_last_reload_success_timestamp_seconds"),
		"a failed reload must not advance the last successful reload timestamp")
}

func TestRecordReloadFailureBeforeAnySuccess(t *testing.T) {
	m := newExporterMetrics()

	m.recordReload(errors.New("bad config"))

	assert.Equal(t, 0.0, gaugeValue(t, m.registry, "junos_exporter_config_last_reload_successful"))
	assert.Equal(t, 0.0, gaugeValue(t, m.registry, "junos_exporter_config_last_reload_success_timestamp_seconds"))
}

func TestExporterMetricsHandlerServesRuntimeAndReloadMetrics(t *testing.T) {
	m := newExporterMetrics()
	m.recordReload(nil)

	rec := httptest.NewRecorder()
	m.handler().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/metrics/exporter", nil))

	assert.Equal(t, http.StatusOK, rec.Code)

	body := rec.Body.String()
	assert.Contains(t, body, "go_goroutines")
	assert.Contains(t, body, "junos_exporter_config_last_reload_successful 1")
	assert.Contains(t, body, "junos_exporter_config_last_reload_success_timestamp_seconds")
}

func TestExporterMetricsHandlerServesNoDeviceMetrics(t *testing.T) {
	m := newExporterMetrics()

	rec := httptest.NewRecorder()
	m.handler().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/metrics/exporter", nil))

	assert.NotContains(t, rec.Body.String(), "junos_up")
}

func TestRegisterRoutesServesExporterMetrics(t *testing.T) {
	selfMetrics = newExporterMetrics()

	mux := http.NewServeMux()
	registerRoutes(mux)

	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, exporterMetricsPath(*metricsPath), nil))

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "go_goroutines")
}

func TestRegisterRoutesLinksExporterMetricsFromIndex(t *testing.T) {
	selfMetrics = newExporterMetrics()

	mux := http.NewServeMux()
	registerRoutes(mux)

	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), `href="`+exporterMetricsPath(*metricsPath)+`"`)
}
