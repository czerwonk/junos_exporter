package interfaces

type InterfaceStats struct {
	Name           string
	AdminStatus    bool
	OperStatus     bool
	Description    string
	Mac            string
	IsPhysical     bool
	ReceiveBytes   float64
	ReceiveErrors  float64
	ReceiveDrops   float64
	TransmitBytes  float64
	TransmitErrors float64
	TransmitDrops  float64
}
