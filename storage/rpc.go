package storage

import "encoding/xml"

type RpcReply struct {
	XMLName                   xml.Name                  `xml:"rpc-reply"`
	MultiRoutingEngineResults MultiRoutingEngineResults `xml:"multi-routing-engine-results"`
}

type MultiRoutingEngineResults struct {
	RoutingEngine []RoutingEngine `xml:"multi-routing-engine-item"`
}

type RoutingEngine struct {
	Name               string             `xml:"re-name"`
	StorageInformation StorageInformation `xml:"system-storage-information"`
}

type StorageInformation struct {
	Filesystems []Filesystem `xml:"filesystem"`
}

type Filesystem struct {
	FilesystemName  string `xml:"filesystem-name"`
	TotalBlocks     int64  `xml:"total-blocks"`
	UsedBlocks      int64  `xml:"used-blocks"`
	AvailableBlocks int64  `xml:"available-blocks"`
	UsedPercent     string `xml:"used-percent"`
	MountedOn       string `xml:"mounted-on"`
}

type RpcReplyNoRE struct {
	XMLName            xml.Name           `xml:"rpc-reply"`
	StorageInformation StorageInformation `xml:"system-storage-information"`
}
