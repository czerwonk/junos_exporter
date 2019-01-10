package storage

type StorageRpc struct {
	Information struct {
		Filesystems []Filesystem `xml:"filesystem"`
	} `xml:"system-storage-information"`
}

type Filesystem struct {
	FilesystemName  string `xml:"filesystem-name"`
	TotalBlocks     int64  `xml:"total-blocks"`
	UsedBlocks      int64  `xml:"used-blocks"`
	AvailableBlocks int64  `xml:"available-blocks"`
	UsedPercent     int64  `xml:"used-percent"`
	MountedOn       string `xml:"mounted-on"`
}
