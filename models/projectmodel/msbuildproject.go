package projectmodel

type MsBuildProject struct {
	TargetFrameworks              []string
	DefaultNamespace              string
	AssemblyName                  string
	AbsolutePath                  string
	RelativeProjectReferencePaths []string
	AbsoluteProjectReferencePaths []string
	IsNetCoreProject              bool
	Name                          string
	Failed                        bool
	Exception                     DotNetException
}
