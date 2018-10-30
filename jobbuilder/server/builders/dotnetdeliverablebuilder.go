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

	publishProfilePaths, err := getPublishProfilePaths(metadata)
	if err != nil {
		return nil, err
	}

	var solutions []projectmodel.MsBuildSolution
	getSolutionList(solutionPaths, metadata, repoClient, msBuildClient, solutions)

	var projects []projectmodel.MsBuildProject
	getProjectList(projectPaths, metadata, repoClient, msBuildClient, projects)

	var publishProfiles []projectmodel.MsBuildPublishProfile
	getPublishProfileList(publishProfilePaths, metadata, repoClient, msBuildClient, publishProfiles)

	populateAbsoluteProjectPaths(projects, fileGraph)

	for _, project := range projects {
		resolveProjectDependencies(&project, projects)
	}

	for _, solution := range solutions {
		linkProjectSolutions(&solution, projects)
	}

	logging.LogInfo("Done building .NET deliverables for " + metadata.Name + " b. " + metadata.Branch)
	return results, nil
}

func resolveProjectDependencies(project *projectmodel.MsBuildProject, projects []projectmodel.MsBuildProject) {
	// TODO: Implement this bitch
}

func linkProjectSolutions(solution *projectmodel.MsBuildSolution, projects []projectmodel.MsBuildProject) {
	// TODO: Implement this bitch
}

func populateAbsoluteProjectPaths(projects []projectmodel.MsBuildProject, graph fileutil.FileGraph) {
	for _, project := range projects {
		for _, relativeProjectPath := range project.RelativeProjectReferencePaths {
			fileGraphItem, err := graph.GetItemByRelativePath(project.AbsolutePath, relativeProjectPath)
			if err != nil {
				logging.LogInfoMultiline("Could not find project path. Ignoring.",
					"Project Name: "+project.Name,
					"Path: "+relativeProjectPath,
					"Error: "+err.Error())
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
	solutions []projectmodel.MsBuildSolution) {
	for _, solutionPath := range solutionPaths {
		solution := getSolutionFromSourceControl(metadata, solutionPath, repoClient, msBuildClient)
		solutions = append(solutions, *solution)
	}
}

func getProjectList(projectPaths []string,
	metadata repomodel.RepositoryMetadata,
	repoClient repoclient.RepoClient,
	msBuildClient msbuildclient.MsBuildClient,
	projects []projectmodel.MsBuildProject) {
	for _, projectPath := range projectPaths {
		project := getProjectFromSourceControl(metadata, projectPath, repoClient, msBuildClient)
		projects = append(projects, *project)
	}
}

func getPublishProfileList(publishProfilePaths []string,
	metadata repomodel.RepositoryMetadata,
	repoClient repoclient.RepoClient,
	msBuildClient msbuildclient.MsBuildClient,
	publishProfiles []projectmodel.MsBuildPublishProfile) {
	for _, publishProfilePath := range publishProfilePaths {
		publishProfile := getPublishProfileFromSourceControl(metadata, publishProfilePath, repoClient, msBuildClient)
		publishProfiles = append(publishProfiles, *publishProfile)
	}
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
