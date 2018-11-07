package builders

import (
	"../../../clients/msbuildclient"
	"../../../clients/repoclient"
	"../../../fileutil"
	"../../../logging"
	"../../../models/filesysmodel"
	"../../../models/projectmodel"
	"../../../models/repomodel"
)

var ValidProjectExtensions = []string{
	`^.*\.csproj$`,
	`^.*\.fsproj$`,
	`^.*\.vbproj$`,
}

var ValidSolutionExtensions = []string{
	`^.*\.sln$`,
}

var ValidPublishProfileExtensions = []string{
	`^.*\.pubxml$`,
}

func BuildDotNetDeliverables(metadata repomodel.RepositoryMetadata,
	fileGraph fileutil.FileGraph) ([]projectmodel.DotNetDeliverable, error) {
	results := make([]projectmodel.DotNetDeliverable, 0)

	repoClient := repoclient.NewRepoClient()
	msBuildClient := msbuildclient.NewMsBuildClient()

	solutionPaths, err := getSolutionPaths(metadata)
	if err != nil {
		return nil, err
	}

	if len(solutionPaths) == 0 {
		logging.LogInfo(metadata.Name + " b. " + metadata.Branch + " contains no solution files.")
		return results, nil
	}

	projectPaths, err := getProjectPaths(metadata)
	if err != nil {
		return nil, err
	}

	if len(projectPaths) == 0 {
		logging.LogInfo(metadata.Name + " b. " + metadata.Branch + " contains no project files.")
		return results, nil
	}

	//publishProfilePaths, err := getPublishProfilePaths(metadata)
	//if err != nil {
	//	return nil, err
	//}

	solutions := getSolutionList(solutionPaths, metadata, repoClient, msBuildClient, fileGraph)
	projects := getProjectList(projectPaths, metadata, repoClient, msBuildClient)
	// getPublishProfileList(publishProfilePaths, metadata, repoClient, msBuildClient)

	resolveProjectReferencePaths(projects, fileGraph)
	resolveSolutionReferencePaths(solutions, projects, fileGraph)

	for _, project := range projects {
		resolveProjectDependencies(project, projects)
	}

	for _, solution := range solutions {
		linkProjectSolutions(solution, projects)
	}

	logging.LogInfo("Done building .NET deliverables for " + metadata.Name + " b. " + metadata.Branch)
	return results, nil
}

func resolveSolutionReferencePaths(
	solutions []*projectmodel.MsBuildSolution,
	projects []*projectmodel.MsBuildProject,
	graph fileutil.FileGraph) {
	for _, solution := range solutions {
		solution.AbsoluteProjectPaths = make([]string, 0)
		for _, relativeProjectPath := range solution.RelativeProjectPaths {
			rootItem, err := graph.GetItemByRootPath(relativeProjectPath)
			if err != nil {
				logging.LogInfoMultiline("Could not find project at path.",
					"Path: "+relativeProjectPath,
					"Error: "+err.Error())
				continue
			}
			rootPath := (*rootItem).GetPathString()
			solution.AbsoluteProjectPaths = append(solution.AbsoluteProjectPaths, rootPath)
			found := false
			for _, project := range projects {
				if rootPath != project.AbsolutePath {
					continue
				}
				solution.Projects = append(solution.Projects, project)
				found = true
				break
			}

			if !found {
				logging.LogInfoMultiline("Failed to find solution reference",
					"Project Path: "+rootPath,
					"Solution Name: "+solution.Name)
			}
		}
	}
}

func resolveProjectDependencies(project *projectmodel.MsBuildProject, projects []*projectmodel.MsBuildProject) {
	logging.LogInfo("Resolving dependencies for: " + project.Name)
	project.ProjectDependencies = make([]*projectmodel.MsBuildProject, 0)
	for _, absolutePath := range project.AbsoluteProjectReferencePaths {
		found := false
		for _, referenceProject := range projects {
			if absolutePath != referenceProject.AbsolutePath ||
				absolutePath == project.AbsolutePath {
				continue
			}
			project.ProjectDependencies = append(project.ProjectDependencies, referenceProject)
			found = true
			break
		}

		if !found {
			logging.LogInfoMultiline("Failed to find project reference",
				"Absolute Path: "+absolutePath,
				"Project Name: "+project.Name)
		}
	}
}

func linkProjectSolutions(solution *projectmodel.MsBuildSolution, projects []*projectmodel.MsBuildProject) {
	for _, project := range projects {
		found := false
		for _, projectPath := range solution.AbsoluteProjectPaths {
			if projectPath != project.AbsolutePath {
				continue
			}
			project.SolutionParents = append(project.SolutionParents, solution)
			solution.Projects = append(solution.Projects, project)
			found = true
			break
		}

		if !found {
			logging.LogInfoMultiline("Failed to link solution to project. Project not found.",
				"Solution Name: "+solution.Name,
				"Project Name: "+project.Name)
		}
	}
}

