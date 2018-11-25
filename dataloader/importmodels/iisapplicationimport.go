package importmodels

type IisApplicationImport struct {
	Name            string
	VirtualPath     string
	ApplicationPool IisApplicationPoolImport
	PhysicalPath    string
}
