package main

import (
	"../models/jenkins"
	"../utilities"
	"net/http"
	"strconv"
)

type JenkinsConfiguration struct {
	Port               int    `json:"port"`
	JenkinsUrl         string `json:"jenkinsUrl"`
	Username           string `json:"username"`
	AccessToken        string `json:"accessToken"`
	FolderTemplatePath string `json:"folderTemplatePath"`
}

type JenkinsEndpoint struct {
	configuration JenkinsConfiguration
	client        http.Client
}

func NewJenkinsEndpoint(configuration JenkinsConfiguration) JenkinsEndpoint {
	return JenkinsEndpoint{
		configuration: configuration,
		client:        http.Client{},
	}
}

func (je *JenkinsEndpoint) CreateUpdateJob(crumb jenkins.Crumb, jobFolderUrl string, jobContents string) (*string, error) {
	utilities.LogInfo("Create/Update Jenkins Job -> " + jobFolderUrl)
	request, err := http.NewRequest(utilities.PostMethod, jobFolderUrl /* jobContents */, nil)
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(request)
	addCrumbHeader(crumb, request)
	utilities.AddJsonHeader(request)

	result, err := utilities.ExecuteRequestAndReadStringBody(&je.client, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (je *JenkinsEndpoint) CreateFolder(crumb jenkins.Crumb, jobFolderUrl string) (*string, error) {
	utilities.LogInfo("Create Jenkins Folder -> " + jobFolderUrl)
	request, err := http.NewRequest(utilities.PostMethod, jobFolderUrl /* folderContents */, nil)
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(request)
	addCrumbHeader(crumb, request)
	utilities.AddJsonHeader(request)

	result, err := utilities.ExecuteRequestAndReadStringBody(&je.client, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (je *JenkinsEndpoint) DeleteJobOrFolder(crumb jenkins.Crumb, jobFolderUrl string) (*string, error) {
	utilities.LogInfo("Delete Jenkins Job/Folder -> " + jobFolderUrl)
	request, err := http.NewRequest(utilities.PostMethod, jobFolderUrl, nil)
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(request)
	addCrumbHeader(crumb, request)
	utilities.AddJsonHeader(request)

	result, err := utilities.ExecuteRequestAndReadStringBody(&je.client, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (je *JenkinsEndpoint) GetJenkinsMetadata(crumb jenkins.Crumb) (*jenkins.JobMetadata, error) {
	utilities.LogInfo("Get Jenkins Job Metadata")
	metadataUrl := je.buildJobMetadataUrl()
	request, err := http.NewRequest(utilities.GetMethod, metadataUrl, nil)
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(request)
	addCrumbHeader(crumb, request)
	utilities.AddJsonHeader(request)

	var result jenkins.JobMetadata
	err = utilities.ExecuteRequestAndReadJsonBody(&je.client, request, &result)
	if err != nil {
		return nil, err
	}

	// give parent meaningful data
	result.Name = "Build Server"
	result.Url = je.configuration.JenkinsUrl
	utilities.LogInfo("Metadata retrieved")

	return &result, nil
}

func (je *JenkinsEndpoint) GetJenkinsCrumb() (*jenkins.Crumb, error) {
	utilities.LogInfo("Get Jenkins Crumb")
	request, err := http.NewRequest(utilities.GetMethod, je.buildCrumbUrl(), nil)
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(request)
	utilities.AddJsonHeader(request)

	var result jenkins.Crumb
	err = utilities.ExecuteRequestAndReadJsonBody(&je.client, request, &result)
	if err != nil {
		return nil, err
	}

	utilities.LogInfoMultiline(
		"Crumb reads: ",
		"Request Field: "+result.CrumbRequestField,
		"Crumb: "+result.Crumb)

	return &result, nil
}

func (je *JenkinsEndpoint) addAuthHeader(r *http.Request) {
	r.SetBasicAuth(je.configuration.Username, je.configuration.AccessToken)
}

func addCrumbHeader(crumb jenkins.Crumb, r *http.Request) {
	r.Header.Add(crumb.CrumbRequestField, crumb.Crumb)
}

func (je *JenkinsEndpoint) buildCrumbUrl() string {
	return je.configuration.JenkinsUrl + "crumbIssuer/api/json"
}

func (je *JenkinsEndpoint) buildJobMetadataUrl() string {
	jobDepthString := strconv.Itoa(jenkins.MaximumJobDepth)
	return je.configuration.JenkinsUrl + "api/json?depth=" +
		jobDepthString + "&pretty=false"
}
