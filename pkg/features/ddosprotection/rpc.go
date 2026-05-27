package ddosprotection

import "encoding/xml"

type statistics struct {
	XMLName                  xml.Name `xml:"rpc-reply"`
	Text                     string   `xml:",chardata"`
	Junos                    string   `xml:"junos,attr"`
	DdosProtocolsInformation struct {
		Text                   string  `xml:",chardata"`
		Xmlns                  string  `xml:"xmlns,attr"`
		Style                  string  `xml:"style,attr"`
		TotalPacketTypes       float64 `xml:"total-packet-types"`
		PacketTypesRcvdPackets float64 `xml:"packet-types-rcvd-packets"`
		PacketTypesInViolation float64 `xml:"packet-types-in-violation"`
		DdosProtocolGroup      []struct {
			Text         string `xml:",chardata"`
			GroupName    string `xml:"group-name"`
			DdosProtocol []struct {
				Text                 string `xml:",chardata"`
				PacketType           string `xml:"packet-type"`
				DdosSystemStatistics struct {
					Text                 string  `xml:",chardata"`
					Style                string  `xml:"style,attr"`
					PacketReceived       float64 `xml:"packet-received"`
					PacketArrivalRate    string  `xml:"packet-arrival-rate"`
					PacketDropped        float64 `xml:"packet-dropped"`
					PacketArrivalRateMax string  `xml:"packet-arrival-rate-max"`
				} `xml:"ddos-system-statistics"`
				DdosInstance []struct {
					Text                   string `xml:",chardata"`
					Style                  string `xml:"style,attr"`
					ProtocolStatesLocale   string `xml:"protocol-states-locale"`
					DdosInstanceStatistics struct {
						Text                 string  `xml:",chardata"`
						Style                string  `xml:"style,attr"`
						PacketReceived       float64 `xml:"packet-received"`
						PacketArrivalRate    string  `xml:"packet-arrival-rate"`
						PacketDropped        float64 `xml:"packet-dropped"`
						PacketArrivalRateMax string  `xml:"packet-arrival-rate-max"`
						PacketDroppedOthers  float64 `xml:"packet-dropped-others"`
						PacketDroppedFlows   float64 `xml:"packet-dropped-flows"`
					} `xml:"ddos-instance-statistics"`
				} `xml:"ddos-instance"`
			} `xml:"ddos-protocol"`
		} `xml:"ddos-protocol-group"`
	} `xml:"ddos-protocols-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

type parameters struct {
	XMLName                  xml.Name `xml:"rpc-reply"`
	Text                     string   `xml:",chardata"`
	Junos                    string   `xml:"junos,attr"`
	DdosProtocolsInformation struct {
		Text              string  `xml:",chardata"`
		Xmlns             string  `xml:"xmlns,attr"`
		Style             string  `xml:"style,attr"`
		TotalPacketTypes  float64 `xml:"total-packet-types"`
		ModPacketTypes    float64 `xml:"mod-packet-types"`
		DdosProtocolGroup []struct {
			Text         string `xml:",chardata"`
			GroupName    string `xml:"group-name"`
			DdosProtocol []struct {
				Text                  string `xml:",chardata"`
				PacketType            string `xml:"packet-type"`
				PacketTypeDescription string `xml:"packet-type-description"`
				DdosBasicParameters   struct {
					Text                   string `xml:",chardata"`
					Style                  string `xml:"style,attr"`
					PolicerBandwidth       string `xml:"policer-bandwidth"`
					PolicerBurst           string `xml:"policer-burst"`
					PolicerTimeRecover     string `xml:"policer-time-recover"`
					PolicerEnable          string `xml:"policer-enable"`
					PolicerPriority        string `xml:"policer-priority"`
					PolicerBypassAggregate string `xml:"policer-bypass-aggregate"`
				} `xml:"ddos-basic-parameters"`
				DdosInstance []struct {
					Text                   string `xml:",chardata"`
					Style                  string `xml:"style,attr"`
					ProtocolStatesLocale   string `xml:"protocol-states-locale"`
					DdosInstanceParameters struct {
						Text                  string  `xml:",chardata"`
						Style                 string  `xml:"style,attr"`
						PolicerBandwidth      string  `xml:"policer-bandwidth"`
						PolicerBurst          string  `xml:"policer-burst"`
						PolicerEnable         string  `xml:"policer-enable"`
						PolicerBandwidthScale string  `xml:"policer-bandwidth-scale"`
						PolicerBurstScale     string  `xml:"policer-burst-scale"`
						HostboundQueue        float64 `xml:"hostbound-queue"`
					} `xml:"ddos-instance-parameters"`
				} `xml:"ddos-instance"`
			} `xml:"ddos-protocol"`
		} `xml:"ddos-protocol-group"`
	} `xml:"ddos-protocols-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

type flowDetection struct {
	XMLName                  xml.Name `xml:"rpc-reply"`
	Text                     string   `xml:",chardata"`
	Junos                    string   `xml:"junos,attr"`
	DdosProtocolsInformation struct {
		Text              string  `xml:",chardata"`
		Xmlns             string  `xml:"xmlns,attr"`
		Style             string  `xml:"style,attr"`
		TotalPacketTypes  float64 `xml:"total-packet-types"`
		ModPacketTypes    float64 `xml:"mod-packet-types"`
		DdosProtocolGroup []struct {
			Text         string `xml:",chardata"`
			GroupName    string `xml:"group-name"`
			DdosProtocol []struct {
				Text              string `xml:",chardata"`
				PacketType        string `xml:"packet-type"`
				DdosFlowDetection struct {
					Text                       string  `xml:",chardata"`
					Style                      string  `xml:"style,attr"`
					DdosFlowDetectionEnabled   string  `xml:"ddos-flow-detection-enabled"`
					DetectionMode              string  `xml:"detection-mode"`
					DetectTime                 float64 `xml:"detect-time"`
					LogFlows                   string  `xml:"log-flows"`
					RecoverTime                float64 `xml:"recover-time"`
					TimeoutActiveFlows         string  `xml:"timeout-active-flows"`
					TimeoutTime                float64 `xml:"timeout-time"`
					FlowAggregationLevelStates struct {
						Text             string  `xml:",chardata"`
						SubDetectionMode string  `xml:"sub-detection-mode"`
						SubControlMode   string  `xml:"sub-control-mode"`
						SubBandwidth     float64 `xml:"sub-bandwidth"`
						IflDetectionMode string  `xml:"ifl-detection-mode"`
						IflControlMode   string  `xml:"ifl-control-mode"`
						IflBandwidth     float64 `xml:"ifl-bandwidth"`
						IfdDetectionMode string  `xml:"ifd-detection-mode"`
						IfdControlMode   string  `xml:"ifd-control-mode"`
						IfdBandwidth     float64 `xml:"ifd-bandwidth"`
					} `xml:"flow-aggregation-level-states"`
				} `xml:"ddos-flow-detection"`
			} `xml:"ddos-protocol"`
		} `xml:"ddos-protocol-group"`
	} `xml:"ddos-protocols-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}
