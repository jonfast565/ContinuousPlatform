// Client for the MSBuild service
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

// Client configuration for the MSBuild services
type ClientConfiguration struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
}

// Client for the MSBuild service, with configuration
type MsBuildClient struct {
	configuration ClientConfiguration
}

// Constructor for an MsBuildClient
func NewMsBuildClient() MsBuildClient {
	var config ClientConfiguration
	jsonutil.DecodeJsonFromFile(SettingsFilePath, &config)
	return MsBuildClient{configuration: config}
}

// Gets the metadata from a solution when passed a file payload containing bytes corresponding to the solution file
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

// Gets the metadata from a project when passed a file payload containing bytes corresponding to the project file
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

// Gets the metadata from a publish profile when passed a file payload containing bytes corresponding to the publish profile file
// TODO: This is similar to obtaining a project, but having a different use case. Maybe endpoints could be merged?
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
