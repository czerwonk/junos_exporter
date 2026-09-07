// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"errors"
	"io"
	"runtime"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"

	"github.com/czerwonk/junos_exporter/pkg/collector"
)

func TestBuildInfoMetric(t *testing.T) {
	m := buildInfoMetric()

	if !strings.Contains(m.Desc().String(), prefix+"exporter_build_info") {
		t.Fatalf("unexpected desc: %s", m.Desc().String())
	}

	d := &dto.Metric{}
	if err := m.Write(d); err != nil {
		t.Fatalf("could not write metric: %v", err)
	}

	if got := d.GetGauge().GetValue(); got != 1 {
		t.Errorf("build_info value = %v, want 1", got)
	}

	labels := make(map[string]string)
	for _, lp := range d.GetLabel() {
		labels[lp.GetName()] = lp.GetValue()
	}

	if got := labels["version"]; got != version {
		t.Errorf("version label = %q, want %q", got, version)
	}

	// goversion comes from the runtime, so it is always populated and proves the
	// labels are wired rather than empty placeholders.
	if labels["goversion"] != runtime.Version() {
		t.Errorf("goversion label = %q, want %q", labels["goversion"], runtime.Version())
	}

	for _, name := range []string{"revision", "branch"} {
		if _, ok := labels[name]; !ok {
			t.Errorf("missing %q label", name)
		}
	}
}

// fakeRPCCollector reports a configurable result without touching a device.
type fakeRPCCollector struct {
	name string
	err  error
}

func (c *fakeRPCCollector) Name() string { return c.name }

func (c *fakeRPCCollector) Describe(_ chan<- *prometheus.Desc) {}

func (c *fakeRPCCollector) Collect(_ collector.Client, _ chan<- prometheus.Metric, _ []string) error {
	return c.err
}

// collectSuccessValue runs one collector and returns the value of the
// junos_collect_success sample it emitted.
func collectSuccessValue(t *testing.T, collectErr error) float64 {
	t.Helper()

	jc := &junosCollector{}
	ch := make(chan prometheus.Metric, 16)

	jc.collectWithCollector(
		context.Background(),
		&fakeRPCCollector{name: "fake", err: collectErr},
		nil,
		ch,
		[]string{"device1"},
	)
	close(ch)

	for m := range ch {
		if !strings.Contains(m.Desc().String(), prefix+"collect_success") {
			continue
		}

		d := &dto.Metric{}
		if err := m.Write(d); err != nil {
			t.Fatalf("could not write metric: %v", err)
		}

		return d.GetGauge().GetValue()
	}

	t.Fatalf("%scollect_success was not emitted", prefix)
	return 0
}

func TestCollectSuccessReportsCollectorOutcome(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want float64
	}{
		{name: "success", err: nil, want: 1},
		{name: "error", err: errors.New("rpc failed"), want: 0},
		// io.EOF is what a dead SSH session returns and must not be treated as
		// a successful scrape.
		{name: "eof", err: io.EOF, want: 0},
		{name: "wrapped eof", err: errors.New("read: " + io.EOF.Error()), want: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := collectSuccessValue(t, tc.err); got != tc.want {
				t.Errorf("%scollect_success = %v, want %v", prefix, got, tc.want)
			}
		})
	}
}
