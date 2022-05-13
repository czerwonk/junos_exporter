package connector

import (
	"net"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"github.com/Juniper/go-netconf/netconf"
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

// WithKeepAliveTimeout sets the timeout after an ssh connection to be determined dead (default 15 seconds)
func WithKeepAliveTimeout(d time.Duration) Option {
	return func(m *SSHConnectionManager) {
		m.keepAliveTimeout = d
	}
}


// SSHConnectionManager manages SSH connections to different devices
type SSHConnectionManager struct {
	connections       map[string]*SSHConnection
	reconnectInterval time.Duration
	keepAliveTimeout  time.Duration

	mu                sync.Mutex
}

// NewConnectionManager creates a new connection manager
func NewConnectionManager(opts ...Option) *SSHConnectionManager {
	m := &SSHConnectionManager{
		connections:       make(map[string]*SSHConnection),
		reconnectInterval: 30 * time.Second,
		keepAliveTimeout:  20 * time.Second,
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

// Connect connects to a device or returns an long living connection
func (m *SSHConnectionManager) Connect(device *Device) (*SSHConnection, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if connection, found := m.connections[device.Host]; found {
		if !connection.isConnected() {
			return nil, errors.New("not connected")
		}

		return connection, nil
	}

	return m.connect(device)
}

func (m *SSHConnectionManager) connect(device *Device) (*SSHConnection, error) {
	session, err := m.connectToDevice(device)
	if err != nil {
		return nil, err
	}

	c := &SSHConnection{
		session: session,
		device: device,
		done:   make(chan struct{}),
	}

	m.connections[device.Host] = c

	return c, nil
}

func (m *SSHConnectionManager) connectToDevice(device *Device) (*netconf.Session, error) {
	cfg := &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeoutInSeconds * time.Second,
	}

	device.Auth(cfg)

	host := m.tcpAddressForHost(device.Host)

	session,err := netconf.DialSSHTimeout(host,cfg,m.keepAliveTimeout)
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to device")
	}

	return session, nil
}

func (m *SSHConnectionManager) tcpAddressForHost(host string) string {
	colonCount := strings.Count(host, ":")

	if colonCount == 0 {
		// either legacy IP or hostname without port
		return host + ":" + defaultPort
	}

	h, p, err := net.SplitHostPort(host)
	if err == nil {
		return m.formatHost(h) + ":" + p
	}

	return m.formatHost(host) + ":" + defaultPort
}

func (m *SSHConnectionManager) formatHost(host string) string {
	ip := net.ParseIP(host)

	if ip == nil || ip.To4() != nil {
		// not an IP or IPv4 address
		return host
	}

	return "[" + host + "]"
}


func (m *SSHConnectionManager) reconnect(connection *SSHConnection) {
	for {
		session, err := m.connectToDevice(connection.device)
		if err == nil {
			connection.session = session
			return
		}

		log.Infof("Reconnect to %s failed: %v", connection.device, err)
		time.Sleep(m.reconnectInterval)
	}
}

// Close closes all TCP connections and stop keep alives
func (m *SSHConnectionManager) Close() error {
	for _, c := range m.connections {
		c.close()
	}

	return nil
}
