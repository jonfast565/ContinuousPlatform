package importmodels

type IisApplicationRecord struct {
	Name            string
	VirtualPath     string
	ApplicationPool IisApplicationPoolRecord
	PhysicalPath    string
}
