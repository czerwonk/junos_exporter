package interfaces

type InterfaceStats struct {
	Name                string
	AdminStatus         bool
	OperStatus          bool
	ErrorStatus         bool
	Description         string
	Mac                 string
	IsPhysical          bool
	ReceiveBytes        float64
	ReceivePackets      float64
	ReceiveErrors       float64
	ReceiveDrops        float64
	TransmitBytes       float64
	TransmitPackets     float64
	TransmitErrors      float64
	TransmitDrops       float64
	IPv6ReceiveBytes    float64
	IPv6ReceivePackets  float64
	IPv6TransmitBytes   float64
	IPv6TransmitPackets float64
}
