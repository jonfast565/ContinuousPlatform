package projectmodel

type MsBuildProjectExport struct {
	TargetFrameworks    []string
	DefaultNamespace    string
	AssemblyName        string
	AbsolutePath        string
	FolderPath          string
	ProjectDependencies []*MsBuildProjectReference
	SolutionParents     []*MsBuildSolutionReference
	IsNetCoreProject    bool
	Name                string
}
