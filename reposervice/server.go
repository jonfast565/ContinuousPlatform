package main

import (
	"../models/filesystem"
	"../models/repos"
	"../models/repos/teamservices"
	"../models/web"
	"../utilities"
	"encoding/json"
	"net/http"
	"strings"
)

type TeamServicesConfiguration struct {
	Port                int    `json:"port"`
	CollectionUrl       string `json:"collectionUrl"`
	ProjectName         string `json:"projectName"`
	Username            string `json:"username"`
	PersonalAccessToken string `json:"personalAccessToken"`
}

type TeamServicesEndpoint struct {
	Client        http.Client
	Configuration TeamServicesConfiguration
}

func NewTeamServicesEndpoint(configuration TeamServicesConfiguration) *TeamServicesEndpoint {
	result := new(TeamServicesEndpoint)
	result.Client = http.Client{}
	result.Configuration = configuration
	return result
}

func (e TeamServicesEndpoint) GetRepositories() (*repos.RepositoryPackage, error) {
	results := make([]repos.RepositoryMetadata, 0)
	resultsChan := make(chan repos.RepositoryMetadata)
	repositories, err := e.getRepositoryInformation()
	if err != nil {
		return nil, err
	}
	refs := make(chan teamservices.TeamServicesGitRefsList, repositories.Count)
	for _, repository := range repositories.Value {
		go func(repository teamservices.TeamServicesGitRepositoryModel) {
			branches, err := e.getBranchInformation(repository)
			if err != nil {
				panic("Branch information not retrieved")
			}
			refs <- *branches
		}(repository)
	}
	for _, repository := range repositories.Value {
		var ref = <-refs
		for _, branch := range ref.Value {
			go func(branch teamservices.TeamServicesGitRefsModel) {
				files, err := e.getBranchFileList(repository, branch)
				if err != nil {
					panic("Files not retrieved")
				}
				result := repos.RepositoryMetadata{
					Name:     repository.Name,
					Branch:   branch.Name,
					Url:      repository.RemoteUrl,
					Metadata: getFileSystemMetadataFromList(*files),
				}
				resultsChan <- result
			}(branch)
		}
	}
	select {
	case result, ok := <-resultsChan:
		if ok {
			results = append(results, result)
		}
	}
	amalgamation := repos.RepositoryPackage{
		Metadata: results,
		Type:     repos.AzureDevOps,
	}
	return &amalgamation, nil
}

func (e TeamServicesEndpoint) GetFile(file repos.RepositoryFileMetadata) (*web.FilePayload, error) {
	// TODO: Implement this
	return nil, nil
}

func (e TeamServicesEndpoint) getRepositoryInformation() (*teamservices.TeamServicesGitRepositoryList, error) {
	url := teamservices.GetRepositoryApiPath(e.Configuration.CollectionUrl, e.Configuration.ProjectName)
	request, err := http.NewRequest(utilities.GetMethod, url, nil)
	if err != nil {
		return nil, err
	}
	buildTeamServiceAuthHeader(request, e)
	utilities.AddJsonHeader(request)
	body, err := utilities.ExecuteRequestAndReadBodyAsString(&e.Client, request)
	var result teamservices.TeamServicesGitRepositoryList
	if err := json.Unmarshal(*body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func buildTeamServiceAuthHeader(request *http.Request, e TeamServicesEndpoint) {
	request.Header.Add(utilities.AuthorizationHeader, utilities.BasicAuthHeaderValue(
		teamservices.BuildAuthorizationHeader(e.Configuration.Username, e.Configuration.PersonalAccessToken)))
}

func (e TeamServicesEndpoint) getFileInformation(
	repository teamservices.TeamServicesGitRepositoryModel,
	branch teamservices.TeamServicesGitRefsModel,
	path string) (*web.FilePayload, error) {
	url := teamservices.GetApiFilesPath(
		e.Configuration.CollectionUrl,
		e.Configuration.ProjectName,
		repository.Id,
		strings.Replace(branch.Name, teamservices.RefsHeadsConstants, "", -1),
		path)
	request, err := http.NewRequest(utilities.GetMethod, url, nil)
	if err != nil {
		return nil, err
	}
	buildTeamServiceAuthHeader(request, e)
	utilities.AddOctetHeader(request)
	body, err := utilities.ExecuteRequestAndReadBodyAsString(&e.Client, request)
	var result web.FilePayload
	if err := json.Unmarshal(*body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (e TeamServicesEndpoint) getBranchInformation(
	repository teamservices.TeamServicesGitRepositoryModel) (*teamservices.TeamServicesGitRefsList, error) {
	url := teamservices.GetBranchApiPath(e.Configuration.CollectionUrl, e.Configuration.ProjectName, repository.Id)
	request, err := http.NewRequest(utilities.GetMethod, url, nil)
	if err != nil {
		return nil, err
	}
	buildTeamServiceAuthHeader(request, e)
	utilities.AddJsonHeader(request)
	body, err := utilities.ExecuteRequestAndReadBodyAsString(&e.Client, request)
	var result teamservices.TeamServicesGitRefsList
	if err := json.Unmarshal(*body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (e TeamServicesEndpoint) getBranchFileList(
	repository teamservices.TeamServicesGitRepositoryModel,
	branch teamservices.TeamServicesGitRefsModel) (*teamservices.TeamServicesGitFileList, error) {
	url := teamservices.GetApiFilesPath(
		e.Configuration.CollectionUrl,
		e.Configuration.ProjectName,
		repository.Id,
		branch.Name,
		"")
	request, err := http.NewRequest(utilities.GetMethod, url, nil)
	if err != nil {
		return nil, err
	}
	buildTeamServiceAuthHeader(request, e)
	utilities.AddJsonHeader(request)
	body, err := utilities.ExecuteRequestAndReadBodyAsString(&e.Client, request)
	var result teamservices.TeamServicesGitFileList
	if err := json.Unmarshal(*body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func getFileSystemMetadataFromList(fileList teamservices.TeamServicesGitFileList) []filesystem.FileSystemMetadata {
	result := make([]filesystem.FileSystemMetadata, 0)
	for _, file := range fileList.Value {
		result = append(result, filesystem.FileSystemMetadata{
			Path:             "." + file.Path,
			OptionalChangeId: file.CommitId,
			Type:             getGitObjectType(file.GitObjectType),
		})
	}
	return result
}

func getGitObjectType(objectType string) filesystem.FileSystemObjectType {
	if objectType == teamservices.BlobConstant {
		return filesystem.FileType
	}
	return filesystem.FolderType
}
