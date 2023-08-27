// SPDX-License-Identifier: MIT

package rpc

import (
	"encoding/xml"
	"fmt"
	"github.com/czerwonk/junos_exporter/pkg/connector"
	"log"
)

// Parser parses XML of RPC-Output
type Parser func([]byte) error

type ClientOption func(*Client)

func WithDebug() ClientOption {
	return func(cl *Client) {
		cl.debug = true
	}
}

func WithSatellite() ClientOption {
	return func(cl *Client) {
		cl.satellite = true
	}
}

func WithLicenseInformation() ClientOption {
	return func(cl *Client) {
		cl.license = true
	}
}

// Client sends commands to JunOS and parses results
type Client struct {
	conn      *connector.SSHConnection
	debug     bool
	satellite bool
	license   bool
}

// NewClient creates a new client to connect to
func NewClient(ssh *connector.SSHConnection, opts ...ClientOption) *Client {
	cl := &Client{conn: ssh}

	for _, opt := range opts {
		opt(cl)
	}

	return cl
}

func UnpackRpcReply(data []byte) ([]byte, error) {
	var reply *rpcReply
	var err = xml.Unmarshal(data, &reply)

	if err != nil {
		return nil, &RpcReplyXmlParserError{
			RawResponse: data,
			Err:         err,
		}
	}

	return reply.Body, nil
}

// RunCommandAndParse runs a command on JunOS and unmarshals the XML result
func (c *Client) RunCommandAndParse(cmd string, obj interface{}) error {
	return c.RunCommandAndParseWithParser(cmd, func(b []byte) error {
		return xml.Unmarshal(b, obj)
	})
}

// RunCommandAndParseWithParser runs a command on JunOS and unmarshals the XML result using the specified parser function
func (c *Client) RunCommandAndParseWithParser(cmd string, parser Parser) error {
	if c.debug {
		log.Printf("Running command on %s: %s\n", c.conn.Host(), cmd)
	}

	var err error
	var b []byte

	b, err = c.conn.RunCommand(fmt.Sprintf("%s | display xml", cmd))

	if err != nil {
		return err
	}

	if c.debug {
		log.Printf("Output for %s: %s\n", c.conn.Host(), string(b))
	}

	b, err = UnpackRpcReply(b)

	if err != nil {
		return err
	}

	err = parser(b)
	return err
}

// Device returns device information for the connected device
func (c *Client) Device() *connector.Device {
	return c.conn.Device()
}

// IsSatelliteEnabled returns if sattelite features are enabled on the device
func (c *Client) IsSatelliteEnabled() bool {
	return c.satellite
}

func (c *Client) IsScrapingLicenseEnabled() bool {
	return c.license
}
