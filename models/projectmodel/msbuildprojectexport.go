package projectmodel

type MsBuildProjectExport struct {
	TargetFrameworks    []string
	DefaultNamespace    string
	AssemblyName        string
	AbsolutePath        string
	FolderPath          string
	ProjectDependencies []*MsBuildProjectReference
	SolutionParents     []*MsBuildSolutionReference
	PublishProfiles     []*MsBuildPublishProfile
	IsNetCoreProject    bool
	Name                string
}
