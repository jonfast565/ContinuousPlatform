package projectmodel

type DotNetDeliverable struct {
}

type DotNetException struct {
	Message string
}

type MsBuildSolution struct {
	Configurations       []string
	ProjectRelativePaths []string
	Name                 string
	Failed               bool
	Exception            DotNetException
}

type MsBuildProject struct {
	TargetFrameworks              []string
	DefaultNamespace              string
	AssemblyName                  string
	RelativeProjectReferencePaths []string
	IsNetCoreProject              bool
	Name                          string
	Failed                        bool
	Exception                     DotNetException
}

type MsBuildPublishProfile struct {
	PublishUrl string
	Name       string
	Failed     bool
	Exception  DotNetException
}
