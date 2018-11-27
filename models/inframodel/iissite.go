package inframodel

type IisSite struct {
	SiteName     string
	PhysicalPath string
	AppPool      IisApplicationPool
	SiteGuid     string
	Applications []IisApplication
	Environments []EnvironmentPart
}

type IisSitePart struct {
	SiteName     string
	PhysicalPath string
	SiteGuid     string
	Environments []EnvironmentPart
}
