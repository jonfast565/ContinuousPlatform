// repo service server
package server

import (
	"errors"
	"fmt"
	"github.com/ahmetb/go-linq"
	"github.com/jonfast565/continuous-platform/constants"
	"github.com/jonfast565/continuous-platform/models/filesysmodel"
	"github.com/jonfast565/continuous-platform/models/miscmodel"
	"github.com/jonfast565/continuous-platform/models/repomodel"
	"github.com/jonfast565/continuous-platform/utilities/limitutil"
	"github.com/jonfast565/continuous-platform/utilities/logging"
	"github.com/jonfast565/continuous-platform/utilities/pathutil"
	"github.com/jonfast565/continuous-platform/utilities/webutil"
	"net/http"
	"strings"
	"sync"
)

// Configuration for a Team Services (aka. Azure Devops) connection
type TeamServicesConfiguration struct {
	Port                int    `json:"port"`
	CollectionUrl       string `json:"collectionUrl"`
	ProjectName         string `json:"projectName"`
	Username            string `json:"username"`
	PersonalAccessToken string `json:"personalAccessToken"`
}

// Endpoint for Team Services
type TeamServicesEndpoint struct {
	Configuration TeamServicesConfiguration
}

// Endpoint constructor
func NewTeamServicesEndpoint(configuration TeamServicesConfiguration) *TeamServicesEndpoint {
	result := new(TeamServicesEndpoint)
	result.Configuration = configuration
	return result
}

// Gets a package object containing a list of repositories, suitable for xfer over the wire.
// If failed, returns an error.
func (e TeamServicesEndpoint) GetRepositories() (*repomodel.RepositoryPackage, error) {
	results := make([]repomodel.RepositoryMetadata, 0)
	resultsChannel := make(chan repomodel.RepositoryMetadata)
	rateLimiter := limitutil.NewRateLimiter(2, 2)
	repositories, err := e.getRepositoryInformation()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(repositories.Value))
	for _, repository := range repositories.Value {
		go e.getRepositoryBranches(repository, &wg, resultsChannel)
		limitutil.WaitForAllow(rateLimiter)
	}
	wg.Wait()

	for {
		noMore := false
		select {
		case msg := <-resultsChannel:
			results = append(results, msg)
		default:
			logging.LogInfo("No more repositories/branches received")
			noMore = true
		}
		if noMore {
			break
		}
	}

	amalgamation := repomodel.RepositoryPackage{
		Metadata: results,
		Type:     repomodel.AzureDevOps,
	}
	return &amalgamation, nil
}

