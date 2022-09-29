package security_policies

import "encoding/xml"

type StatsRpcReply struct {
	XMLName                   xml.Name                       `xml:"rpc-reply"`
	MultiRoutingEngineResults StatsMultiRoutingEngineResults `xml:"multi-routing-engine-results"`
}

type StatsMultiRoutingEngineResults struct {
	RoutingEngine []StatsRoutingEngine `xml:"multi-routing-engine-item"`
}

type StatsRoutingEngine struct {
	Name     string           `xml:"re-name"`
	Policies SecurityPolicies `xml:"security-policies"`
}

type StatsRpcReplyNoRE struct {
	XMLName  xml.Name         `xml:"rpc-reply"`
	Policies SecurityPolicies `xml:"security-policies"`
}

type SecurityPolicies struct {
	Contexts []SecurityContext `xml:"security-context"`
}

type SecurityContext struct {
	ContextInformation SecurityContextInformation `xml:"context-information"`
	Policies           []SecurityPolicy           `xml:"policies"`
}

type SecurityContextInformation struct {
	FromZone      string    `xml:"source-zone-name"`
	ToZone        string    `xml:"destination-zone-name"`
	GlobalContext *struct{} `xml:"global-context"`
}

type SecurityPolicy struct {
	PolicyInformation SecurityPolicyInformation `xml:"policy-information"`
}

type SecurityPolicyInformation struct {
	Name   string `xml:"policy-name"`
	Action struct {
		ActionType string `xml:"action-type"`
	} `xml:"policy-action"`
	StatisticsInformation *SecurityPolicyStatisticsInformation `xml:"policy-statistics-information"`
}

type SecurityPolicyStatisticsInformation struct {
	InputBytesInit     float64 `xml:"input-bytes-init"`
	InputBytesReply    float64 `xml:"input-bytes-reply"`
	OutputBytesInit    float64 `xml:"output-bytes-init"`
	OutputBytesReply   float64 `xml:"output-bytes-reply"`
	InputPacketsInit   float64 `xml:"input-packets-init"`
	InputPacketsReply  float64 `xml:"input-packets-reply"`
	OutputPacketsInit  float64 `xml:"output-packets-init"`
	OutputPacketsReply float64 `xml:"output-packets-reply"`
	SessionCreations   float64 `xml:"session-creations"`
	SessionDeletions   float64 `xml:"session-deletions"`
}

type HitsRpcReply struct {
	XMLName                   xml.Name                      `xml:"rpc-reply"`
	MultiRoutingEngineResults HitsMultiRoutingEngineResults `xml:"multi-routing-engine-results"`
}

type HitsMultiRoutingEngineResults struct {
	RoutingEngine []HitsRoutingEngine `xml:"multi-routing-engine-item"`
}

type HitsRoutingEngine struct {
	Name     string   `xml:"re-name"`
	HitCount HitCount `xml:"policy-hit-count"`
}

type HitsRpcReplyNoRE struct {
	XMLName  xml.Name `xml:"rpc-reply"`
	HitCount HitCount `xml:"policy-hit-count"`
}

type HitCount struct {
	LSName   string          `xml:"logical-system-name"`
	Policies []HitCountEntry `xml:"policy-hit-count-entry"`
}

type HitCountEntry struct {
	PolicyName string  `xml:"policy-hit-count-policy-name"`
	FromZone   string  `xml:"policy-hit-count-from-zone"`
	ToZone     string  `xml:"policy-hit-count-to-zone"`
	Count      float64 `xml:"policy-hit-count-count"`
}
