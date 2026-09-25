// SPDX-License-Identifier: MIT

package main

import (
	"runtime"
	"strings"
	"testing"

	dto "github.com/prometheus/client_model/go"
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