func resolveProjectReferencePaths(projects []*projectmodel.MsBuildProject, graph fileutil.FileGraph) {
	for _, project := range projects {
		project.AbsoluteProjectReferencePaths = make([]string, 0)
		for _, relativeProjectPath := range project.RelativeProjectReferencePaths {
			rootItem, err := graph.GetItemByRootPath(project.AbsolutePath)
			if err != nil {
				logging.LogInfoMultiline("Could not find project path. Ignoring.",
					"Project Name: "+project.Name,
					"Path: "+relativeProjectPath,
					"Error: "+err.Error())
				continue
			}
			rootParent := (*rootItem).GetParent()
			rootPath := (*rootParent).GetPathString()
			fileGraphItem, err := graph.GetItemByRelativePath(rootPath, relativeProjectPath)
			if err != nil {
				logging.LogInfoMultiline("Could not find relative project path. Ignoring.",
					"Project Name: "+project.Name,
					"Path: "+relativeProjectPath,
					"Error: "+err.Error())
				continue
			} else {
				pathOfItem := (*fileGraphItem).GetPathString()
				project.AbsoluteProjectReferencePaths = append(project.AbsoluteProjectReferencePaths, pathOfItem)
			}
		}
	}
}

func getSolutionList(solutionPaths []string,
	metadata repomodel.RepositoryMetadata,
	repoClient repoclient.RepoClient,
	msBuildClient msbuildclient.MsBuildClient,
	f fileutil.FileGraph) []*projectmodel.MsBuildSolution {
	var solutions []*projectmodel.MsBuildSolution
	for _, solutionPath := range solutionPaths {
		solution := getSolutionFromSourceControl(metadata, solutionPath, repoClient, msBuildClient)
		solution.AbsolutePath = solutionPath
		solutions = append(solutions, solution)
		prettyPaths, err := f.PrettifyPaths(solution.RelativeProjectPaths)
		if err != nil {
			panic(err)
		}
		solution.AbsoluteProjectPaths = prettyPaths
	}
	return solutions
}

func getProjectList(projectPaths []string,
	metadata repomodel.RepositoryMetadata,
	repoClient repoclient.RepoClient,
	msBuildClient msbuildclient.MsBuildClient) []*projectmodel.MsBuildProject {
	var projects []*projectmodel.MsBuildProject
	for _, projectPath := range projectPaths {
		project := getProjectFromSourceControl(metadata, projectPath, repoClient, msBuildClient)
		project.AbsolutePath = projectPath
		projects = append(projects, project)
	}
	return projects
}

func getPublishProfileList(publishProfilePaths []string,
	metadata repomodel.RepositoryMetadata,
	repoClient repoclient.RepoClient,
	msBuildClient msbuildclient.MsBuildClient) []*projectmodel.MsBuildPublishProfile {
	var publishProfiles []*projectmodel.MsBuildPublishProfile
	for _, publishProfilePath := range publishProfilePaths {
		publishProfile := getPublishProfileFromSourceControl(metadata, publishProfilePath, repoClient, msBuildClient)
		publishProfiles = append(publishProfiles, publishProfile)
	}
	return publishProfiles
}

// TODO: Implement in processing loop
func getPublishProfileFromSourceControl(
	metadata repomodel.RepositoryMetadata,
	path string,
	repoClient repoclient.RepoClient,
	msBuildClient msbuildclient.MsBuildClient) *projectmodel.MsBuildPublishProfile {
	repoMetadata := getRepositoryFileMetadataFromPath(metadata, path)

	payload, err := repoClient.GetFile(repoMetadata)
	if err != nil {
		panic(err)
	}

	publishProfile, err := msBuildClient.GetPublishProfile(*payload)
	if err != nil {
		panic(err)
	}

	return publishProfile
}

func getProjectFromSourceControl(
	metadata repomodel.RepositoryMetadata,
	path string,
	repoClient repoclient.RepoClient,
	msBuildClient msbuildclient.MsBuildClient) *projectmodel.MsBuildProject {
	repoMetadata := getRepositoryFileMetadataFromPath(metadata, path)

	payload, err := repoClient.GetFile(repoMetadata)
	if err != nil {
		panic(err)
	}

	project, err := msBuildClient.GetProject(*payload)
	if err != nil {
		panic(err)
	}

	return project
}

func getSolutionFromSourceControl(
	metadata repomodel.RepositoryMetadata,
	path string,
	repoClient repoclient.RepoClient,
	msBuildClient msbuildclient.MsBuildClient) *projectmodel.MsBuildSolution {
	repoMetadata := getRepositoryFileMetadataFromPath(metadata, path)

	payload, err := repoClient.GetFile(repoMetadata)
	if err != nil {
		panic(err)
	}

	solution, err := msBuildClient.GetSolution(*payload)
	if err != nil {
		panic(err)
	}

	return solution
}

func getRepositoryFileMetadataFromPath(metadata repomodel.RepositoryMetadata,
	path string) repomodel.RepositoryFileMetadata {
	fileMetadata := filesysmodel.FileSystemMetadata{
		Path: path,
		Type: filesysmodel.FileType,
	}
	return repomodel.NewRepositoryFileMetadata(metadata.Name, metadata.Branch, path, fileMetadata)
}

func getSolutionPaths(metadata repomodel.RepositoryMetadata) ([]string, error) {
	return metadata.GetMatchingFiles(ValidSolutionExtensions)
}

func getProjectPaths(metadata repomodel.RepositoryMetadata) ([]string, error) {
	return metadata.GetMatchingFiles(ValidProjectExtensions)
}

func getPublishProfilePaths(metadata repomodel.RepositoryMetadata) ([]string, error) {
	return metadata.GetMatchingFiles(ValidPublishProfileExtensions)
}
