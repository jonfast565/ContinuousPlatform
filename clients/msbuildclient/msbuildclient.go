package msbuildclient

import (
	"bytes"
	"encoding/json"
	"github.com/jonfast565/continuous-platform/constants"
	"github.com/jonfast565/continuous-platform/models/miscmodel"
	"github.com/jonfast565/continuous-platform/models/projectmodel"
	"github.com/jonfast565/continuous-platform/utilities/jsonutil"
	"github.com/jonfast565/continuous-platform/utilities/webutil"
	"net/http"
	"strconv"
)

var (
	SettingsFilePath = "./msbuildclient-settings.json"
)

type ClientConfiguration struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
}

type MsBuildClient struct {
	configuration ClientConfiguration
}

func NewMsBuildClient() MsBuildClient {
	var config ClientConfiguration
	jsonutil.DecodeJsonFromFile(SettingsFilePath, &config)
	return MsBuildClient{configuration: config}
}

func (msbc MsBuildClient) GetSolution(
	payload miscmodel.FilePayload) (*projectmodel.MsBuildSolution, error) {
	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		msbc.configuration.Hostname,
		strconv.Itoa(msbc.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "GetSolution", "Bytes"})

	requestJson, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// execute request
	request, err := http.NewRequest(constants.PostMethod,
		myUrl.GetUrlStringValue(),
		bytes.NewReader(requestJson))
	if err != nil {
		return nil, err
	}

	var value projectmodel.MsBuildSolution
	err = webutil.ExecuteRequestAndReadJsonBody(request, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (msbc MsBuildClient) GetProject(
	payload miscmodel.FilePayload) (*projectmodel.MsBuildProject, error) {
	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		msbc.configuration.Hostname,
		strconv.Itoa(msbc.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "GetProject", "Bytes"})

	requestJson, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// execute request
	request, err := http.NewRequest(constants.PostMethod,
		myUrl.GetUrlStringValue(),
		bytes.NewReader(requestJson))
	if err != nil {
		return nil, err
	}

	var value projectmodel.MsBuildProject
	err = webutil.ExecuteRequestAndReadJsonBody(request, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (msbc MsBuildClient) GetPublishProfile(
	payload miscmodel.FilePayload) (*projectmodel.MsBuildPublishProfile, error) {
	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		msbc.configuration.Hostname,
		strconv.Itoa(msbc.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "GetPublishProfile", "Bytes"})

	requestJson, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// execute request
	request, err := http.NewRequest(constants.PostMethod,
		myUrl.GetUrlStringValue(),
		bytes.NewReader(requestJson))
	if err != nil {
		return nil, err
	}

	var value projectmodel.MsBuildPublishProfile
	err = webutil.ExecuteRequestAndReadJsonBody(request, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}
