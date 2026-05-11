// SPDX-License-Identifier: MIT

package virtualchassis

import (
    "encoding/xml"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestParseChassisClusterStatus(t *testing.T) {
    body := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/23.4R2-S7.4/junos">
    <virtual-chassis-information>
        <preprovisioned-virtual-chassis-information>
            <virtual-chassis-id>1234.1234.1234</virtual-chassis-id>
            <virtual-chassis-mode>Enabled</virtual-chassis-mode>
        </preprovisioned-virtual-chassis-information>
        <member-list style="single-slot">
            <member>
                <member-status>Prsnt</member-status>
                <member-id>0</member-id>
                <fpc-slot>(FPC 0)</fpc-slot>
                <member-priority>129</member-priority>
                <member-mixed-mode>N</member-mixed-mode>
                <member-route-mode>VC</member-route-mode>
                <member-role>Backup</member-role>
                <neighbor-list>
                    <neighbor>
                        <neighbor-id>1</neighbor-id>
                        <neighbor-interface>vcp-255/0/54</neighbor-interface>
                    </neighbor>
                </neighbor-list>
            </member>
            <member>
                <member-status>Prsnt</member-status>
                <member-id>1</member-id>
                <fpc-slot>(FPC 1)</fpc-slot>
                <member-priority>129</member-priority>
                <member-mixed-mode>N</member-mixed-mode>
                <member-route-mode>VC</member-route-mode>
                <member-role>Master*</member-role>
                <neighbor-list>
                    <neighbor>
                        <neighbor-id>0</neighbor-id>
                        <neighbor-interface>vcp-255/0/55</neighbor-interface>
                    </neighbor>
                </neighbor-list>
            </member>
        </member-list>
    </virtual-chassis-information>
</rpc-reply>`

    var reply VirtualChassisReply
    err := xml.Unmarshal([]byte(body), &reply)

    if err != nil {
        t.Fatal(err)
    }

    assert.Equal(t, string{"1234.1234.1234"}, reply.VirtChassInfo.PreProvisVCInfo, "virtual-chassis-id")
    assert.Len(t, reply.VirtChassInfo.MemberList, 2)

    member0 := reply.VirtChassInfo.MemberList[0]
    assert.Equal(t, 0, member0.MemberID, "member-id")
    assert.Equal(t, 129, member0.MemberPriority, "member-priority")
    assert.Equal(t, []string{"Y","N"}, member0.MemberMixedMode, "member-mixed-mode")
    assert.Equal(t, []string{"Master*", "Backup"}, member0.MemberRole, "member-role")

    member1 := reply.VirtChassInfo.MemberList[1]
    assert.Equal(t, 1, member1.MemberID, "member-id")
    assert.Equal(t, 129, member1.MemberPriority, "member-priority")
    assert.Equal(t, []string{"Y","N"}, member1.MemberMixedMode, "member-mixed-mode")
    assert.Equal(t, []string{"Master*", "Backup"}, member1.MemberRole, "member-role")
}
