package storage

import "encoding/xml"

type multiEngineResult struct {
	XMLName xml.Name       `xml:"rpc-reply"`
	Results routingEngines `xml:"multi-routing-engine-results"`
}

type routingEngines struct {
	RoutingEngine []routingEngine `xml:"multi-routing-engine-item"`
}

type routingEngine struct {
	Name               string             `xml:"re-name"`
	StorageInformation storageInformation `xml:"system-storage-information"`
}

type storageInformation struct {
	Filesystems []filesystem `xml:"filesystem"`
}

type filesystem struct {
	FilesystemName  string `xml:"filesystem-name"`
	TotalBlocks     int64  `xml:"total-blocks"`
	UsedBlocks      int64  `xml:"used-blocks"`
	AvailableBlocks int64  `xml:"available-blocks"`
	UsedPercent     string `xml:"used-percent"`
	MountedOn       string `xml:"mounted-on"`
}

type singleEngineResult struct {
	XMLName            xml.Name           `xml:"rpc-reply"`
	StorageInformation storageInformation `xml:"system-storage-information"`
}
