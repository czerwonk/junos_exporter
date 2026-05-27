package krt

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseXML(t *testing.T) {
	resultsData := `
<rpc-reply xmlns:junos="http://xml.juniper.net/junos/23.4R2-S3.9/junos">
    <krt-queue-information xmlns="http://xml.juniper.net/junos/23.4R0/junos-routing">
        <krt-queue>
            <krtq-type>Routing table add queue</krtq-type>
            <krtq-queue-length>1</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Interface add/delete/change queue</krtq-type>
            <krtq-queue-length>11</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Top-priority deletion queue</krtq-type>
            <krtq-queue-length>111</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Top-priority change queue</krtq-type>
            <krtq-queue-length>1111</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Top-priority add queue</krtq-type>
            <krtq-queue-length>11111</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>high priority V4oV6 tcnh delete queue</krtq-type>
            <krtq-queue-length>111111</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>high prioriy anchor gencfg delete queue</krtq-type>
            <krtq-queue-length>2</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>High-priority multicast add/change</krtq-type>
            <krtq-queue-length>22</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Indirect next hop top priority add/change</krtq-type>
            <krtq-queue-length>222</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Indirect next hop add/change</krtq-type>
            <krtq-queue-length>2222</krtq-queue-length>
        </krt-queue>                    
        <krt-queue>
            <krtq-type>high prioriy anchor gencfg add-change queue</krtq-type>
            <krtq-queue-length>22222</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>MPLS add queue</krtq-type>
            <krtq-queue-length>222222</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Indirect next hop delete</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>High-priority deletion queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>MPLS change queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>High-priority change queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>High-priority add queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Normal-priority indirect next hop queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Normal-priority deletion queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Normal-priority composite next hop deletion queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>                     
            <krtq-type>Low prioriy Statistics-id-group deletion queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Normal-priority change queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Normal-priority add queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Least-priority delete queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Least-priority change queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Least-priority add queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Normal-priority pfe table nexthop queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>EVPN gencfg queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Normal-priority gmp queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Routing table delete queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Low priority route retry queue</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
        <krt-queue>
            <krtq-type>Priority queue dependency management system</krtq-type>
            <krtq-queue-length>0</krtq-queue-length>
        </krt-queue>
    </krt-queue-information>
    <cli>
        <banner></banner>
    </cli>
</rpc-reply>`

	var results resultKRT

	// Parse the XML data for krt queue
	err := xml.Unmarshal([]byte(resultsData), &results)
	assert.NoError(t, err)
	// check the values pf queue length
	assert.Equal(t, float64(1), results.KrtQueueInformation.KrtQueue[0].KrtqQueueLength)
	assert.Equal(t, float64(11), results.KrtQueueInformation.KrtQueue[1].KrtqQueueLength)
	assert.Equal(t, float64(111), results.KrtQueueInformation.KrtQueue[2].KrtqQueueLength)
	assert.Equal(t, float64(1111), results.KrtQueueInformation.KrtQueue[3].KrtqQueueLength)
	assert.Equal(t, float64(11111), results.KrtQueueInformation.KrtQueue[4].KrtqQueueLength)
	assert.Equal(t, float64(111111), results.KrtQueueInformation.KrtQueue[5].KrtqQueueLength)
	assert.Equal(t, float64(2), results.KrtQueueInformation.KrtQueue[6].KrtqQueueLength)
	assert.Equal(t, float64(22), results.KrtQueueInformation.KrtQueue[7].KrtqQueueLength)
	assert.Equal(t, float64(222), results.KrtQueueInformation.KrtQueue[8].KrtqQueueLength)
	assert.Equal(t, float64(2222), results.KrtQueueInformation.KrtQueue[9].KrtqQueueLength)
	assert.Equal(t, float64(22222), results.KrtQueueInformation.KrtQueue[10].KrtqQueueLength)
	assert.Equal(t, float64(222222), results.KrtQueueInformation.KrtQueue[11].KrtqQueueLength)
	assert.Equal(t, float64(0), results.KrtQueueInformation.KrtQueue[12].KrtqQueueLength)
}
