package projectmodel

type MsBuildSolution struct {
	Configurations       []string
	AbsolutePath         string
	RelativeProjectPaths []string
	AbsoluteProjectPaths []string
	Projects             []MsBuildProject
	Name                 string
	Failed               bool
	Exception            DotNetException
}
