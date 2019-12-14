package storage

type MultiRoutingEngineResults struct {
	Results []MultiRoutingEngineItem `xml:"multi-routing-engine-results"`
}

type MultiRoutingEngineItem struct {
	Name    string             `xml:"re-name"`
	Storage StorageInformation `xml:"multi-routing-engine-item"`
}

type StorageInformation struct {
	Information struct {
		Filesystems []Filesystem `xml:"filesystem"`
	} `xml:"system-storage-information"`
}

type Filesystem struct {
	FilesystemName  string `xml:"filesystem-name"`
	TotalBlocks     int64  `xml:"total-blocks"`
	UsedBlocks      int64  `xml:"used-blocks"`
	AvailableBlocks int64  `xml:"available-blocks"`
	UsedPercent     string `xml:"used-percent"`
	MountedOn       string `xml:"mounted-on"`
}
