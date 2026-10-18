// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"sync"
	"testing"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"

	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
)

// spanRecorder installs the recording tracer provider once per test binary and
// hands each test an empty recorder. The package level tracer resolves the global
// provider on first use and keeps it, so a provider per test would not be seen.
var (
	installRecorder sync.Once
	recorder        = tracetest.NewSpanRecorder()
)

func spanRecorder(t *testing.T) *tracetest.SpanRecorder {
	t.Helper()

	installRecorder.Do(func() {
		otel.SetTracerProvider(sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(recorder)))
	})
	recorder.Reset()

	return recorder
}

// disconnectedAdapter returns an adapter whose client has no SSH client, so every
// command fails without touching the network.
func disconnectedAdapter() *clientTracingAdapter {
	conn := connector.NewSSHConnection(&connector.Device{Host: "device.example"}, time.Second, time.Second)

	return &clientTracingAdapter{
		cl:  rpc.NewClient(conn),
		ctx: context.Background(),
	}
}

func TestRunCommandAndParseIsTraced(t *testing.T) {
	sr := spanRecorder(t)
	cta := disconnectedAdapter()

	err := cta.RunCommandAndParse("show version", &struct{}{})
	if err == nil {
		t.Fatal("expected an error from a disconnected client")
	}

	spans := sr.Ended()
	if len(spans) != 1 {
		t.Fatalf("recorded %d spans, want 1", len(spans))
	}

	span := spans[0]
	if span.Name() != "RunCommandAndParse" {
		t.Errorf("span name = %q, want %q", span.Name(), "RunCommandAndParse")
	}

	var cmd string
	for _, a := range span.Attributes() {
		if a.Key == "command" {
			cmd = a.Value.AsString()
		}
	}
	if cmd != "show version" {
		t.Errorf("command attribute = %q, want %q", cmd, "show version")
	}

	if span.Status().Code != codes.Error {
		t.Errorf("status = %v, want %v", span.Status().Code, codes.Error)
	}

	if len(span.Events()) == 0 {
		t.Error("expected the error to be recorded on the span")
	}
}

func TestRunCommandAndParseWithParserIsTraced(t *testing.T) {
	sr := spanRecorder(t)
	cta := disconnectedAdapter()

	err := cta.RunCommandAndParseWithParser("show interfaces", func([]byte) error { return nil })
	if err == nil {
		t.Fatal("expected an error from a disconnected client")
	}

	spans := sr.Ended()
	if len(spans) != 1 {
		t.Fatalf("recorded %d spans, want 1", len(spans))
	}

	if got := spans[0].Name(); got != "RunCommandAndParseWithParser" {
		t.Errorf("span name = %q, want %q", got, "RunCommandAndParseWithParser")
	}
}
