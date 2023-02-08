// SPDX-License-Identifier: MIT

package securitypolicies

import "encoding/xml"

type statsMultiEngineResult struct {
	XMLName xml.Name                       `xml:"rpc-reply"`
	Results statsMultiRoutingEngineResults `xml:"multi-routing-engine-results"`
}

type statsMultiRoutingEngineResults struct {
	RoutingEngines []statsRoutingEngine `xml:"multi-routing-engine-item"`
}

type statsRoutingEngine struct {
	Name     string           `xml:"re-name"`
	Policies securityPolicies `xml:"security-policies"`
}

type statsSingleEngineResult struct {
	XMLName  xml.Name         `xml:"rpc-reply"`
	Policies securityPolicies `xml:"security-policies"`
}

type securityPolicies struct {
	Contexts []securityContext `xml:"security-context"`
}

type securityContext struct {
	ContextInformation securityContextInformation `xml:"context-information"`
	Policies           []securityPolicy           `xml:"policies"`
}

type securityContextInformation struct {
	FromZone      string    `xml:"source-zone-name"`
	ToZone        string    `xml:"destination-zone-name"`
	GlobalContext *struct{} `xml:"global-context"`
}

type securityPolicy struct {
	PolicyInformation securityPolicyInformation `xml:"policy-information"`
}

type securityPolicyInformation struct {
	Name   string `xml:"policy-name"`
	Action struct {
		ActionType string `xml:"action-type"`
	} `xml:"policy-action"`
	StatisticsInformation *securityPolicyStatisticsInformation `xml:"policy-statistics-information"`
}

type securityPolicyStatisticsInformation struct {
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

type hitsMultiEngineResult struct {
	XMLName                   xml.Name           `xml:"rpc-reply"`
	MultiRoutingEngineResults hitsRoutingEngines `xml:"multi-routing-engine-results"`
}

type hitsRoutingEngines struct {
	RoutingEngine []hitsRoutingEngine `xml:"multi-routing-engine-item"`
}

type hitsRoutingEngine struct {
	Name     string   `xml:"re-name"`
	HitCount hitCount `xml:"policy-hit-count"`
}

type hitsSingleEngineResult struct {
	XMLName  xml.Name `xml:"rpc-reply"`
	HitCount hitCount `xml:"policy-hit-count"`
}

type hitCount struct {
	LSName   string          `xml:"logical-system-name"`
	Policies []hitCountEntry `xml:"policy-hit-count-entry"`
}

type hitCountEntry struct {
	PolicyName string  `xml:"policy-hit-count-policy-name"`
	FromZone   string  `xml:"policy-hit-count-from-zone"`
	ToZone     string  `xml:"policy-hit-count-to-zone"`
	Count      float64 `xml:"policy-hit-count-count"`
}
