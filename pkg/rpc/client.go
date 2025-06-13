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
	empty_element := "<outbound-secure-channel>\n            <sci></sci>\n            <outgoing-packet-number></outgoing-packet-number>\n            <outbound-secure-association>\n                <association-number></association-number>\n                <association-number-status></association-number-status>\n                <create-time junos:seconds=\"\"></create-time>\n            </outbound-secure-association>\n        </outbound-secure-channel>"
	//slice of interfaces
	//!Next line may not return what is expected!
	str_int := returnTextInBetween(string(b), "<interface-name>", "<interface-name>")
	for i, _ := range str_int {
		fmt.Println(str_int[i])
	}
	empty_outbound_connections := make([]bool, len(str_int))
	for i, _ := range str_int {
		//fmt.Println(str_int[i])
		empty_outbound_connections[i] = strings.Contains(string(str_int[i]), "<outbound-secure-channel>")
	}
	for i, _ := range empty_outbound_connections {
		fmt.Println(empty_outbound_connections[i])
	}
	for i, _ := range empty_outbound_connections {
		if empty_outbound_connections[i] == true {
			//interface_to_adust := returnTextInBetween(str_int[i], "<interface-name>", "</interface-name>")
			//fmt.Println(interface_to_adust)
			//b = []byte(addStringToXML(string(b), interface_to_adust[0], empty_element))
			b = []byte(addStringToXML(string(b), str_int[i], empty_element))
		}
	}
	//fmt.Println(string(b))
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

// returnTextInBetween extracts all substrings between specified start and end strings within the given input string.
func returnTextInBetween(str string, start string, end string) (result []string) {
	// Initialize an empty slice to store results
	result = []string{}

	// Current position in the string
	currentPos := 0

	for {
		// Find the next occurrence of the start string
		startIndex := strings.Index(str[currentPos:], start)
		if startIndex == -1 {
			// No more occurrences of the start string
			break
		}

		// Adjust the start index to the absolute position
		startIndex += currentPos

		// Move position past the start string
		startPos := startIndex + len(start)

		// Find the end string after the start string
		endIndex := strings.Index(str[startPos:], end)
		if endIndex == -1 {
			// No corresponding end string
			break
		}

		// Adjust the end index to the absolute position
		endIndex += startPos

		// Extract the text between start and end
		extracted := str[startPos:endIndex]

		// Add to results
		result = append(result, extracted)

		// Move current position past the end string for next iteration
		currentPos = endIndex + len(end)
	}

	return result
}

// The function inserts a string () into an XML document () at a specific location - specifically,
// right before the first occurrence of that appears after a specified marker string (). `addStringToXML“stringToWrite“str“<inbound-secure-channel>“start`
func addStringToXML(str string, start string, stringToWrite string) string {
	// Find the position of the first occurrence of 'start'
	startPos := strings.Index(str, start)
	fmt.Println("printing start")
	fmt.Println(start)
	if startPos == -1 {
		// 'start' not found, return original string
		return str
	}

	// Find the position of '<inbound-secure-channel>' after 'start'
	searchFrom := startPos + len(start)
	inboundPos := strings.Index(str[searchFrom:], "<inbound-secure-channel>")
	if inboundPos == -1 {
		// '<inbound-secure-channel>' not found after 'start', return original string
		return str
	}

	// Adjust to get the absolute position in the original string
	inboundPos += searchFrom

	// Insert stringToWrite just before '<inbound-secure-channel>'
	return str[:inboundPos] + stringToWrite + str[inboundPos:]
}
