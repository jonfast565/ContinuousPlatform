package jenkinsclient

import (
	"../../constants"
	"../../jsonutil"
	"../../models/jenkinsmodel"
	"../../webutil"
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

	var value jenkinsmodel.JenkinsJobMetadata
	err = webutil.ExecuteRequestAndReadJsonBody(&jc.client, request, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}
