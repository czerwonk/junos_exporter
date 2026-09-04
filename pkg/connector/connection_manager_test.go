// SPDX-License-Identifier: MIT

package connector

import (
	"sync"
	"testing"
	"time"

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
			name:     "IPv6 without port and brackets",
			host:     "2001:678:1e0:f00::1",
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
			assert.Equal(t, test.expected, tcpAddressForHost(test.host))
		})
	}
}

// TestCloseAllConcurrentWithConnect exercises CloseAll against in-flight
// connection setup. CloseAll ranges m.connections without holding
// connectionsMu, while connect() writes that map under the lock.
func TestCloseAllConcurrentWithConnect(t *testing.T) {
	s := startStubSSHServer(t)

	m := NewConnectionManager(WithExpiredConnectionTimeout(time.Hour))

	done := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(done)

		for i := 0; i < 100; i++ {
			d := &Device{
				Host: s.addr,
				Auth: AuthByPassword("test", "test"),
			}
			// CloseAll marks the cached connection dead, so the next call falls
			// through to connect() and writes the map again.
			_, _ = m.GetSSHConnection(d)
		}
	}()

	// Spin for as long as connections are being established, rather than a fixed
	// count: ranging a one-entry map finishes far faster than an SSH handshake.
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			default:
				m.CloseAll()
			}
		}
	}()

	wg.Wait()
}
