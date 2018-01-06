package interfaces

type InterfaceStats struct {
	Name string
	Description string
	Mac string
	ReceiveBytes   float64
	ReceiveErrors  float64
	ReceiveDrops   float64
	TransmitBytes  float64
	TransmitErrors float64
	TransmitDrops  float64
}
