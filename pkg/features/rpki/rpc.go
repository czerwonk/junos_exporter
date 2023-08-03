// SPDX-License-Identifier: MIT

package rpki

type sessionResult struct {
	Information struct {
		Sessions []session `xml:"rv-session"`
	} `xml:"rv-session-information"`
}

type session struct {
	IPAddress       string `xml:"ip-address"`
	State           string `xml:"session-state"`
	Flaps           int64  `xml:"session-flaps"`
	IPv4PrefixCount int64  `xml:"ip-prefix-count"`
	IPv6PrefixCount int64  `xml:"ip6-prefix-count"`
}

type statisticResult struct {
	Information struct {
		Statistics statistics `xml:"rv-statistics"`
	} `xml:"rv-statistics-information"`
}

type statistics struct {
	RecordCount            int64 `xml:"rv-record-count"`
	ReplicationRecordCount int64 `xml:"rv-replication-record-count"`
	PrefixCount            int64 `xml:"rv-prefix-count"`
	OriginASCount          int64 `xml:"rv-origin-as-count"`
	MemoryUtilization      int64 `xml:"rv-memory-utilization"`
	OriginResultsValid     int64 `xml:"rv-policy-origin-validation-results-valid"`
	OriginResultsInvalid   int64 `xml:"rv-policy-origin-validation-results-invalid"`
	OriginResultsUnknown   int64 `xml:"rv-policy-origin-validation-results-unknown"`
}
