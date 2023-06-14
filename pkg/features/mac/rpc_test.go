// SPDX-License-Identifier: MIT

package mac

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/czerwonk/junos_exporter/pkg/rpc"
)

func TestParseOutput(t *testing.T) {
	body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/ZZZ/junos">
    <ethernet-switching-table-information junos:style="summary">
        <ethernet-switching-table junos:style="summary">
            <mac-table-entry junos:style="summary">
                <mac-table-total-count>37</mac-table-total-count>
                <mac-table-dot1x-count>5</mac-table-dot1x-count>
                <mac-table-recieve-count>1</mac-table-recieve-count>
                <mac-table-dynamic-count>28</mac-table-dynamic-count>
                <mac-table-flood-count>3</mac-table-flood-count>
            </mac-table-entry>
        </ethernet-switching-table>
    </ethernet-switching-table-information>
    <cli>
        <banner>{master:1}</banner>
    </cli>
</rpc-reply>`

	var b, err = rpc.UnpackRpcReply([]byte(body))

	if err != nil {
		t.Fatal(err)
	}

	result := ethernetSwitchingTableInformation{}
	err = xml.Unmarshal(b, &result)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int64(37), result.Table.Entry.TotalCount, "total-count")
	assert.Equal(t, int64(5), result.Table.Entry.Dot1XCount, "dot1x-count")
	assert.Equal(t, int64(1), result.Table.Entry.ReceiveCount, "receive-count")
	assert.Equal(t, int64(28), result.Table.Entry.DynamicCount, "dynamic-count")
	assert.Equal(t, int64(3), result.Table.Entry.FloodCount, "flood-count")
}
