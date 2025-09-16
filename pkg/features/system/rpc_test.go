package system

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemCommit(t *testing.T) {
	input := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/24.2R2-S1.6/junos">
    <commit-information>
        <commit-history>
            <sequence-number>0</sequence-number>
            <user>user</user>
            <client>cli</client>
            <date-time junos:seconds="1757518615">2025-09-10 15:36:55 UTC</date-time>
            <log>Some log A</log>
        </commit-history>
        <commit-history>
            <sequence-number>1</sequence-number>
            <user>user</user>
            <client>cli</client>
            <date-time junos:seconds="1757518493">2025-09-10 15:34:53 UTC</date-time>
            <comment>Some log B</comment>
        </commit-history>
	</commit-information>
	</rpc-reply>`

	data := &systemCommit{}
	err := xml.Unmarshal([]byte(input), data)
	if err != nil {
		t.Errorf("unable to unmarshal XML: %v", err)
	}

	assert.Equal(t, 1757518615, data.CommitInfo.CommitHistory[0].DateTime.Seconds)
}
