// SPDX-License-Identifier: MIT

package rpc

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"

	"github.com/czerwonk/junos_exporter/pkg/connector"
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

// RunCommandAndParse runs a command on JunOS and unmarshals the XML result
func (c *Client) RunCommandAndParse(cmd string, obj interface{}) error {
	return c.RunCommandAndParseWithParser(cmd, func(b []byte) error {
		return xml.Unmarshal(b, obj)
	})
}

func (c *Client) RunCommandAndParseCustom(cmd string, obj interface{}) error {
	return c.RunCommandAndParseWithParserCustom(cmd, func(b []byte) error {
		return xml.Unmarshal(b, obj)
	})
}

// RunCommandAndParseWithParser runs a command on JunOS and unmarshals the XML result using the specified parser function
func (c *Client) RunCommandAndParseWithParser(cmd string, parser Parser) error {
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

	err = parser(b)
	return err
}

// This custom parser takes care of the case where we have an empty <outbound-secure-channel> in macsec feature
func (c *Client) RunCommandAndParseWithParserCustom(cmd string, parser Parser) error {
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

	if strings.Contains(cmd, "show security macsec connections") {
		if strings.Contains(string(b), "<macsec-connection-information>") &&
			!strings.Contains(string(b), "<outbound-secure-channel>") {
			// Find the position where we need to insert the empty outbound-secure-channel element
			// Typically, this would be before the closing tag of macsec-connection-information or before inbound-secure-channel
			insertPos := -1
			if strings.Contains(string(b), "<inbound-secure-channel>") {
				// Insert before inbound-secure-channel
				insertPos = strings.Index(string(b), "<inbound-secure-channel>")
			} else if strings.Contains(string(b), "</macsec-connection-information>") {
				// Insert before closing macsec-connection-information tag
				insertPos = strings.Index(string(b), "</macsec-connection-information>")
			}
			if insertPos > 0 {
				emptyOutboundElement := "<outbound-secure-channel/>"
				// Insert the empty element at the determined position
				modifiedXML := string(b[:insertPos]) + emptyOutboundElement + string(b[insertPos:])
				b = []byte(modifiedXML)
			}
		}
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
