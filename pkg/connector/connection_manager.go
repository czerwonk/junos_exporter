// SPDX-License-Identifier: MIT

package connector

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

const timeoutInSeconds = 5
const defaultPort = "22"

// Option defines options for the manager which are applied on creation
type Option func(*SSHConnectionManager)

// WithReconnectInterval sets the reconnect interval (default 10 seconds)
func WithReconnectInterval(d time.Duration) Option {
	return func(m *SSHConnectionManager) {
		m.reconnectInterval = d
	}
}

// WithKeepAliveInterval sets the keep alive interval (default 10 seconds)
func WithKeepAliveInterval(d time.Duration) Option {
	return func(m *SSHConnectionManager) {
		m.keepAliveInterval = d
	}
}

// WithKeepAliveTimeout sets the timeout after an ssh connection to be determined dead (default 15 seconds)
func WithKeepAliveTimeout(d time.Duration) Option {
	return func(m *SSHConnectionManager) {
		m.keepAliveTimeout = d
	}
}

// WithExpiredConnectionTime sets the timeout after an unused ssh connection will not be keepalived
func WithExpiredConnectionTimeout(d time.Duration) Option {
	return func(m *SSHConnectionManager) {
		m.expiredConnectionTimeout = d
	}
}

// SSHConnectionManager manages SSH connections to different devices
type SSHConnectionManager struct {
	connections              map[string]*SSHConnection
	connectionsMu            sync.RWMutex
	reconnectInterval        time.Duration
	keepAliveInterval        time.Duration
	keepAliveTimeout         time.Duration
	expiredConnectionTimeout time.Duration
}

// NewConnectionManager creates a new connection manager
func NewConnectionManager(opts ...Option) *SSHConnectionManager {
	m := &SSHConnectionManager{
		connections:       make(map[string]*SSHConnection),
		reconnectInterval: 30 * time.Second,
		keepAliveInterval: 10 * time.Second,
		keepAliveTimeout:  15 * time.Second,
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

// GetSSHConnection gets a cached SSHConnection or creates a fresh one, if necessary
func (m *SSHConnectionManager) GetSSHConnection(device *Device) (*SSHConnection, error) {
	connection := m.getExistingConnection(device)
	if connection != nil {
		log.Infof("Re-using existing connection with %s", device.Host)
		return connection, nil
	}

	return m.connect(device)
}

func (m *SSHConnectionManager) getExistingConnection(device *Device) *SSHConnection {
	m.connectionsMu.RLock()
	defer m.connectionsMu.RUnlock()

	if connection, found := m.connections[device.Host]; found {
		if connection.IsConnected() {
			return connection
		}
	}

	return nil
}

func (m *SSHConnectionManager) connect(device *Device) (*SSHConnection, error) {
	log.Infof("Creating SSH connection with %s", device.Host)
	c := NewSSHConnection(device, m.keepAliveInterval, m.keepAliveTimeout)
	err := c.Start(m.expiredConnectionTimeout)
	if err != nil {
		return nil, fmt.Errorf("unable to get new SSH connection: %w", err)
	}

	m.connectionsMu.Lock()
	defer m.connectionsMu.Unlock()

	if existingCon, exists := m.connections[device.Host]; exists && existingCon.IsConnected() {
		c.Stop(fmt.Errorf("connection conflict"))
		return existingCon, nil
	}

	m.connections[device.Host] = c
	return c, nil
}

func tcpAddressForHost(host string) string {
	colonCount := strings.Count(host, ":")

	if colonCount == 0 {
		// either legacy IP or hostname without port
		return host + ":" + defaultPort
	}

	h, p, err := net.SplitHostPort(host)
	if err == nil {
		return formatHost(h) + ":" + p
	}

	return formatHost(host) + ":" + defaultPort
}

func formatHost(host string) string {
	ip := net.ParseIP(host)

	if ip == nil || ip.To4() != nil {
		// not an IP or IPv4 address
		return host
	}

	return "[" + host + "]"
}

// CloseAll closes all TCP connections and stops keep alives
func (m *SSHConnectionManager) CloseAll() error {
	for _, c := range m.connections {
		c.Stop(fmt.Errorf("end of world"))
	}

	return nil
}
