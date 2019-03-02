package rpc

import (
	"encoding/xml"
	"fmt"

	"log"

	"github.com/czerwonk/junos_exporter/connector"
)

// Client sends commands to JunOS and parses results
type Client struct {
	conn  *connector.SSHConnection
	debug bool
}

// NewClient creates a new client to connect to
func NewClient(ssh *connector.SSHConnection, debug bool) *Client {
	rpc := &Client{conn: ssh, debug: debug}

	return rpc
}

// RunCommandAndParse runs a command on JunOS and unmarshals the XML result
func (c *Client) RunCommandAndParse(cmd string, obj interface{}) error {
	if c.debug {
		log.Printf("Running command on %s: %s\n", c.conn.Host(), cmd)
	}

	b, err := c.conn.RunCommand(fmt.Sprintf("%s | display xml", cmd))
	if err != nil {
		return err
	}

	if c.debug {
		log.Printf("Output for %s: %s\n", c.conn.Host(), string(b))
	}

	err = xml.Unmarshal(b, obj)
	return err
}
