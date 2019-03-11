package repoclient

import (
	"bytes"
	"encoding/json"
	"github.com/jonfast565/continuous-platform/constants"
	"github.com/jonfast565/continuous-platform/models/miscmodel"
	"github.com/jonfast565/continuous-platform/models/repomodel"
	"github.com/jonfast565/continuous-platform/utilities/jsonutil"
	"github.com/jonfast565/continuous-platform/utilities/webutil"
	"net/http"
	"strconv"
)

var (
	SettingsFilePath = "./repoclient-settings.json"
)

type ClientConfiguration struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
}

type RepoClient struct {
	configuration ClientConfiguration
}

func NewRepoClient() RepoClient {
	var config ClientConfiguration
	jsonutil.DecodeJsonFromFile(SettingsFilePath, &config)
	return RepoClient{configuration: config}
}

func (rc RepoClient) GetRepositories() (*repomodel.RepositoryPackage, error) {
	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		rc.configuration.Hostname,
		strconv.Itoa(rc.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "GetRepositories"})

	// execute request
	request, err := http.NewRequest(constants.PostMethod,
		myUrl.GetUrlStringValue(),
		nil)
	if err != nil {
		return nil, err
	}

	var value repomodel.RepositoryPackage
	err = webutil.ExecuteRequestAndReadJsonBody(request, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (rc RepoClient) GetFile(
	info repomodel.RepositoryFileMetadata) (*miscmodel.FilePayload, error) {
	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		rc.configuration.Hostname,
		strconv.Itoa(rc.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "GetFile"})

	// execute request
	requestBody := info
	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(constants.PostMethod,
		myUrl.GetUrlStringValue(),
		bytes.NewReader(requestJson))
	if err != nil {
		return nil, err
	}

	var value miscmodel.FilePayload
	err = webutil.ExecuteRequestAndReadJsonBody(request, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}
