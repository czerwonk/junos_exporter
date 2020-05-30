package l2circuit

type L2circuitRpc struct {
	Information l2circuitInformation `xml:"l2circuit-connection-information"`
}

type l2circuitInformation struct {
	Neighbors []l2circuitNeighbor `xml:"l2circuit-neighbor"`
}

type l2circuitNeighbor struct {
	Address     string                `xml:"neighbor-address"`
	Connections []l2circuitConnection `xml:"connection"`
}

type l2circuitConnection struct {
	ID           string `xml:"connection-id"`
	Type         string `xml:"connection-type"`
	StatusString string `xml:"connection-status"`
}
