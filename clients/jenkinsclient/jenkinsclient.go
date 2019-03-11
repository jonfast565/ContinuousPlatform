// Client for the Jenkins service
package jenkinsclient

import (
	"bytes"
	"encoding/json"
	"github.com/jonfast565/continuous-platform/constants"
	"github.com/jonfast565/continuous-platform/models/jenkinsmodel"
	"github.com/jonfast565/continuous-platform/models/miscmodel"
	"github.com/jonfast565/continuous-platform/utilities/jsonutil"
	"github.com/jonfast565/continuous-platform/utilities/webutil"
	"net/http"
	"strconv"
)

var (
	SettingsFilePath = "./jenkinsclient-settings.json"
)

// Configuration for the client that calls the Jenkins service
type ClientConfiguration struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
}

// Jenkins client with configuration
type JenkinsClient struct {
	configuration ClientConfiguration
}

// Creates a new Jenkins client,
// requires configuration
func NewJenkinsClient() JenkinsClient {
	var config ClientConfiguration
	jsonutil.DecodeJsonFromFile(SettingsFilePath, &config)
	return JenkinsClient{configuration: config}
}

// Gets the Jenkins Crumb, used to validate subsequent requests to the Jenkins API,
// and allows for CSRF protection
// TODO: Should not use this API directly, should be called and cached on the server for performance
// TODO: Above technique when used may present errors when there are none, say, if the crumb changes. Investigate.
func (jc JenkinsClient) GetJenkinsCrumb() (*jenkinsmodel.JenkinsCrumb, error) {
	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		jc.configuration.Hostname,
		strconv.Itoa(jc.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "GetJenkinsCrumb"})

	// execute request
	request, err := http.NewRequest(constants.PostMethod,
		myUrl.GetUrlStringValue(),
		nil)
	if err != nil {
		return nil, err
	}
	webutil.AddFormHeader(request)

	var value jenkinsmodel.JenkinsCrumb
	err = webutil.ExecuteRequestAndReadJsonBody(request, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

// Gets a tree of all Jenkins jobs
func (jc JenkinsClient) GetJenkinsMetadata() (*jenkinsmodel.JenkinsJobMetadata, error) {
	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		jc.configuration.Hostname,
		strconv.Itoa(jc.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "GetJenkinsMetadata"})

	// execute request
	request, err := http.NewRequest(constants.PostMethod,
		myUrl.GetUrlStringValue(),
		nil)
	if err != nil {
		return nil, err
	}
	webutil.AddFormHeader(request)

	var value jenkinsmodel.JenkinsJobMetadata
	err = webutil.ExecuteRequestAndReadJsonBody(request, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

// Creates a job
// NOTE: This method is not idempotent.
func (jc JenkinsClient) CreateJob(jobRequest jenkinsmodel.JenkinsJobRequest) (*string, error) {
	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		jc.configuration.Hostname,
		strconv.Itoa(jc.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "CreateJob"})

	requestJson, err := json.Marshal(jobRequest)
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
	webutil.AddFormHeader(request)

	value, err := webutil.ExecuteRequestAndReadStringBody(request, false)
	if err != nil {
		return nil, err
	}

	return value, nil
}

// Updates a job idempotently
func (jc JenkinsClient) UpdateJob(jobRequest jenkinsmodel.JenkinsJobRequest) (*string, error) {
	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		jc.configuration.Hostname,
		strconv.Itoa(jc.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "UpdateJob"})

	requestJson, err := json.Marshal(jobRequest)
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
	webutil.AddFormHeader(request)

	value, err := webutil.ExecuteRequestAndReadStringBody(request, false)
	if err != nil {
		return nil, err
	}

	return value, nil
}

// Creates a Jenkins folder
// NOTE: This is not idempotent
func (jc JenkinsClient) CreateFolder(jobRequest jenkinsmodel.JenkinsJobRequest) (*string, error) {
	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		jc.configuration.Hostname,
		strconv.Itoa(jc.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "CreateFolder"})

	requestJson, err := json.Marshal(jobRequest)
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
	webutil.AddFormHeader(request)

	value, err := webutil.ExecuteRequestAndReadStringBody(request, false)
	if err != nil {
		return nil, err
	}

	return value, nil
}

// Deletes a Jenkins item (most commonly a job of any type or folder)
func (jc JenkinsClient) DeleteJobOrFolder(jobRequest jenkinsmodel.JenkinsJobRequest) (*string, error) {
	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		jc.configuration.Hostname,
		strconv.Itoa(jc.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "DeleteJobOrFolder"})

	requestJson, err := json.Marshal(jobRequest)
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
	webutil.AddFormHeader(request)

	value, err := webutil.ExecuteRequestAndReadStringBody(request, false)
	if err != nil {
		return nil, err
	}

	return value, nil
}

// Check if a job exists, will return a string indicating a yes or no
func (jc JenkinsClient) CheckJobExists(jobRequest jenkinsmodel.JenkinsJobRequest) (*bool, error) {
	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		jc.configuration.Hostname,
		strconv.Itoa(jc.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "CheckJob"})

	requestJson, err := json.Marshal(jobRequest)
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

	var value miscmodel.YesNoResult
	err = webutil.ExecuteRequestAndReadJsonBody(request, &value)
	if err != nil {
		return nil, err
	}

	return &value.Value, nil
}
