package main

import (
	"context"

	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type clientTracingAdapter struct {
	cl  *rpc.Client
	ctx context.Context
}

// RunCommandAndParse implements RunCommandAndParse of the collector.Client interface
func (cta *clientTracingAdapter) RunCommandAndParse(cmd string, obj interface{}) error {
	return cta.cl.RunCommandAndParse(cmd, obj)
}

// RunCommandAndParseWithParser implements RunCommandAndParseWithParser of the collector.Client interface
func (cta *clientTracingAdapter) RunCommandAndParseWithParser(cmd string, parser rpc.Parser) error {
	_, span := tracer.Start(cta.ctx, "RunCommandAndParseWithParser", trace.WithAttributes(
		attribute.String("command", cmd),
	))
	defer span.End()

	err := cta.cl.RunCommandAndParseWithParser(cmd, parser)
	return err
}

// IsSatelliteEnabled implements IsSatelliteEnabled of the collector.Client interface
func (cta *clientTracingAdapter) IsSatelliteEnabled() bool {
	return cta.cl.IsSatelliteEnabled()
}

// Device implements Device of the collector.Client interface
func (cta *clientTracingAdapter) Device() *connector.Device {
	return cta.cl.Device()
}

// Context implements Context of the collector.Client interface
func (cta *clientTracingAdapter) Context() context.Context {
	return cta.ctx
}
