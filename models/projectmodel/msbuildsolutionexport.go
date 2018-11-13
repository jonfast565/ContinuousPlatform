package projectmodel

type MsBuildSolutionExport struct {
	Configurations []string
	AbsolutePath   string
	FolderPath     string
	Projects       []*MsBuildProjectExport
	Name           string
}
