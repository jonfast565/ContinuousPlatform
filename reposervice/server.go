package main

import (
	"../models/filesystem"
	"../models/repos"
	"../models/repos/teamservices"
	"../models/web"
	"../utilities"
	"log"
	"net/http"
	"strings"
	"sync"
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
	resultsChannel := make(chan repos.RepositoryMetadata)
	repositories, err := e.getRepositoryInformation()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(repositories.Value))
	for _, repository := range repositories.Value {
		go e.getRepositoryBranches(repository, &wg, resultsChannel)
	}
	wg.Wait()

	for {
		noMore := false
		select {
		case msg := <-resultsChannel:
			results = append(results, msg)
		default:
			log.Print("No more repositories/branches received")
			noMore = true
		}
		if noMore {
			break
		}
	}

	amalgamation := repos.RepositoryPackage{
		Metadata: results,
		Type:     repos.AzureDevOps,
	}
	return &amalgamation, nil
}

func (e *TeamServicesEndpoint) getRepositoryBranches(
	repository teamservices.TsGitRepositoryModel,
	wg *sync.WaitGroup,
	resultsChannel chan repos.RepositoryMetadata) {
	branches, err := e.getBranchInformation(repository)
	if err != nil {
		panic("Branch information not retrieved")
	}
	wg.Add(len(branches.Value))
	for _, branch := range branches.Value {
		if !isValidBranch(branch) {
			wg.Done()
			continue
		}
		go e.getRepositoryBranchFiles(repository, branch, wg, resultsChannel)
	}
	wg.Done()
}

func (e *TeamServicesEndpoint) getRepositoryBranchFiles(
	repository teamservices.TsGitRepositoryModel,
	branch teamservices.TsGitRefsModel,
	wg *sync.WaitGroup,
	resultsChannel chan repos.RepositoryMetadata) {
	files, err := e.getBranchFileList(repository, branch)
	if err != nil {
		panic("Files not retrieved")
	}
	result := e.buildRepositoryMetadata(repository, branch, files)
	log.Printf("Result built:\n%s", result)
	go func() { resultsChannel <- result }()
	wg.Done()
}

func (e TeamServicesEndpoint) buildRepositoryMetadata(
	repository teamservices.TsGitRepositoryModel,
	branch teamservices.TsGitRefsModel,
	files *teamservices.TsGitFileList) repos.RepositoryMetadata {
	return repos.RepositoryMetadata{
		Name:     repository.Name,
		Branch:   getCleanBranchName(branch),
		Url:      repository.RemoteUrl,
		Metadata: *getFileSystemMetadataFromList(*files),
	}
}

func (e TeamServicesEndpoint) GetFile(file repos.RepositoryFileMetadata) (*web.FilePayload, error) {
	// TODO: Implement this
	return nil, nil
}

func (e TeamServicesEndpoint) getRepositoryInformation() (*teamservices.TsGitRepositoryList, error) {
	url := teamservices.GetRepositoryApiPath(
		e.Configuration.CollectionUrl,
		e.Configuration.ProjectName)
	request, err := http.NewRequest(utilities.GetMethod, url, nil)
	if err != nil {
		return nil, err
	}
	buildTeamServiceAuthHeader(request, e)
	utilities.AddJsonHeader(request)
	var result teamservices.TsGitRepositoryList
	err = utilities.ExecuteRequestAndReadJsonBody(&e.Client, request, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func buildTeamServiceAuthHeader(request *http.Request, e TeamServicesEndpoint) {
	request.SetBasicAuth(e.Configuration.Username, e.Configuration.PersonalAccessToken)
}

func (e TeamServicesEndpoint) getFileInformation(
	repository teamservices.TsGitRepositoryModel,
	branch teamservices.TsGitRefsModel,
	path string) (*web.FilePayload, error) {
	url := teamservices.GetApiFilesPath(
		e.Configuration.CollectionUrl,
		e.Configuration.ProjectName,
		repository.Id,
		getCleanBranchName(branch),
		path)
	request, err := http.NewRequest(utilities.GetMethod, url, nil)
	if err != nil {
		return nil, err
	}
	buildTeamServiceAuthHeader(request, e)
	utilities.AddOctetHeader(request)
	var result web.FilePayload
	err = utilities.ExecuteRequestAndReadJsonBody(&e.Client, request, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (e TeamServicesEndpoint) getBranchInformation(
	repository teamservices.TsGitRepositoryModel) (*teamservices.TsGitRefsList, error) {
	url := teamservices.GetBranchApiPath(
		e.Configuration.CollectionUrl,
		e.Configuration.ProjectName,
		repository.Id)
	request, err := http.NewRequest(utilities.GetMethod, url, nil)
	if err != nil {
		return nil, err
	}
	buildTeamServiceAuthHeader(request, e)
	utilities.AddJsonHeader(request)
	var result teamservices.TsGitRefsList
	err = utilities.ExecuteRequestAndReadJsonBody(&e.Client, request, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (e TeamServicesEndpoint) getBranchFileList(
	repository teamservices.TsGitRepositoryModel,
	branch teamservices.TsGitRefsModel) (*teamservices.TsGitFileList, error) {
	url := teamservices.GetApiFilesPath(
		e.Configuration.CollectionUrl,
		e.Configuration.ProjectName,
		repository.Id,
		getCleanBranchName(branch),
		"")
	request, err := http.NewRequest(utilities.GetMethod, url, nil)
	if err != nil {
		return nil, err
	}
	buildTeamServiceAuthHeader(request, e)
	utilities.AddJsonHeader(request)
	var result teamservices.TsGitFileList
	err = utilities.ExecuteRequestAndReadJsonBody(&e.Client, request, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func getFileSystemMetadataFromList(
	fileList teamservices.TsGitFileList) *[]filesystem.FileSystemMetadata {
	result := make([]filesystem.FileSystemMetadata, 0)
	for _, file := range fileList.Value {
		result = append(result, filesystem.FileSystemMetadata{
			Path:             "." + file.Path,
			OptionalChangeId: file.CommitId,
			Type:             getGitObjectType(file.GitObjectType),
		})
	}
	return &result
}

func isValidBranch(branch teamservices.TsGitRefsModel) bool {
	return strings.Contains(branch.Name, teamservices.RefsHeadsConstants)
}

func getCleanBranchName(branch teamservices.TsGitRefsModel) string {
	return strings.Replace(branch.Name, teamservices.RefsHeadsConstants, "", -1)
}

func getGitObjectType(objectType string) filesystem.FileSystemObjectType {
	if objectType == teamservices.BlobConstant {
		return filesystem.FileType
	}
	return filesystem.FolderType
}
