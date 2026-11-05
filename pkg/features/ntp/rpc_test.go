// SPDX-License-Identifier: MIT

package ntp

import (
	"context"
	"encoding/xml"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/czerwonk/junos_exporter/pkg/connector"
	"github.com/czerwonk/junos_exporter/pkg/rpc"
)

const textReply = `
<rpc-reply xmlns:junos="http://xml.juniper.net/junos/18.4R3/junos">
    <output>
        status=00f4 leap_none, sync_unspec, 15 events, event_peer/strat_chg,
        version="ntpd 4.2.0-a Wed Mar 18 11:34:58  2020 (1)",
        processor="powerpc", system="JUNOS18.4R3-S1.3", leap=00, stratum=4,
        precision=-18, rootdelay=13.569, rootdispersion=3231.830, peer=0,
        refid=192.0.2.13,
        reftime=ecebca0f.815513fb  Tue, Dec 16 2025 12:07:43.505, poll=4,
        clock=ecee831b.508a968f  Thu, Dec 18 2025 13:41:47.314, state=4,
        offset=1.847, frequency=-1.630, jitter=0.455, stability=0.107
    </output>
    <cli>
        <banner>{master:0}</banner>
    </cli>
</rpc-reply>`

const structuredReply = `
<rpc-reply xmlns:junos="http://xml.juniper.net/junos/23.4R2-S7.4/junos">
    <ntp-status>
        <associd>0</associd>
        <statusinfo>0615 leap_none, sync_ntp, 1 event, clock_sync</statusinfo>
        <version>"ntpd 4.2.8p15-a Thu Jan  8 19:18:33  2026 (1)"</version>
        <processor>"amd64"</processor>
        <system>"FreeBSDJNPR-12.1-20250723.90ae8a5_buil"</system>
        <leap>00</leap>
        <stratum>2</stratum>
        <precision>-22</precision>
        <rootdelay>0.745</rootdelay>
        <rootdisp>45.486</rootdisp>
        <refid>192.0.2.10</refid>
        <reftime>ee973723.e5fc6f17  Thu, Nov  5 2026 17:10:59.898</reftime>
        <clock>ee973eec.bf24bb2c  Thu, Nov  5 2026 17:44:12.746</clock>
        <peer>11362</peer>
        <tc>10</tc>
        <mintc>3</mintc>
        <offset>+0.118471</offset>
        <frequency>+22.018</frequency>
        <sys-jitter>0.007207</sys-jitter>
        <clk-jitter>0.030</clk-jitter>
        <clk-wander>0.006</clk-wander>
    </ntp-status>
    <cli>
        <banner></banner>
    </cli>
</rpc-reply>`

func TestMetricsFromTextOutput(t *testing.T) {
	var reply rpcReply
	require.NoError(t, xml.Unmarshal([]byte(textReply), &reply))

	m := reply.metrics()

	assert.Equal(t, "4", m["stratum"])
	assert.Equal(t, "192.0.2.13", m["refid"])
	assert.Equal(t, "1.847", m["offset"])
	assert.Equal(t, "00", m["leap"])
}

func TestMetricsFromStructuredOutput(t *testing.T) {
	var reply rpcReply
	require.NoError(t, xml.Unmarshal([]byte(structuredReply), &reply))

	m := reply.metrics()

	assert.Equal(t, "0", m["associd"])
	assert.Equal(t, "00", m["leap"])
	assert.Equal(t, "2", m["stratum"])
	assert.Equal(t, "-22", m["precision"])
	assert.Equal(t, "0.745", m["rootdelay"])
	assert.Equal(t, "192.0.2.10", m["refid"])
	assert.Equal(t, "10", m["tc"])
	assert.Equal(t, "+0.118471", m["offset"])
	assert.Equal(t, "0.007207", m["sys_jitter"])
	assert.Equal(t, "0.030", m["clk_jitter"])
}

func TestMetricsEmptyReply(t *testing.T) {
	var reply rpcReply
	require.NoError(t, xml.Unmarshal([]byte(`<rpc-reply><cli><banner></banner></cli></rpc-reply>`), &reply))

	assert.Empty(t, reply.metrics())
}

// replyClient answers every RPC with a canned XML document.
type replyClient struct {
	reply string
}

func (c *replyClient) RunCommandAndParse(_ string, obj any) error {
	return xml.Unmarshal([]byte(c.reply), obj)
}

func (c *replyClient) RunCommandAndParseWithParser(_ string, _ rpc.Parser) error { return nil }

func (c *replyClient) IsSatelliteEnabled() bool { return false }

func (c *replyClient) IsScrapingLicenseEnabled() bool { return false }

func (c *replyClient) Device() *connector.Device { return &connector.Device{Host: "test-device"} }

func (c *replyClient) Context() context.Context { return context.Background() }

func collect(t *testing.T, reply string) (map[string]float64, error) {
	t.Helper()

	ch := make(chan prometheus.Metric, 16)
	err := NewCollector().Collect(&replyClient{reply: reply}, ch, []string{"test-device"})
	close(ch)

	values := make(map[string]float64)
	for m := range ch {
		var pb dto.Metric
		require.NoError(t, m.Write(&pb))

		for _, name := range []string{"stratum", "offset", "system_jitter", "clock_jitter", "root_delay", "leap", "precision", "poll_interval"} {
			if strings.Contains(m.Desc().String(), `"`+prefix+name+`"`) {
				values[name] = pb.GetGauge().GetValue()
			}
		}
	}
	return values, err
}

func TestCollectStructuredOutput(t *testing.T) {
	values, err := collect(t, structuredReply)
	require.NoError(t, err)

	assert.Equal(t, map[string]float64{
		"stratum":       2,
		"offset":        0.118471,
		"system_jitter": 0.007207,
		"clock_jitter":  0.030,
		"root_delay":    0.745,
		"leap":          0,
		"precision":     -22,
		"poll_interval": 1024,
	}, values)
}

func TestCollectTextOutput(t *testing.T) {
	values, err := collect(t, textReply)
	require.NoError(t, err)

	assert.Equal(t, 4.0, values["stratum"])
	assert.Equal(t, 1.847, values["offset"])
	assert.Equal(t, 13.569, values["root_delay"])
}

func TestCollectWithoutNTPData(t *testing.T) {
	_, err := collect(t, `<rpc-reply><cli><banner></banner></cli></rpc-reply>`)

	assert.Error(t, err)
}
