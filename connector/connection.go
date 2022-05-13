package connector

import (
	"net"
	"sync"

	"github.com/pkg/errors"

	"github.com/Juniper/go-netconf/netconf"
)

// SSHConnection encapsulates the connection to the device
type SSHConnection struct {
	device *Device
	session *netconf.Session
	conn   net.Conn
	mu     sync.Mutex
	done   chan struct{}
}

// RunCommand runs a command against the device
func (c *SSHConnection) RunCommand(cmd string) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.session == nil {
		return nil, errors.New("not connected")
	}

	reply, err := c.session.Exec(netconf.RawMethod(cmd))
	if err != nil {
		return nil, errors.Wrap(err, "could not run command")
	}

	return []byte(reply.RawReply), nil
}

func (c *SSHConnection) isConnected() bool {
	return c.conn != nil
}

func (c *SSHConnection) terminate() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.conn.Close()

	c.session = nil
	c.conn = nil
}

func (c *SSHConnection) close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.session != nil {
		c.session.Close()
	}

	c.done <- struct{}{}
	c.conn = nil
	c.session = nil
}

// Host returns the hostname of the connected device
func (c *SSHConnection) Host() string {
	return c.device.Host
}

// Device returns the device information of the connected device
func (c *SSHConnection) Device() *Device {
	return c.device
}
