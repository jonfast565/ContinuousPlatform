package projectmodel

type MsBuildProject struct {
	TargetFrameworks              []string
	DefaultNamespace              string
	AssemblyName                  string
	AbsolutePath                  string
	FolderPath                    string
	RelativeProjectReferencePaths []string
	AbsoluteProjectReferencePaths []string
	ProjectDependencies           []*MsBuildProject
	SolutionParents               []*MsBuildSolution
	IsNetCoreProject              bool
	Name                          string
	Failed                        bool
	Exception                     DotNetException
}
