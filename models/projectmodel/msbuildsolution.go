package projectmodel

type MsBuildSolution struct {
	Configurations       []string
	AbsolutePath         string
	FolderPath           string
	RelativeProjectPaths []string
	AbsoluteProjectPaths []string
	Projects             []*MsBuildProject
	Name                 string
	Failed               bool
	Exception            DotNetException
}

func (msbs MsBuildSolution) GetSolutionReference() MsBuildSolutionReference {
	return MsBuildSolutionReference{
		Name:         msbs.Name,
		AbsolutePath: msbs.AbsolutePath,
	}
}

func (msbs MsBuildSolution) ExportProjects() []*MsBuildProjectExport {
	var projectExports []*MsBuildProjectExport
	for _, project := range msbs.Projects {
		projectExports = append(projectExports, project.Export())
	}
	return projectExports
}

func (msbs MsBuildSolution) Export() *MsBuildSolutionExport {
	return &MsBuildSolutionExport{
		Configurations: msbs.Configurations,
		AbsolutePath:   msbs.AbsolutePath,
		FolderPath:     msbs.FolderPath,
		Projects:       msbs.ExportProjects(),
		Name:           msbs.Name,
	}
}
