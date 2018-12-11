package projectmodel

type MsBuildProject struct {
	TargetFrameworks              []string
	DefaultNamespace              string
	AssemblyName                  string
	AbsolutePath                  string
	FolderPath                    string
	RelativeProjectReferencePaths []string
	AbsoluteProjectReferencePaths []string
	ProjectDependencies           []*MsBuildProjectReference
	SolutionParents               []*MsBuildSolutionReference
	PublishProfiles               []*MsBuildPublishProfile
	IsNetCoreProject              bool
	Name                          string
	Failed                        bool
	Exception                     DotNetException
}

func (p *MsBuildProject) GetReference() *MsBuildProjectReference {
	return &MsBuildProjectReference{
		Name:         p.Name,
		AbsolutePath: p.AbsolutePath,
	}
}

func (p MsBuildProject) Export() *MsBuildProjectExport {
	return &MsBuildProjectExport{
		TargetFrameworks:    p.TargetFrameworks,
		DefaultNamespace:    p.DefaultNamespace,
		PublishProfiles:     p.PublishProfiles,
		AssemblyName:        p.AssemblyName,
		AbsolutePath:        p.AbsolutePath,
		FolderPath:          p.FolderPath,
		ProjectDependencies: p.ProjectDependencies,
		SolutionParents:     p.SolutionParents,
		IsNetCoreProject:    p.IsNetCoreProject,
		Name:                p.Name,
	}
}
