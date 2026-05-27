// SPDX-License-Identifier: MIT

package connector

import (
	"bytes"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"golang.org/x/crypto/ssh"
)

// SSHConnection encapsulates the connection to the device
type SSHConnection struct {
	device            *Device
	sshClient         *ssh.Client
	tcpConn           net.Conn
	isConnected       bool
	mu                sync.RWMutex // protects sshClient, tcpConn and isConnected
	lastUsed          time.Time
	lastUsedMu        sync.RWMutex
	done              chan struct{}
	keepAliveInterval time.Duration
	keepAliveTimeout  time.Duration
}

func NewSSHConnection(device *Device, keepAliveInterval time.Duration, keepAliveTimeout time.Duration) *SSHConnection {
	return &SSHConnection{
		device:            device,
		keepAliveInterval: keepAliveInterval,
		keepAliveTimeout:  keepAliveTimeout,
		done:              make(chan struct{}),
	}
}

func (c *SSHConnection) Start(expiredConnectionTimeout time.Duration) error {
	err := c.connect()
	if err != nil {
		return err
	}

	go c.keepalive(expiredConnectionTimeout)
	return nil
}

func (c *SSHConnection) Stop(err error) {
	log.Infof("Stopping SSH connection with %s (reason: %v)", c.device.Host, err)

	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.isConnected {
		return
	}

	close(c.done)

	if c.sshClient != nil {
		c.sshClient.Close()
		c.sshClient = nil
	}

	if c.tcpConn != nil {
		c.tcpConn.Close()
		c.tcpConn = nil
	}

	c.isConnected = false
}

// RunCommand runs a command against the device
func (c *SSHConnection) RunCommand(cmd string) ([]byte, error) {
	c.setLastUsed(time.Now())

	sshClient := c.getSSHClient()
	if sshClient == nil {
		c.Stop(fmt.Errorf("No ssh client"))
		return nil, errors.New(fmt.Sprintf("no SSH client to %s", c.device.Host))
	}

	session, err := c.sshClient.NewSession()
	if err != nil {
		c.Stop(fmt.Errorf("SSH session failure"))
		return nil, errors.Wrapf(err, "could not open session with %s", c.device.Host)
	}
	defer session.Close()

	var b = &bytes.Buffer{}
	session.Stdout = b

	err = session.Run(cmd)
	if err != nil {
		c.Stop(fmt.Errorf("failed running command"))
		return nil, errors.Wrapf(err, "could not run command %q on %s", cmd, c.device.Host)
	}

	return b.Bytes(), nil
}

func (c *SSHConnection) keepalive(expiredConnectionTimeout time.Duration) {
	for {
		select {
		case <-time.After(c.keepAliveInterval):
			terminated := c.terminateIfLifetimeExpired(expiredConnectionTimeout)
			if terminated {
				return
			}

			_ = c.tcpConn.SetDeadline(time.Now().Add(c.keepAliveTimeout))

			ok := c.testSSHClient()
			if !ok {
				return
			}
		case <-c.done:
			return
		}
	}
}

func (c *SSHConnection) terminateIfLifetimeExpired(expiredConnectionTimeout time.Duration) bool {
	if time.Since(c.GetLastUsed()) > expiredConnectionTimeout {
		c.Stop(fmt.Errorf("lifetime expired"))
		return true
	}

	return false
}

func (c *SSHConnection) testSSHClient() bool {
	sshClient := c.getSSHClient()

	_, _, err := sshClient.SendRequest("keepalive@golang.org", true, nil)
	if err != nil {
		log.Infof("SSH keepalive request to %s failed: %v", c.device, err)
		c.Stop(fmt.Errorf("keepalive failed"))
		return false
	}

	return true
}

func (c *SSHConnection) connect() error {
	cfg := &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeoutInSeconds * time.Second,
	}

	c.device.Auth(cfg)

	host := tcpAddressForHost(c.device.Host)
	log.Infof("Establishing TCP connection with %s", host)

	tcpConn, err := net.DialTimeout("tcp", host, cfg.Timeout)
	if err != nil {
		return fmt.Errorf("could not open tcp connection: %w", err)
	}

	sshConn, chans, reqs, err := ssh.NewClientConn(tcpConn, host, cfg)
	if err != nil {
		tcpConn.Close()
		return fmt.Errorf("could not connect to device: %w", err)
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.tcpConn = tcpConn
	c.sshClient = ssh.NewClient(sshConn, chans, reqs)
	c.isConnected = true

	return nil
}

func (c *SSHConnection) setLastUsed(t time.Time) {
	c.lastUsedMu.Lock()
	defer c.lastUsedMu.Unlock()

	c.lastUsed = t
}

func (c *SSHConnection) GetLastUsed() time.Time {
	c.lastUsedMu.RLock()
	defer c.lastUsedMu.RUnlock()

	return c.lastUsed
}

func (c *SSHConnection) getSSHClient() *ssh.Client {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.sshClient
}

func (c *SSHConnection) IsConnected() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.isConnected
}

// Host returns the hostname of the connected device
func (c *SSHConnection) Host() string {
	return c.device.Host
}

// Device returns the device information of the connected device
func (c *SSHConnection) Device() *Device {
	return c.device
}