func (e *TeamServicesEndpoint) getRepositoryBranches(
	repository repomodel.TsGitRepositoryModel,
	wg *sync.WaitGroup,
	resultsChannel chan repomodel.RepositoryMetadata) {
	defer wg.Done()
	branches, err := e.getBranchInformation(repository)

	if err != nil {
		logging.LogFatal(err)
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
}

func (e *TeamServicesEndpoint) getRepositoryBranchFiles(
	repository repomodel.TsGitRepositoryModel,
	branch repomodel.TsGitRefsModel,
	wg *sync.WaitGroup,
	resultsChannel chan repomodel.RepositoryMetadata) {
	defer wg.Done()
	files, err := e.getBranchFileList(repository, branch)

	if err != nil {
		logging.LogFatal(err)
		panic("Files not retrieved")
	}

	result := e.buildRepositoryMetadata(repository, branch, files)
	logging.LogInfoMultiline("Repository metadata built: ",
		fmt.Sprintf("Repo: %s", result.Name),
		fmt.Sprintf("Branch: %s", result.Branch),
		fmt.Sprintf("Url: %s", result.OptionalUrl),
	)

	go func() { resultsChannel <- result }()
}

func (e TeamServicesEndpoint) buildRepositoryMetadata(
	repository repomodel.TsGitRepositoryModel,
	branch repomodel.TsGitRefsModel,
	files *repomodel.TsGitFileList) repomodel.RepositoryMetadata {
	return repomodel.RepositoryMetadata{
		Name:        repository.Name,
		Branch:      getCleanBranchName(branch),
		OptionalUrl: repository.RemoteUrl,
		Files:       *getFileSystemMetadataFromList(*files),
	}
}

// Gets a file by providing metadata about its repository and path. Returns the payload as an object containing
// a bytestring. If failed, returns an error.
func (e TeamServicesEndpoint) GetFile(file repomodel.RepositoryFileMetadata) (*miscmodel.FilePayload, error) {

	if len(file.File.Path) == 0 {
		return nil, errors.New("file path not specified")
	}
	filePath := file.File.Path[1:len(file.File.Path)]

	repoInfo, err := e.getRepositoryInformation()
	if err != nil {
		return nil, err
	}
	repo := linq.From(repoInfo.Value).FirstWithT(func(r repomodel.TsGitRepositoryModel) bool {
		return r.Name == file.Repo
	})

	if repo == nil {
		return nil, errors.New(fmt.Sprintf("repository not found: '%s'", file.Repo))
	}

	repoTyped := repo.(repomodel.TsGitRepositoryModel)
	branches, err := e.getBranchInformation(repoTyped)
	if err != nil {
		return nil, err
	}

	branch := linq.From(branches.Value).FirstWithT(func(ref repomodel.TsGitRefsModel) bool {
		return getCleanBranchName(ref) == file.Branch
	})

	if branch == nil {
		return nil, errors.New(fmt.Sprintf("branch not found: %s", file.Branch))
	}

	branchTyped := branch.(repomodel.TsGitRefsModel)
	logging.LogInfo(fmt.Sprintf("Getting file %s from %s b. %s",
		file.Name,
		file.Repo,
		file.Branch))

	fileValue, err := e.getFileInformation(repoTyped, branchTyped, filePath)
	if err != nil {
		return nil, err
	}

	if fileValue.Bytes == nil {
		return nil, errors.New(fmt.Sprintf("file not found: %s", file.Name))
	}

	logging.LogInfoMultiline("File downloaded!",
		fmt.Sprintf("Path: %s", file.File.Path))
	return fileValue, nil
}

func (e TeamServicesEndpoint) getRepositoryInformation() (*repomodel.TsGitRepositoryList, error) {
	url := constants.GetRepositoryApiPath(
		e.Configuration.CollectionUrl,
		e.Configuration.ProjectName)
	request, err := http.NewRequest(constants.GetMethod, url, nil)
	if err != nil {
		return nil, err
	}
	buildTeamServiceAuthHeader(request, e)
	webutil.AddJsonHeader(request)

	var result repomodel.TsGitRepositoryList
	err = webutil.ExecuteRequestAndReadJsonBody(request, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func buildTeamServiceAuthHeader(request *http.Request, e TeamServicesEndpoint) {
	request.SetBasicAuth(e.Configuration.Username, e.Configuration.PersonalAccessToken)
}

func (e TeamServicesEndpoint) getFileInformation(
	repository repomodel.TsGitRepositoryModel,
	branch repomodel.TsGitRefsModel,
	path string) (*miscmodel.FilePayload, error) {
	url := constants.GetApiFilesPath(
		e.Configuration.CollectionUrl,
		e.Configuration.ProjectName,
		repository.Id,
		getCleanBranchName(branch),
		path)

	request, err := http.NewRequest(constants.GetMethod, url, nil)
	if err != nil {
		return nil, err
	}

	buildTeamServiceAuthHeader(request, e)
	webutil.AddOctetHeader(request)

	var result miscmodel.FilePayload
	var resultBytes *[]byte
	resultBytes, err = webutil.ExecuteRequestAndReadBinaryBody(request)
	if err != nil {
		return nil, err
	}

	result = miscmodel.FilePayload{
		Name:  pathutil.GetLastPathComponent("." + path),
		Bytes: *resultBytes,
	}
	return &result, nil
}

func (e TeamServicesEndpoint) getBranchInformation(
	repository repomodel.TsGitRepositoryModel) (*repomodel.TsGitRefsList, error) {
	url := constants.GetBranchApiPath(
		e.Configuration.CollectionUrl,
		e.Configuration.ProjectName,
		repository.Id)
	request, err := http.NewRequest(constants.GetMethod, url, nil)
	if err != nil {
		return nil, err
	}

	buildTeamServiceAuthHeader(request, e)
	webutil.AddJsonHeader(request)

	var result repomodel.TsGitRefsList
	err = webutil.ExecuteRequestAndReadJsonBody(request, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (e TeamServicesEndpoint) getBranchFileList(
	repository repomodel.TsGitRepositoryModel,
	branch repomodel.TsGitRefsModel) (*repomodel.TsGitFileList, error) {
	url := constants.GetApiFilesPath(
		e.Configuration.CollectionUrl,
		e.Configuration.ProjectName,
		repository.Id,
		getCleanBranchName(branch),
		"")

	request, err := http.NewRequest(constants.GetMethod, url, nil)
	if err != nil {
		return nil, err
	}

	buildTeamServiceAuthHeader(request, e)
	webutil.AddJsonHeader(request)

	var result repomodel.TsGitFileList
	err = webutil.ExecuteRequestAndReadJsonBody(request, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func getFileSystemMetadataFromList(
	fileList repomodel.TsGitFileList) *[]filesysmodel.FileSystemMetadata {
	result := make([]filesysmodel.FileSystemMetadata, 0)
	for _, file := range fileList.Value {
		result = append(result, filesysmodel.FileSystemMetadata{
			Path:             "." + file.Path,
			OptionalChangeId: file.CommitId,
			Type:             getGitObjectType(file.GitObjectType),
		})
	}

	return &result
}

func isValidBranch(branch repomodel.TsGitRefsModel) bool {
	// TODO: Remove master restriction
	return strings.Contains(branch.Name, constants.RefsHeadsConstants) && strings.Contains(branch.Name, "master")
}

func getCleanBranchName(branch repomodel.TsGitRefsModel) string {
	return strings.Replace(branch.Name, constants.RefsHeadsConstants, "", -1)
}

func getGitObjectType(objectType string) filesysmodel.FileSystemObjectType {
	if objectType == constants.BlobConstant {
		return filesysmodel.FileType
	}
	return filesysmodel.FolderType
}
