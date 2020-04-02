package connector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPAddressForHost(t *testing.T) {
	tests := []struct {
		name     string
		host     string
		expected string
	}{
		{
			name:     "hostname without port",
			host:     "test.routing.rocks",
			expected: "test.routing.rocks:22",
		},
		{
			name:     "hostname with port",
			host:     "test.routing.rocks:22",
			expected: "test.routing.rocks:22",
		},
		{
			name:     "IPv4 without port",
			host:     "127.0.0.1",
			expected: "127.0.0.1:22",
		},
		{
			name:     "IPv4 with port",
			host:     "127.0.0.1:22",
			expected: "127.0.0.1:22",
		},
		{
			name:     "IPv6 without port",
			host:     "[2001:678:1e0:f00::1]",
			expected: "[2001:678:1e0:f00::1]:22",
		},
		{
			name:     "IPv6 with port",
			host:     "[2001:678:1e0:f00::1]:22",
			expected: "[2001:678:1e0:f00::1]:22",
		},
	}

	t.Parallel()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := NewConnectionManager()
			assert.Equal(t, test.expected, m.tcpAddressForHost(test.host))
		})
	}
}
