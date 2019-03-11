package jenkinsclient

import (
	"bytes"
	"encoding/json"
	"github.com/jonfast565/continuous-platform/constants"
	"github.com/jonfast565/continuous-platform/jsonutil"
	"github.com/jonfast565/continuous-platform/models/jenkinsmodel"
	"github.com/jonfast565/continuous-platform/models/miscmodel"
	"github.com/jonfast565/continuous-platform/webutil"
	"net/http"
	"strconv"
)

var (
	SettingsFilePath = "./jenkinsclient-settings.json"
)

type ClientConfiguration struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
}

type JenkinsClient struct {
	configuration ClientConfiguration
}

func NewJenkinsClient() JenkinsClient {
	var config ClientConfiguration
	jsonutil.DecodeJsonFromFile(SettingsFilePath, &config)
	return JenkinsClient{configuration: config}
}

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
