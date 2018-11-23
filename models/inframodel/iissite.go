package inframodel

type IisSite struct {
	SiteName     string
	PhysicalPath string
	AppPool      IisApplicationPool
	SiteGuid     string
	Environments []string
	Applications []IisApplication
}
