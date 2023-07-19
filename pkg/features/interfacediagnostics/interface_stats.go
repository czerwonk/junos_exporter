package interfacediagnostics

type transceiverInformation struct {
	Name                string
	ChassisHardwareInfo *chassisSubSubModule
	PicPort             *picPort
}
