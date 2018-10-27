package repoclient

import (
	"../../constants"
	"../../jsonutil"
	"../../models/miscmodel"
	"../../models/repomodel"
	"../../webutil"
	"bytes"
	"encoding/json"
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
	client        http.Client
}

func NewRepoClient() RepoClient {
	var config ClientConfiguration
	jsonutil.DecodeJsonFromFile(SettingsFilePath, &config)
	return RepoClient{configuration: config, client: http.Client{}}
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
	err = webutil.ExecuteRequestAndReadJsonBody(&rc.client, request, value)
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

	request, err := http.NewRequest(constants.GetMethod,
		myUrl.GetUrlStringValue(),
		bytes.NewReader(requestJson))
	if err != nil {
		return nil, err
	}

	var value miscmodel.FilePayload
	err = webutil.ExecuteRequestAndReadJsonBody(&rc.client, request, value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}
