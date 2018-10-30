package projectmodel

type MsBuildSolution struct {
	Configurations       []string
	AbsolutePath         string
	RelativeProjectPaths []string
	Name                 string
	Failed               bool
	Exception            DotNetException
}
