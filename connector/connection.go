package connector

import (
	"net"
	"sync"

	"github.com/pkg/errors"

	"golang.org/x/crypto/ssh"
	"github.com/Juniper/go-netconf/netconf"
)

// SSHConnection encapsulates the connection to the device
type SSHConnection struct {
	device *Device
	client *ssh.Client
	conn   net.Conn
	mu     sync.Mutex
	done   chan struct{}
}

type TransportSSH struct {
    transportBasicIO
    sshClient  *ssh.Client
    sshSession *ssh.Session
}


// RunCommand runs a command against the device
func (c *SSHConnection) RunCommand(cmd string) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	var err error

	if c.client == nil {
		return nil, errors.New("not connected")
	}

	t := &TransportSSH{}
	t.sshSession, err = c.client.NewSession()
	if err != nil {
		return nil, errors.Wrap(err, "could not open session")
	}
	defer t.sshSession.Close()

	writer, err := t.sshSession.StdinPipe()
	if err != nil {
		return nil, errors.Wrap(err, "could not open session stdin")
	}

	reader, err := t.sshSession.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "could not open session stdout")
	}

	t.ReadWriteCloser = netconf.NewReadWriteCloser(reader, writer)
	t.sshSession.RequestSubsystem("netconf")
	session := netconf.NewSession(t)


	reply, err := session.Exec(netconf.RawMethod(cmd))
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
