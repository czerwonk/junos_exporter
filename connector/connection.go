package connector

import (
	"bytes"
	"net"
	"sync"

	"github.com/pkg/errors"

	"golang.org/x/crypto/ssh"
)

// SSHConnection encapsulates the connection to the device
type SSHConnection struct {
	device *Device
	client *ssh.Client
	conn   net.Conn
	mu     sync.Mutex
	done   chan struct{}
}

// RunCommand runs a command against the device
func (c *SSHConnection) RunCommand(cmd string) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.client == nil {
		return nil, errors.New("not conneted")
	}

	session, err := c.client.NewSession()
	if err != nil {
		return nil, errors.Wrap(err, "could not open session")
	}
	defer session.Close()

	var b = &bytes.Buffer{}
	session.Stdout = b

	err = session.Run(cmd)
	if err != nil {
		return nil, errors.Wrap(err, "could not run command")
	}

	return b.Bytes(), nil
}

func (c *SSHConnection) isConnected() bool {
	return c.conn != nil
}

func (c *SSHConnection) terminate() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.conn.Close()

	c.client = nil
	c.conn = nil
}

func (c *SSHConnection) close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.client != nil {
		c.client.Close()
	}

	c.done <- struct{}{}
	c.conn = nil
	c.client = nil
}

// Host returns the hostname of the connected device
func (c *SSHConnection) Host() string {
	return c.device.Host
}

// Device returns the device information of the connected device
func (c *SSHConnection) Device() *Device {
	return c.device
}
