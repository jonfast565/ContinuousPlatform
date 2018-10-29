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
	for _, solutionPath := range solutionPaths {
		solution := getSolutionFromSourceControl(metadata, solutionPath, repoClient, msBuildClient)
		solutions = append(solutions, *solution)
	}

	var projects []projectmodel.MsBuildProject
	for _, projectPath := range projectPaths {
		project := getProjectFromSourceControl(metadata, projectPath, repoClient, msBuildClient)
		projects = append(projects, *project)
	}

	var publishProfiles []projectmodel.MsBuildPublishProfile
	for _, publishProfilePath := range publishProfilePaths {
		publishProfile := getPublishProfileFromSourceControl(metadata, publishProfilePath, repoClient, msBuildClient)
		publishProfiles = append(publishProfiles, *publishProfile)
	}

	logging.LogInfo("Done building .NET deliverables for " + metadata.Name + " b. " + metadata.Branch)
	return results, nil
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
