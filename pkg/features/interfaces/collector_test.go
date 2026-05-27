// SPDX-License-Identifier: MIT

package interfaces

import (
	"context"
	"testing"

	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
)

type mockClient struct {
	lastCmd string
}

func (m *mockClient) RunCommandAndParse(cmd string, obj interface{}) error {
	m.lastCmd = cmd
	return nil
}

func (m *mockClient) RunCommandAndParseWithParser(cmd string, parser rpc.Parser) error {
	m.lastCmd = cmd
	return nil
}

func (m *mockClient) IsSatelliteEnabled() bool {
	return false
}

func (m *mockClient) IsScrapingLicenseEnabled() bool {
	return false
}

func (m *mockClient) Device() *connector.Device {
	return &connector.Device{Host: "test-device"}
}

func (m *mockClient) Context() context.Context {
	return context.TODO()
}

func TestInterfaceCollectorCommand(t *testing.T) {
	tests := []struct {
		name              string
		extensiveArgument string
		expectedCmd       string
	}{
		{
			name:              "default (empty)",
			extensiveArgument: "",
			expectedCmd:       "show interfaces extensive",
		},
		{
			name:              "custom argument",
			extensiveArgument: "[!(d)][!(i)]*",
			expectedCmd:       "show interfaces extensive [!(d)][!(i)]*",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			col := NewCollector(nil, tt.extensiveArgument).(*interfaceCollector)
			client := &mockClient{}
			_, _ = col.interfaceStats(client)
			if client.lastCmd != tt.expectedCmd {
				t.Errorf("expected command %q, got %q", tt.expectedCmd, client.lastCmd)
			}
		})
	}
}
