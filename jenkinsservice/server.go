package main

import (
	"../models/jenkins"
	"../utilities"
	"../utilities/iteration"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
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

func (je *JenkinsEndpoint) CreateUpdateJob(crumb jenkins.Crumb, request jenkins.JobRequest) (*string, error) {
	utilities.LogInfo("Create/Update Jenkins Job -> " + request.FolderUrl)
	req, err := http.NewRequest(utilities.PostMethod, request.FolderUrl /* jobContents */, nil)
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(req)
	addCrumbHeader(crumb, req)
	utilities.AddJsonHeader(req)

	result, err := utilities.ExecuteRequestAndReadStringBody(&je.client, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (je *JenkinsEndpoint) CreateFolder(crumb jenkins.Crumb, request *jenkins.JobRequest) (*string, error) {
	utilities.LogInfo("Create Jenkins Folder -> " + request.FolderUrl)

	checkFrag := je.buildCreateUpdateCheckUrl(*request, true)
	folderTemplate, err := utilities.RunTemplateFromFile(je.configuration.FolderTemplatePath, utilities.Empty{})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(utilities.PostMethod, checkFrag, strings.NewReader(*folderTemplate))
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(req)
	addCrumbHeader(crumb, req)
	utilities.AddJsonHeader(req)

	result, err := utilities.ExecuteRequestAndReadStringBody(&je.client, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (je *JenkinsEndpoint) DeleteJobOrFolder(crumb jenkins.Crumb, request *jenkins.JobRequest) (*string, error) {
	utilities.LogInfo("Delete Jenkins Job/Folder -> " + request.FolderUrl)

	req, err := http.NewRequest(utilities.PostMethod, je.buildDeleteUrl(*request), nil)
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(req)
	addCrumbHeader(crumb, req)
	utilities.AddJsonHeader(req)

	result, err := utilities.ExecuteRequestAndReadStringBody(&je.client, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (je *JenkinsEndpoint) GetJenkinsMetadata(crumb jenkins.Crumb) (*jenkins.JobMetadata, error) {
	utilities.LogInfo("Get Jenkins Job Metadata")

	request, err := http.NewRequest(utilities.GetMethod, je.buildJobMetadataUrl(), nil)
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

	utilities.LogInfoMultiline(
		"Metadata retrieved: ",
		"Num Top-Level Jobs: "+strconv.Itoa(len(result.Jobs)))
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

func (je *JenkinsEndpoint) buildCreateUpdateCheckUrl(request jenkins.JobRequest, existenceCheck bool) string {
	folderUrlNoBase := strings.Replace(request.FolderUrl, je.configuration.JenkinsUrl, "", -1)
	pathParser := new(iteration.PathParser)
	pathParser.SetActionSeries(folderUrlNoBase)
	pathParser.RemoveLastNActions(2)
	pathString := pathParser.GetPathString(false) + "/"
	folderJobName := iteration.GetLastPathComponent(folderUrlNoBase)
	var actionFragment string
	if existenceCheck {
		actionFragment = "checkJobName?value="
	} else {
		actionFragment = "createItem?name="
	}
	jenkinsFolderQuery := je.configuration.JenkinsUrl + "/" +
		url.QueryEscape(pathString+actionFragment+folderJobName)
	return jenkinsFolderQuery
}

func (je *JenkinsEndpoint) buildDeleteUrl(request jenkins.JobRequest) string {
	return path.Join(je.configuration.JenkinsUrl, request.FolderUrl, "doDelete")
}

func (je *JenkinsEndpoint) buildCrumbUrl() string {
	return je.configuration.JenkinsUrl + "crumbIssuer/api/json"
}

func (je *JenkinsEndpoint) buildJobMetadataUrl() string {
	jobDepthString := strconv.Itoa(jenkins.MaximumJobDepth)
	return je.configuration.JenkinsUrl + "api/json?depth=" +
		jobDepthString + "&pretty=false"
}
