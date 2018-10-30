package inframodel

type IisApplication struct {
	Sites           []IisSite
	ApplicationName string
	PhysicalPath    string
	AppPool         IisApplicationPool
	ApplicationGuid string
}
