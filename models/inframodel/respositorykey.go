package inframodel

type ResourceKey struct {
	RepositoryName string
	SolutionName   string
	ProjectName    string
}

func (rk ResourceKey) String() string {
	return "(" + rk.RepositoryName + " " + rk.SolutionName + " " + rk.ProjectName + ")"
}
