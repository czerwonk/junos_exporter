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

func (m *mockClient) RunCommandAndParse(cmd string, obj any) error {
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
		name               string
		interfaceNameRegex string
		expectedCmd        string
	}{
		{
			name:               "default (empty)",
			interfaceNameRegex: "",
			expectedCmd:        "show interfaces extensive",
		},
		{
			name:               "custom argument",
			interfaceNameRegex: "[!(d)][!(i)]*",
			expectedCmd:        "show interfaces extensive [!(d)][!(i)]*",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			col := NewCollector(nil, tt.interfaceNameRegex).(*interfaceCollector)
			client := new(mockClient)
			_, _ = col.interfaceStats(client)
			if client.lastCmd != tt.expectedCmd {
				t.Errorf("expected command %q, got %q", tt.expectedCmd, client.lastCmd)
			}
		})
	}
}

func TestInterfaceSpeed(t *testing.T) {
	tests := []struct {
		name       string
		speed      string
		ifSpeedCfg string
		ifName     string
		expected   float64
	}{
		{
			name:       "configured 1G overrides 10G hardware speed",
			speed:      "10Gbps",
			ifSpeedCfg: "1G",
			ifName:     "xe-0/1/5",
			expected:   1e9,
		},
		{
			name:       "configured 100M",
			speed:      "1Gbps",
			ifSpeedCfg: "100M",
			ifName:     "ge-0/0/0",
			expected:   100e6,
		},
		{
			name:       "configured lowercase 1g",
			speed:      "10Gbps",
			ifSpeedCfg: "1g",
			ifName:     "xe-0/1/5",
			expected:   1e9,
		},
		{
			name:       "configured 40G",
			speed:      "40Gbps",
			ifSpeedCfg: "40G",
			ifName:     "et-0/0/0",
			expected:   40e9,
		},
		{
			name:       "configured fractional 2.5G",
			speed:      "10Gbps",
			ifSpeedCfg: "2.5G",
			ifName:     "ge-0/0/0",
			expected:   2.5e9,
		},
		{
			name:       "configured Auto falls back to hardware speed",
			speed:      "10Gbps",
			ifSpeedCfg: "Auto",
			ifName:     "xe-0/1/5",
			expected:   10e9,
		},
		{
			name:       "configured AUTO (uppercase) falls back to hardware speed",
			speed:      "10Gbps",
			ifSpeedCfg: "AUTO",
			ifName:     "xe-0/1/5",
			expected:   10e9,
		},
		{
			name:       "no configured speed uses hardware speed",
			speed:      "1Gbps",
			ifSpeedCfg: "",
			ifName:     "ge-0/0/0",
			expected:   1e9,
		},
		{
			name:       "hardware speed in mbps",
			speed:      "1000mbps",
			ifSpeedCfg: "",
			ifName:     "ge-0/0/0",
			expected:   1e9,
		},
		{
			name:       "hardware speed with uppercase Mbps",
			speed:      "100Mbps",
			ifSpeedCfg: "",
			ifName:     "ge-0/0/0",
			expected:   100e6,
		},
		{
			name:       "hardware speed with uppercase GBPS",
			speed:      "10GBPS",
			ifSpeedCfg: "",
			ifName:     "xe-0/0/0",
			expected:   10e9,
		},
		{
			name:       "Auto on a non ge/xe interface yields 0",
			speed:      "Auto",
			ifSpeedCfg: "",
			ifName:     "et-0/0/0",
			expected:   0,
		},
		{
			name:       "Unspecified on a ge interface yields 1G",
			speed:      "Unspecified",
			ifSpeedCfg: "",
			ifName:     "ge-0/0/0",
			expected:   1e9,
		},
		{
			name:       "Unlimited yields 0",
			speed:      "Unlimited",
			ifSpeedCfg: "",
			ifName:     "lo0",
			expected:   0,
		},
		{
			name:       "configured 400G",
			speed:      "100Gbps",
			ifSpeedCfg: "400G",
			ifName:     "et-0/0/0",
			expected:   4e11,
		},
		{
			name:       "unrecognized value yields 0",
			speed:      "bogus",
			ifSpeedCfg: "",
			ifName:     "ge-0/0/0",
			expected:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &interfaceStats{
				IsPhysical: true,
				Name:       tt.ifName,
				Speed:      tt.speed,
				IfSpeedCfg: tt.ifSpeedCfg,
			}

			if got := interfaceSpeedBPS(s); got != tt.expected {
				t.Errorf("interfaceSpeedBPS() = %v, want %v", got, tt.expected)
			}
		})
	}
}
