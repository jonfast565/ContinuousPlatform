package projectmodel

type MsBuildSolution struct {
	Configurations       []string
	AbsolutePath         string
	RelativeProjectPaths []string
	Projects             []MsBuildProject
	Name                 string
	Failed               bool
	Exception            DotNetException
}
