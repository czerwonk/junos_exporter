// SPDX-License-Identifier: MIT

package isis

import "encoding/xml"

type result struct {
	Information struct {
		Adjacencies []adjacency `xml:"isis-adjacency"`
	} `xml:"isis-adjacency-information"`
}

type adjacency struct {
	InterfaceName  string `xml:"interface-name"`
	SystemName     string `xml:"system-name"`
	Level          int64  `xml:"level"`
	AdjacencyState string `xml:"adjacency-state"`
	Holdtime       int64  `xml:"holdtime"`
	SNPA           string `xml:"snpa"`
}

type interfaces struct {
	XMLName                  xml.Name `xml:"rpc-reply"`
	Text                     string   `xml:",chardata"`
	Junos                    string   `xml:"junos,attr"`
	IsisInterfaceInformation struct {
		Text          string `xml:",chardata"`
		Xmlns         string `xml:"xmlns,attr"`
		IsisInterface []struct {
			InterfaceName      string  `xml:"interface-name"`
			LSPInterval        float64 `xml:"lsp-interval"`
			CSNPInterval       float64 `xml:"csnp-interval"`
			HelloPadding       string  `xml:"hello-padding"`
			MaxHelloSize       float64 `xml:"max-hello-size"`
			InterfaceLevelData struct {
				Level             string  `xml:"level"`
				AdjacencyCount    float64 `xml:"adjacency-count"`
				InterfacePriority float64 `xml:"interface-priority"`
				Metric            float64 `xml:"metric"`
				HelloTime         float64 `xml:"hello-time"`
				HoldTime          float64 `xml:"holdtime"`
				Passive           string  `xml:"passive"`
			} `xml:"interface-level-data"`
		} `xml:"isis-interface"`
	} `xml:"isis-interface-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

type backupCoverage struct {
	XMLName                       xml.Name `xml:"rpc-reply"`
	Text                          string   `xml:",chardata"`
	Junos                         string   `xml:"junos,attr"`
	IsisBackupCoverageInformation struct {
		Text               string `xml:",chardata"`
		Xmlns              string `xml:"xmlns,attr"`
		IsisBackupCoverage struct {
			Text                          string `xml:",chardata"`
			IsisTopologyID                string `xml:"isis-topology-id"`
			Level                         string `xml:"level"`
			IsisNodeCoverage              string `xml:"isis-node-coverage"`
			IsisRouteCoverageIpv4         string `xml:"isis-route-coverage-ipv4"`
			IsisRouteCoverageIpv6         string `xml:"isis-route-coverage-ipv6"`
			IsisRouteCoverageClns         string `xml:"isis-route-coverage-clns"`
			IsisRouteCoverageIpv4Mpls     string `xml:"isis-route-coverage-ipv4-mpls"`
			IsisRouteCoverageIpv6Mpls     string `xml:"isis-route-coverage-ipv6-mpls"`
			IsisRouteCoverageIpv4MplsSspf string `xml:"isis-route-coverage-ipv4-mpls-sspf"`
			IsisRouteCoverageIpv6MplsSspf string `xml:"isis-route-coverage-ipv6-mpls-sspf"`
		} `xml:"isis-backup-coverage"`
	} `xml:"isis-backup-coverage-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

type backupSPF struct {
	XMLName            xml.Name `xml:"rpc-reply"`
	Text               string   `xml:",chardata"`
	Junos              string   `xml:"junos,attr"`
	IsisSpfInformation struct {
		Text    string `xml:",chardata"`
		Xmlns   string `xml:"xmlns,attr"`
		IsisSpf []struct {
			Text                 string `xml:",chardata"`
			IsisSpfResultsHeader struct {
				Level int `xml:"level"` // Converted to int since it's numeric
			} `xml:"isis-spf-results-header"`
			NodeCount           float64 `xml:"node-count"` // Converted to int since it's numeric
			IsisBackupSpfResult []struct {
				NodeID         string `xml:"node-id"`
				NodeAddress    string `xml:"node-address"`
				NextHopElement []struct {
					InterfaceName   string `xml:"interface-name"`
					IsisNextHopType string `xml:"isis-next-hop-type"`
					IsisNextHop     string `xml:"isis-next-hop"`
					SNPA            string `xml:"snpa"`
				} `xml:"next-hop-element"`
				BackupRoot              string  `xml:"backup-root"`        // Singular field
				BackupRootMetric        float64 `xml:"backup-root-metric"` // Converted to int
				Metric                  float64 `xml:"metric"`             // Converted to int
				BackupRootPreference    string  `xml:"backup-root-preference"`
				NoCoverageReasonElement []struct {
					IsisNextHopType  string `xml:"isis-next-hop-type"`
					NoCoverageReason string `xml:"no-coverage-reason"`
				} `xml:"no-coverage-reason-element"`
				TrackItem            []string `xml:"track-item"`
				BackupNextHopElement struct {
					InterfaceName            string  `xml:"interface-name"`
					IsisNextHopType          string  `xml:"isis-next-hop-type"`
					IsisBackupPrefixRefcount float64 `xml:"isis-backup-prefix-refcount"` // Converted to int
					IsisNextHop              string  `xml:"isis-next-hop"`
					SNPA                     string  `xml:"snpa"`
				} `xml:"backup-next-hop-element"`
			} `xml:"isis-backup-spf-result"`
		} `xml:"isis-spf"`
	} `xml:"isis-spf-information"`
	Cli struct {
		Banner string `xml:"banner"`
	} `xml:"cli"`
}

/*type backupSPF struct {
	XMLName            xml.Name `xml:"rpc-reply"`
	Text               string   `xml:",chardata"`
	Junos              string   `xml:"junos,attr"`
	IsisSpfInformation struct {
		Text    string `xml:",chardata"`
		Xmlns   string `xml:"xmlns,attr"`
		IsisSpf []struct {
			Text                 string `xml:",chardata"`
			IsisSpfResultsHeader struct {
				Text  string `xml:",chardata"`
				Level string `xml:"level"`
			} `xml:"isis-spf-results-header"`
			NodeCount           string `xml:"node-count"`
			IsisBackupSpfResult []struct {
				Text           string `xml:",chardata"`
				NodeID         string `xml:"node-id"`
				NodeAddress    string `xml:"node-address"`
				NextHopElement []struct {
					Text            string `xml:",chardata"`
					InterfaceName   string `xml:"interface-name"`
					IsisNextHopType string `xml:"isis-next-hop-type"`
					IsisNextHop     string `xml:"isis-next-hop"`
					Snpa            string `xml:"snpa"`
				} `xml:"next-hop-element"`
				BackupRoot              []string `xml:"backup-root"`
				BackupRootMetric        []string `xml:"backup-root-metric"`
				Metric                  []string `xml:"metric"`
				BackupRootPreference    []string `xml:"backup-root-preference"`
				NoCoverageReasonElement []struct {
					Text             string `xml:",chardata"`
					IsisNextHopType  string `xml:"isis-next-hop-type"`
					NoCoverageReason string `xml:"no-coverage-reason"`
				} `xml:"no-coverage-reason-element"`
				TrackItem            []string `xml:"track-item"`
				BackupNextHopElement struct {
					Text                     string `xml:",chardata"`
					InterfaceName            string `xml:"interface-name"`
					IsisNextHopType          string `xml:"isis-next-hop-type"`
					IsisBackupPrefixRefcount string `xml:"isis-backup-prefix-refcount"`
					IsisNextHop              string `xml:"isis-next-hop"`
					Snpa                     string `xml:"snpa"`
				} `xml:"backup-next-hop-element"`
			} `xml:"isis-backup-spf-result"`
		} `xml:"isis-spf"`
	} `xml:"isis-spf-information"`
	Cli struct {
		Text   string `xml:",chardata"`
		Banner string `xml:"banner"`
	} `xml:"cli"`
}
*/
