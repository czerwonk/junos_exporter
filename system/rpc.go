package system

type BuffersRPC struct {
	// some versions don't provide parsed output
	Output           string `xml:"output"`
	MemoryStatistics struct {
		MbufsCurrent int `xml:"current-mbufs"`
		MbufsCache   int `xml:"cached-mbufs"`
		MbufsTotal   int `xml:"total-mbufs"`
		MbufsDenied  int `xml:"mbuf-failures"`

		MbufClustersCurrent int `xml:"current-mbuf-clusters"`
		MbufClustersCache   int `xml:"cached-mbuf-clusters"`
		MbufClustersTotal   int `xml:"total-mbuf-clusters"`
		MbufClustersMax     int `xml:"max-mbuf-clusters"`
		MbufClustersDenied  int `xml:"cluster-failures"`

		MbufClustersFromPacketZoneCurrent int `xml:"packet-count"`
		MbufClustersFromPacketZoneCache   int `xml:"packet-free"`

		JumboClustersCurrent4K int `xml:"current-jumbo-clusters-4k"`
		JumboClustersCache4K   int `xml:"cached-jumbo-clusters-4k"`
		JumboClustersTotal4K   int `xml:"total-jumbo-clusters-4k"`
		JumboClustersMax4K     int `xml:"max-jumbo-clusters-4k"`
		JumboClustersDenied4K  int `xml:"jumbo-cluster-failures-4k"`

		JumboClustersCurrent9K int `xml:"current-jumbo-clusters-9k"`
		JumboClustersCache9K   int `xml:"cached-jumbo-clusters-9k"`
		JumboClustersTotal9K   int `xml:"total-jumbo-clusters-9k"`
		JumboClustersMax9K     int `xml:"max-jumbo-clusters-9k"`
		JumboClustersDenied9K  int `xml:"jumbo-cluster-failures-9k"`

		JumboClustersCurrent16K int `xml:"current-jumbo-clusters-16k"`
		JumboClustersCache16K   int `xml:"cached-jumbo-clusters-16k"`
		JumboClustersTotal16K   int `xml:"total-jumbo-clusters-16k"`
		JumboClustersMax16K     int `xml:"max-jumbo-clusters-16k"`
		JumboClustersDenied16K  int `xml:"jumbo-cluster-failures-16k"`

		NetworkAllocCurrent int `xml:"current-bytes-in-use"`
		NetworkAllocCache   int `xml:"cached-bytes"`
		NetworkAllocTotal   int `xml:"total-bytes"`

		SfbufsDenied  int `xml:"sfbuf-requests-denied"`
		SfbufsDelayed int `xml:"sfbuf-requests-delayed"`

		MbufAndClustersDenied int `xml:"packet-failures"`
		IoInit                int `xml:"io-initiated"`
	} `xml:"memory-statistics"`
}

type SystemInformationRPC struct {
	SysInfo struct {
		Model     string `xml:"hardware-model"`
		OS        string `xml:"os-name"`
		OSVersion string `xml:"os-version"`
		Serial    string `xml:"serial-number"`
		Hostname  string `xml:"host-name"`
	} `xml:"system-information"`
}

type SatelliteChassisRPC struct {
	SatelliteInfo struct {
		Satellite []struct {
			Alias   string `xml:"satellite-alias"`
			SlotId  int    `xml:"slot-id"`
			State   string `xml:"operation-state"`
			Model   string `xml:"product-model"`
			Serial  string `xml:"serial-number"`
			Version string `xml:"version"`
		} `xml:"satellite"`
	} `xml:"satellite-information"`
}
