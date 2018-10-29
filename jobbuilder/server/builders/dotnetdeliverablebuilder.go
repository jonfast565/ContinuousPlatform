package builders

import (
	"../../../clients/msbuildclient"
	"../../../clients/repoclient"
	"../../../fileutil"
	"../../../logging"
	"../../../models/projectmodel"
	"../../../models/repomodel"
)

var ValidProjectExtensions = []string{
	"^\\.csproj$",
	"^\\.fsproj$",
	"^\\.vbproj$",
}

var ValidSolutionExtensions = []string{
	"^\\.sln$",
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

	for _, solutionPath := range solutionPaths {
		getSolutionFromSourceControl(metadata, solutionPath, repoClient, msBuildClient)
	}

	for _, projectPath := range projectPaths {
		getProjectFromSourceControl(metadata, projectPath, repoClient, msBuildClient)
	}

	logging.LogInfo("Done building .NET deliverables for " + metadata.Name + " b. " + metadata.Branch)
	return results, nil
}

func getProjectFromSourceControl(
	metadata repomodel.RepositoryMetadata,
	path string,
	repoClient repoclient.RepoClient,
	msBuildClient msbuildclient.MsBuildClient) {
	repoMetadata := getRepositoryFileMetadataFromPath(metadata, path)
	payload, err := repoClient.GetFile(repoMetadata)
	if err != nil {
		panic(err)
	}
	project := msBuildClient
}

func getRepositoryFileMetadataFromPath(metadata repomodel.RepositoryMetadata,
	path string) repomodel.RepositoryFileMetadata {
	return repomodel.NewRepositoryFileMetadata(metadata.Name, metadata.Branch, path)
}

func getSolutionFromSourceControl(
	metadata repomodel.RepositoryMetadata,
	path string,
	repoClient repoclient.RepoClient,
	msBuildClient msbuildclient.MsBuildClient) {
	repoMetadata := getRepositoryFileMetadataFromPath(metadata, path)
	payload, err := repoClient.GetFile(repoMetadata)
	if err != nil {
		panic(err)
	}
	solution := msBuildClient
}

func getSolutionPaths(metadata repomodel.RepositoryMetadata) ([]string, error) {
	return metadata.GetMatchingFiles(ValidSolutionExtensions)
}

func getProjectPaths(metadata repomodel.RepositoryMetadata) ([]string, error) {
	return metadata.GetMatchingFiles(ValidProjectExtensions)
}
