package rpc

type IsisRpc struct {
	Information struct {
		Adjacencies []IsisAdjacencies `xml:"isis-adjacency"`
	} `xml:"isis-adjacency-information"`
}

type IsisAdjacencies struct {
	InterfaceName  string `xml:"interface-name"`
	SystemName     string `xml:"system-name"`
	Level          int64  `xml:"level"`
	AdjacencyState string `xml:"adjacency-state"`
	Holdtime       int64  `xml:"holdtime"`
	SNPA           string `xml:"snpa"`
}
