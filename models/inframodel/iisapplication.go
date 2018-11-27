package inframodel

type IisApplication struct {
	ApplicationName string
	PhysicalPath    string
	AppPool         IisApplicationPool
	ApplicationGuid string
	Environments    []EnvironmentPart
}
