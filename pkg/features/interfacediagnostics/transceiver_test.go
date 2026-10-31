// SPDX-License-Identifier: MIT

package interfacediagnostics

import (
	"context"
	"encoding/xml"
	"testing"

	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// picPortsClient answers "show chassis pic ..." with a fixed reply.
type picPortsClient struct {
	reply string
}

func (c *picPortsClient) RunCommandAndParse(cmd string, obj any) error {
	return xml.Unmarshal([]byte(c.reply), obj)
}

func (c *picPortsClient) RunCommandAndParseWithParser(cmd string, parser rpc.Parser) error {
	return parser([]byte(c.reply))
}

func (c *picPortsClient) IsSatelliteEnabled() bool { return false }

func (c *picPortsClient) IsScrapingLicenseEnabled() bool { return false }

func (c *picPortsClient) Device() *connector.Device {
	return &connector.Device{Host: "test-device"}
}

func (c *picPortsClient) Context() context.Context { return context.Background() }

func chassisWithTransceivers(names ...string) chassisHardware {
	xcvrs := make([]chassisSubSubModule, 0, len(names))
	for _, n := range names {
		xcvrs = append(xcvrs, chassisSubSubModule{Name: n, SerialNumber: "SN-" + n})
	}

	hw := chassisHardware{}
	hw.ChassisInventory.Chassis.ChassisModule = []chassisModule{
		{
			Name: "FPC 0",
			ChassisSubModule: []chassisSubModule{
				{Name: "PIC 1", ChassisSubSubModule: xcvrs},
			},
		},
	}

	return hw
}

// The PIC port list only has the populated ports, and it does not have to
// list the same transceivers as the chassis inventory. Each transceiver must
// get the port details for its own port number, and a missing port must not
// panic.
func TestTransceiverInfoMatchesPicPortsByNumber(t *testing.T) {
	cl := &picPortsClient{reply: `<rpc-reply>
  <fpc-information>
    <fpc>
      <pic-detail>
        <slot>0</slot>
        <pic-slot>1</pic-slot>
        <port-information>
          <port>
            <port-number>2</port-number>
            <fiber-mode>SM</fiber-mode>
            <sfp-vendor-name>VENDOR-2</sfp-vendor-name>
            <sfp-vendor-pno>PNO-2</sfp-vendor-pno>
            <wavelength>1310 nm</wavelength>
          </port>
          <port>
            <port-number>3</port-number>
            <fiber-mode>MM</fiber-mode>
            <sfp-vendor-name>VENDOR-3</sfp-vendor-name>
            <sfp-vendor-pno>PNO-3</sfp-vendor-pno>
            <wavelength>850 nm</wavelength>
          </port>
        </port-information>
      </pic-detail>
    </fpc>
  </fpc-information>
</rpc-reply>`}

	c := &interfaceDiagnosticsCollector{}
	got, err := c.transceiverInfoFromRPCResult(cl, chassisWithTransceivers("Xcvr 0", "Xcvr 2", "Xcvr 3"))
	require.NoError(t, err)
	require.Len(t, got, 3)

	byName := make(map[string]*transceiverInformation, len(got))
	for _, tr := range got {
		byName[tr.Name] = tr
	}

	require.Contains(t, byName, "0/1/0")
	assert.Equal(t, "SN-Xcvr 0", byName["0/1/0"].ChassisHardwareInfo.SerialNumber)
	assert.Empty(t, byName["0/1/0"].PicPort.SFPVendorName, "port 0 is not in the PIC port list")

	require.Contains(t, byName, "0/1/2")
	assert.Equal(t, "VENDOR-2", byName["0/1/2"].PicPort.SFPVendorName)
	assert.Equal(t, "1310 nm", byName["0/1/2"].PicPort.Wavelength)

	require.Contains(t, byName, "0/1/3")
	assert.Equal(t, "VENDOR-3", byName["0/1/3"].PicPort.SFPVendorName)
	assert.Equal(t, "MM", byName["0/1/3"].PicPort.FiberMode)
}
