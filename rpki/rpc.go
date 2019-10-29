package rpki

type RpkiSessionRpc struct {
	Information struct {
		RpkiSessions []RpkiSession `xml:"rv-session"`
	} `xml:"rv-session-information"`
}

type RpkiSession struct {
	IpAddress       string `xml:"ip-address"`
	SessionState    string `xml:"session-state"`
	SessionFlaps    int64  `xml:"session-flaps"`
	Ipv4PrefixCount int64  `xml:"ip-prefix-count"`
	Ipv6PrefixCount int64  `xml:"ip6-prefix-count"`
}

type RpkiStatisticsRpc struct {
	Information struct {
		Statistics RpkiStatistics `xml:"rv-statistics"`
	} `xml:"rv-statistics-information"`
}

type RpkiStatistics struct {
	RecordCount            int64 `xml:"rv-record-count"`
	ReplicationRecordCount int64 `xml:"rv-replication-record-count"`
	PrefixCount            int64 `xml:"rv-prefix-count"`
	OriginASCount          int64 `xml:"rv-origin-as-count"`
	MemoryUtilization      int64 `xml:"rv-memory-utilization"`
	OriginResultsValid     int64 `xml:"rv-policy-origin-validation-results-valid"`
	OriginResultsInvalid   int64 `xml:"rv-policy-origin-validation-results-invalid"`
	OriginResultsUnknown   int64 `xml:"rv-policy-origin-validation-results-unknown"`
}
