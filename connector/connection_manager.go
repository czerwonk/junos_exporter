package connector

import (
	"io"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/common/log"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

const timeoutInSeconds = 5

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
	config            *ssh.ClientConfig
	connections       map[string]*SSHConnection
	reconnectInterval time.Duration
	keepAliveInterval time.Duration
	keepAliveTimeout  time.Duration
	mu                sync.Mutex
}

// AuthMethod ithe method to use to authenticate agaist the device
type AuthMethod func(*SSHConnectionManager)

// AuthByPassword uses password authentication
func AuthByPassword(password string) AuthMethod {
	return func(m *SSHConnectionManager) {
		m.config.Auth = append(m.config.Auth, ssh.Password(password))
	}
}

// AuthByKey uses public key authentication
func AuthByKey(key io.Reader) (AuthMethod, error) {
	pk, err := loadPrivateKey(key)
	if err != nil {
		return nil, err
	}

	return func(m *SSHConnectionManager) {
		m.config.Auth = append(m.config.Auth, pk)
	}, nil
}

// NewConnectionManager creates a new connection manager
func NewConnectionManager(user string, auth AuthMethod, opts ...Option) *SSHConnectionManager {
	cfg := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeoutInSeconds * time.Second,
	}

	m := &SSHConnectionManager{
		config:            cfg,
		connections:       make(map[string]*SSHConnection),
		reconnectInterval: 30 * time.Second,
		keepAliveInterval: 10 * time.Second,
		keepAliveTimeout:  15 * time.Second,
	}

	auth(m)

	for _, opt := range opts {
		opt(m)
	}

	return m
}

// Connect connects to a device or returns an long living connection
func (m *SSHConnectionManager) Connect(host string) (*SSHConnection, error) {
	if !strings.Contains(host, ":") {
		host = host + ":22"
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if connection, found := m.connections[host]; found {
		if !connection.isConnected() {
			return nil, errors.New("not connected")
		}

		return connection, nil
	}

	return m.connect(host)
}

func (m *SSHConnectionManager) connect(host string) (*SSHConnection, error) {
	client, conn, err := m.connectToServer(host)
	if err != nil {
		return nil, err
	}

	c := &SSHConnection{
		conn:   conn,
		client: client,
		host:   host,
		done:   make(chan struct{}),
	}
	go m.keepAlive(c)

	m.connections[host] = c

	return c, nil
}

func (m *SSHConnectionManager) connectToServer(host string) (*ssh.Client, net.Conn, error) {
	conn, err := net.DialTimeout("tcp", host, m.config.Timeout)
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not open tcp connection")
	}

	c, chans, reqs, err := ssh.NewClientConn(conn, host, m.config)
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not connect to device")
	}

	return ssh.NewClient(c, chans, reqs), conn, nil
}

func (m *SSHConnectionManager) keepAlive(connection *SSHConnection) {
	for {
		select {
		case <-time.After(m.keepAliveInterval):
			log.Debugf("Sending keepalive for ")
			connection.conn.SetDeadline(time.Now().Add(m.keepAliveTimeout))
			_, _, err := connection.client.SendRequest("keepalive@golang.org", true, nil)
			if err != nil {
				log.Infof("Lost connection to %s (%v). Trying to reconnect...", connection.Host(), err)
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
		client, conn, err := m.connectToServer(connection.Host())
		if err == nil {
			connection.client = client
			connection.conn = conn
			return
		}

		log.Infof("Reconnect to %s failed: %v", connection.Host(), err)
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
