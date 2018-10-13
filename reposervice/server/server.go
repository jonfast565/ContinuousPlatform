package server

import (
	"../../constants"
	"../../logging"
	"../../models"
	"../../pathutil"
	"../../webutil"
	"errors"
	"fmt"
	"github.com/ahmetb/go-linq"
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

func (e TeamServicesEndpoint) GetRepositories() (*models.RepositoryPackage, error) {
	results := make([]models.RepositoryMetadata, 0)
	resultsChannel := make(chan models.RepositoryMetadata)
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
			logging.LogInfo("No more repositories/branches received")
			noMore = true
		}
		if noMore {
			break
		}
	}

	amalgamation := models.RepositoryPackage{
		Metadata: results,
		Type:     models.AzureDevOps,
	}
	return &amalgamation, nil
}

func (e *TeamServicesEndpoint) getRepositoryBranches(
	repository models.TsGitRepositoryModel,
	wg *sync.WaitGroup,
	resultsChannel chan models.RepositoryMetadata) {
	defer wg.Done()
	branches, err := e.getBranchInformation(repository)
	if err != nil {
		logging.LogFatal(err.Error())
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
	repository models.TsGitRepositoryModel,
	branch models.TsGitRefsModel,
	wg *sync.WaitGroup,
	resultsChannel chan models.RepositoryMetadata) {
	defer wg.Done()
	files, err := e.getBranchFileList(repository, branch)
	if err != nil {
		logging.LogFatal(err.Error())
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
	repository models.TsGitRepositoryModel,
	branch models.TsGitRefsModel,
	files *models.TsGitFileList) models.RepositoryMetadata {
	return models.RepositoryMetadata{
		Name:        repository.Name,
		Branch:      getCleanBranchName(branch),
		OptionalUrl: repository.RemoteUrl,
		Files:       *getFileSystemMetadataFromList(*files),
	}
}

func (e TeamServicesEndpoint) GetFile(file models.RepositoryFileMetadata) (*models.FilePayload, error) {
	logging.LogInfo(fmt.Sprintf("Downloading file: %s", file.File.Path))

	if len(file.File.Path) == 0 {
		return nil, errors.New("file path not specified")
	}
	filePath := file.File.Path[1:len(file.File.Path)]

	repoInfo, err := e.getRepositoryInformation()
	if err != nil {
		return nil, err
	}
	repo := linq.From(repoInfo.Value).FirstWithT(func(r models.TsGitRepositoryModel) bool {
		return r.Name == file.Repo
	})

	if repo == nil {
		return nil, errors.New(fmt.Sprintf("repository not found: '%s'", file.Repo))
	}

	repoTyped := repo.(models.TsGitRepositoryModel)
	branches, err := e.getBranchInformation(repoTyped)
	if err != nil {
		return nil, err
	}

	branch := linq.From(branches.Value).FirstWithT(func(ref models.TsGitRefsModel) bool {
		return getCleanBranchName(ref) == file.Branch
	})

	if branch == nil {
		return nil, errors.New(fmt.Sprintf("branch not found: %s", file.Branch))
	}

	branchTyped := branch.(models.TsGitRefsModel)
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

	logging.LogInfo("File downloaded!")
	return fileValue, nil
}

func (e TeamServicesEndpoint) getRepositoryInformation() (*models.TsGitRepositoryList, error) {
	url := constants.GetRepositoryApiPath(
		e.Configuration.CollectionUrl,
		e.Configuration.ProjectName)
	request, err := http.NewRequest(constants.GetMethod, url, nil)
	if err != nil {
		return nil, err
	}
	buildTeamServiceAuthHeader(request, e)
	webutil.AddJsonHeader(request)

	var result models.TsGitRepositoryList
	err = webutil.ExecuteRequestAndReadJsonBody(&e.Client, request, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func buildTeamServiceAuthHeader(request *http.Request, e TeamServicesEndpoint) {
	request.SetBasicAuth(e.Configuration.Username, e.Configuration.PersonalAccessToken)
}

func (e TeamServicesEndpoint) getFileInformation(
	repository models.TsGitRepositoryModel,
	branch models.TsGitRefsModel,
	path string) (*models.FilePayload, error) {
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

	var result models.FilePayload
	var resultBytes *[]byte
	resultBytes, err = webutil.ExecuteRequestAndReadBinaryBody(&e.Client, request)
	if err != nil {
		return nil, err
	}

	result = models.FilePayload{
		Name:  pathutil.GetLastPathComponent("." + path),
		Bytes: *resultBytes,
	}
	return &result, nil
}

func (e TeamServicesEndpoint) getBranchInformation(
	repository models.TsGitRepositoryModel) (*models.TsGitRefsList, error) {
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

	var result models.TsGitRefsList
	err = webutil.ExecuteRequestAndReadJsonBody(&e.Client, request, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (e TeamServicesEndpoint) getBranchFileList(
	repository models.TsGitRepositoryModel,
	branch models.TsGitRefsModel) (*models.TsGitFileList, error) {
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

	var result models.TsGitFileList
	err = webutil.ExecuteRequestAndReadJsonBody(&e.Client, request, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func getFileSystemMetadataFromList(
	fileList models.TsGitFileList) *[]models.FileSystemMetadata {
	result := make([]models.FileSystemMetadata, 0)
	for _, file := range fileList.Value {
		result = append(result, models.FileSystemMetadata{
			Path:             "." + file.Path,
			OptionalChangeId: file.CommitId,
			Type:             getGitObjectType(file.GitObjectType),
		})
	}

	return &result
}

func isValidBranch(branch models.TsGitRefsModel) bool {
	return strings.Contains(branch.Name, constants.RefsHeadsConstants)
}

func getCleanBranchName(branch models.TsGitRefsModel) string {
	return strings.Replace(branch.Name, constants.RefsHeadsConstants, "", -1)
}

func getGitObjectType(objectType string) models.FileSystemObjectType {
	if objectType == constants.BlobConstant {
		return models.FileType
	}
	return models.FolderType
}
