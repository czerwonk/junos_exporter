// SPDX-License-Identifier: MIT

package ddosprotection

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParseXML tests the XML parsing of the ddos-protection protocols statistics
func TestParseXML(t *testing.T) {
	resultStatsData := `
<rpc-reply xmlns:junos="http://xml.juniper.net/junos/23.4R2-S3.9/junos">
    <ddos-protocols-information xmlns="http://xml.juniper.net/junos/23.4R0/junos-jddosd" junos:style="statistics">
        <total-packet-types>253</total-packet-types>
        <packet-types-rcvd-packets>45</packet-types-rcvd-packets>
        <packet-types-in-violation>0</packet-types-in-violation>
        <ddos-protocol-group>
            <group-name>resolve</group-name>
            <ddos-protocol>
                <packet-type>aggregate</packet-type>
                <ddos-system-statistics junos:style="clean-aggr">
                    <packet-received>6</packet-received>
                    <packet-arrival-rate>0</packet-arrival-rate>
                    <packet-dropped>0</packet-dropped>
                    <packet-arrival-rate-max>0</packet-arrival-rate-max>
                </ddos-system-statistics>
                <ddos-instance junos:style="detail">
                    <protocol-states-locale>Routing Engine</protocol-states-locale>
                    <ddos-instance-statistics junos:style="clean-aggr">
                        <packet-received>1</packet-received>
                        <packet-arrival-rate>2</packet-arrival-rate>
                        <packet-dropped>3</packet-dropped>
                        <packet-arrival-rate-max>4</packet-arrival-rate-max>
                        <packet-dropped-others>5</packet-dropped-others>
                    </ddos-instance-statistics>
                </ddos-instance>
                <ddos-instance junos:style="detail">
                    <protocol-states-locale>FPC slot 0</protocol-states-locale>
                    <ddos-instance-statistics junos:style="clean-aggr">
                        <packet-received>6</packet-received>
                        <packet-arrival-rate>7</packet-arrival-rate>
                        <packet-dropped>8</packet-dropped>
                        <packet-arrival-rate-max>9</packet-arrival-rate-max>
                        <packet-dropped-others>10</packet-dropped-others>
                        <packet-dropped-flows>11</packet-dropped-flows>
                    </ddos-instance-statistics>
                </ddos-instance>
            </ddos-protocol>
        </ddos-protocol-group>
    </ddos-protocols-information>
    <cli>
        <banner></banner>
    </cli>
</rpc-reply>`

	resultParamsData := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/23.4R2-S3.9/junos">
    <ddos-protocols-information xmlns="http://xml.juniper.net/junos/23.4R0/junos-jddosd" junos:style="parameters">
        <total-packet-types>253</total-packet-types>
        <mod-packet-types>0</mod-packet-types>
                <ddos-protocol-group>
            <group-name>PFCP</group-name>
            <ddos-protocol>
                <packet-type>aggregate</packet-type>
                <packet-type-description>Aggregate for all PFCP control traffic</packet-type-description>
                <ddos-basic-parameters junos:style="aggr-ext">
                    <policer-bandwidth>6000</policer-bandwidth>
                    <policer-burst>6001</policer-burst>
                    <policer-priority>Medium</policer-priority>
                    <policer-time-recover>300</policer-time-recover>
                    <policer-enable>Yes</policer-enable>
                </ddos-basic-parameters>
                <ddos-instance junos:style="detail">
                    <protocol-states-locale>Routing Engine</protocol-states-locale>
                    <ddos-instance-parameters junos:style="re">
                        <policer-bandwidth>6002</policer-bandwidth>
                        <policer-burst>6003</policer-burst>
                        <policer-enable>enabled</policer-enable>
                    </ddos-instance-parameters>
                </ddos-instance>
                <ddos-instance junos:style="detail">
                    <protocol-states-locale>FPC slot 0</protocol-states-locale>
                    <ddos-instance-parameters junos:style="fpc">
                        <policer-bandwidth-scale>100</policer-bandwidth-scale>
                        <policer-bandwidth>10000</policer-bandwidth>
                        <policer-burst-scale>80</policer-burst-scale>
                        <policer-burst>8001</policer-burst>
                        <policer-enable>enabled</policer-enable>
                        <hostbound-queue>0</hostbound-queue>
                    </ddos-instance-parameters>
                </ddos-instance>
            </ddos-protocol>
        </ddos-protocol-group>
    </ddos-protocols-information>
    <cli>
        <banner></banner>
    </cli>
</rpc-reply>`

	resultFlowDetectionData := `<rpc-reply xmlns:junos="http://xml.juniper.net/junos/23.4R2-S3.9/junos">
    <ddos-protocols-information xmlns="http://xml.juniper.net/junos/23.4R0/junos-jddosd" junos:style="parameters">
        <total-packet-types>253</total-packet-types>
        <mod-packet-types>0</mod-packet-types>
          <ddos-protocol-group>
            <group-name>sctp</group-name>
            <ddos-protocol>
                <packet-type>aggregate</packet-type>
                <ddos-flow-detection junos:style="detail">
                    <ddos-flow-detection-enabled>off</ddos-flow-detection-enabled>
                    <detection-mode>Automatic</detection-mode>
                    <detect-time>4</detect-time>
                    <log-flows>Yes</log-flows>
                    <recover-time>80</recover-time>
                    <timeout-active-flows>No</timeout-active-flows>
                    <timeout-time>400</timeout-time>
                    <flow-aggregation-level-states>
                        <sub-detection-mode>Automatic</sub-detection-mode>
                        <sub-control-mode>Drop</sub-control-mode>
                        <sub-bandwidth>11</sub-bandwidth>
                        <ifl-detection-mode>Automatic</ifl-detection-mode>
                        <ifl-control-mode>Drop</ifl-control-mode>
                        <ifl-bandwidth>12</ifl-bandwidth>
                        <ifd-detection-mode>Automatic</ifd-detection-mode>
                        <ifd-control-mode>Drop</ifd-control-mode>
                        <ifd-bandwidth>20001</ifd-bandwidth>
                    </flow-aggregation-level-states>
                </ddos-flow-detection>
            </ddos-protocol>
        </ddos-protocol-group>
    </ddos-protocols-information>
    <cli>
        <banner></banner>
    </cli>
</rpc-reply>`
	var resultStats statistics
	var resultParams parameters
	var resultFlowDetection flowDetection

	// Parse the XML data for statistics
	err := xml.Unmarshal([]byte(resultStatsData), &resultStats)
	assert.NoError(t, err)

	assert.Equal(t, "resolve", resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].GroupName)
	assert.Equal(t, "aggregate", resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].PacketType)

	assert.Equal(t, float64(6), resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosSystemStatistics.PacketReceived)
	assert.Equal(t, "0", resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosSystemStatistics.PacketArrivalRate)
	assert.Equal(t, float64(0), resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosSystemStatistics.PacketDropped)
	assert.Equal(t, "0", resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosSystemStatistics.PacketArrivalRateMax)

	assert.Equal(t, "Routing Engine", resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[0].ProtocolStatesLocale)
	assert.Equal(t, float64(1), resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[0].DdosInstanceStatistics.PacketReceived)
	assert.Equal(t, "2", resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[0].DdosInstanceStatistics.PacketArrivalRate)
	assert.Equal(t, float64(3), resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[0].DdosInstanceStatistics.PacketDropped)
	assert.Equal(t, "4", resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[0].DdosInstanceStatistics.PacketArrivalRateMax)
	assert.Equal(t, float64(5), resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[0].DdosInstanceStatistics.PacketDroppedOthers)

	assert.Equal(t, "FPC slot 0", resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[1].ProtocolStatesLocale)
	assert.Equal(t, float64(6), resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[1].DdosInstanceStatistics.PacketReceived)
	assert.Equal(t, "7", resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[1].DdosInstanceStatistics.PacketArrivalRate)
	assert.Equal(t, float64(8), resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[1].DdosInstanceStatistics.PacketDropped)
	assert.Equal(t, "9", resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[1].DdosInstanceStatistics.PacketArrivalRateMax)
	assert.Equal(t, float64(10), resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[1].DdosInstanceStatistics.PacketDroppedOthers)
	assert.Equal(t, float64(11), resultStats.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[1].DdosInstanceStatistics.PacketDroppedFlows)

	//parse the XML data for parameters
	err = xml.Unmarshal([]byte(resultParamsData), &resultParams)
	assert.NoError(t, err)

	//basic parameters
	assert.Equal(t, "PFCP", resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].GroupName)
	assert.Equal(t, "aggregate", resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].PacketType)
	assert.Equal(t, "6000", resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosBasicParameters.PolicerBandwidth)
	assert.Equal(t, "6001", resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosBasicParameters.PolicerBurst)
	assert.Equal(t, "Medium", resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosBasicParameters.PolicerPriority)
	assert.Equal(t, "300", resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosBasicParameters.PolicerTimeRecover)
	assert.Equal(t, "Yes", resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosBasicParameters.PolicerEnable)

	//instance parameters
	assert.Equal(t, "Routing Engine", resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[0].ProtocolStatesLocale)
	assert.Equal(t, "6002", resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[0].DdosInstanceParameters.PolicerBandwidth)
	assert.Equal(t, "6003", resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[0].DdosInstanceParameters.PolicerBurst)
	assert.Equal(t, "enabled", resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[0].DdosInstanceParameters.PolicerEnable)
	assert.Equal(t, "FPC slot 0", resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[1].ProtocolStatesLocale)
	assert.Equal(t, "100", resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[1].DdosInstanceParameters.PolicerBandwidthScale)
	assert.Equal(t, "10000", resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[1].DdosInstanceParameters.PolicerBandwidth)
	assert.Equal(t, "80", resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[1].DdosInstanceParameters.PolicerBurstScale)
	assert.Equal(t, "8001", resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[1].DdosInstanceParameters.PolicerBurst)
	assert.Equal(t, "enabled", resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[1].DdosInstanceParameters.PolicerEnable)
	assert.Equal(t, float64(0), resultParams.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosInstance[1].DdosInstanceParameters.HostboundQueue)

	//parse the XML data for flow detection
	err = xml.Unmarshal([]byte(resultFlowDetectionData), &resultFlowDetection)
	assert.NoError(t, err)

	assert.Equal(t, "sctp", resultFlowDetection.DdosProtocolsInformation.DdosProtocolGroup[0].GroupName)
	assert.Equal(t, "aggregate", resultFlowDetection.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].PacketType)

	assert.Equal(t, float64(4), resultFlowDetection.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosFlowDetection.DetectTime)
	assert.Equal(t, float64(80), resultFlowDetection.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosFlowDetection.RecoverTime)
	assert.Equal(t, float64(400), resultFlowDetection.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosFlowDetection.TimeoutTime)
	assert.Equal(t, "Automatic", resultFlowDetection.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosFlowDetection.DetectionMode)
	assert.Equal(t, "Yes", resultFlowDetection.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosFlowDetection.LogFlows)
	assert.Equal(t, "No", resultFlowDetection.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosFlowDetection.TimeoutActiveFlows)

	//flow aggregate level states
	assert.Equal(t, "Automatic", resultFlowDetection.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosFlowDetection.FlowAggregationLevelStates.SubDetectionMode)
	assert.Equal(t, "Drop", resultFlowDetection.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosFlowDetection.FlowAggregationLevelStates.SubControlMode)
	assert.Equal(t, float64(11), resultFlowDetection.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosFlowDetection.FlowAggregationLevelStates.SubBandwidth)
	assert.Equal(t, "Automatic", resultFlowDetection.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosFlowDetection.FlowAggregationLevelStates.IflDetectionMode)
	assert.Equal(t, "Drop", resultFlowDetection.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosFlowDetection.FlowAggregationLevelStates.IflControlMode)
	assert.Equal(t, float64(12), resultFlowDetection.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosFlowDetection.FlowAggregationLevelStates.IflBandwidth)
	assert.Equal(t, "Automatic", resultFlowDetection.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosFlowDetection.FlowAggregationLevelStates.IfdDetectionMode)
	assert.Equal(t, "Drop", resultFlowDetection.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosFlowDetection.FlowAggregationLevelStates.IfdControlMode)
	assert.Equal(t, float64(20001), resultFlowDetection.DdosProtocolsInformation.DdosProtocolGroup[0].DdosProtocol[0].DdosFlowDetection.FlowAggregationLevelStates.IfdBandwidth)
}
