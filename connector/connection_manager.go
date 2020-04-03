package connector

import (
	"net"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/common/log"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
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

// SSHConnectionManager manages SSH connections to different devices
type SSHConnectionManager struct {
	connections       map[string]*SSHConnection
	reconnectInterval time.Duration
	keepAliveInterval time.Duration
	keepAliveTimeout  time.Duration
	mu                sync.Mutex
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
	client, conn, err := m.connectToDevice(device)
	if err != nil {
		return nil, err
	}

	c := &SSHConnection{
		conn:   conn,
		client: client,
		device: device,
		done:   make(chan struct{}),
	}
	go m.keepAlive(c)

	m.connections[device.Host] = c

	return c, nil
}

func (m *SSHConnectionManager) connectToDevice(device *Device) (*ssh.Client, net.Conn, error) {
	cfg := &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeoutInSeconds * time.Second,
	}

	device.Auth(cfg)

	host := m.tcpAddressForHost(device.Host)

	conn, err := net.DialTimeout("tcp", host, cfg.Timeout)
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not open tcp connection")
	}

	c, chans, reqs, err := ssh.NewClientConn(conn, host, cfg)
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not connect to device")
	}

	return ssh.NewClient(c, chans, reqs), conn, nil
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

func (m *SSHConnectionManager) keepAlive(connection *SSHConnection) {
	for {
		select {
		case <-time.After(m.keepAliveInterval):
			log.Debugf("Sending keepalive for ")
			connection.conn.SetDeadline(time.Now().Add(m.keepAliveTimeout))
			_, _, err := connection.client.SendRequest("keepalive@golang.org", true, nil)
			if err != nil {
				log.Infof("Lost connection to %s (%v). Trying to reconnect...", connection.device, err)
				connection.terminate()
				m.reconnect(connection)
			}
		case <-connection.done:
			return
		}
	}
}

func (m *SSHConnectionManager) reconnect(connection *SSHConnection) {
	for {
		client, conn, err := m.connectToDevice(connection.device)
		if err == nil {
			connection.client = client
			connection.conn = conn
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
