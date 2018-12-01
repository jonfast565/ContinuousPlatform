package jenkinsclient

import (
	"../../constants"
	"../../jsonutil"
	"../../models/jenkinsmodel"
	"../../webutil"
	"bytes"
	"encoding/json"
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
	client        http.Client
}

func NewJenkinsClient() JenkinsClient {
	var config ClientConfiguration
	jsonutil.DecodeJsonFromFile(SettingsFilePath, &config)
	return JenkinsClient{configuration: config, client: http.Client{Timeout: constants.ClientTimeout}}
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
	err = webutil.ExecuteRequestAndReadJsonBody(&jc.client, request, &value)
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
	err = webutil.ExecuteRequestAndReadJsonBody(&jc.client, request, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (jc JenkinsClient) CreateUpdateJob(jobRequest jenkinsmodel.JenkinsJobRequest) (*string, error) {
	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		jc.configuration.Hostname,
		strconv.Itoa(jc.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "CreateUpdateJob"})

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

	value, err := webutil.ExecuteRequestAndReadStringBody(&jc.client, request)
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

	value, err := webutil.ExecuteRequestAndReadStringBody(&jc.client, request)
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

	value, err := webutil.ExecuteRequestAndReadStringBody(&jc.client, request)
	if err != nil {
		return nil, err
	}

	return value, nil
}
