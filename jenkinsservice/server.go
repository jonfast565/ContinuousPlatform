package main

import (
	"../models/jenkins"
	"../utilities"
	"../utilities/web"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type JenkinsConfiguration struct {
	Port               int    `json:"port"`
	JenkinsScheme      string `json:"jenkinsScheme"`
	JenkinsHost        string `json:"jenkinsHost"`
	JenkinsPort        string `json:"jenkinsPort"`
	JenkinsUsername    string `json:"jenkinsUsername"`
	JenkinsAccessToken string `json:"jenkinsAccessToken"`
	FolderTemplatePath string `json:"folderTemplatePath"`
}

func (jc *JenkinsConfiguration) GetJenkinsUrl() string {
	myUrl := web.NewEmptyUrl()
	myUrl.SetBase(jc.JenkinsScheme, jc.JenkinsHost, jc.JenkinsPort)
	return myUrl.GetBasePath()
}

func (jc *JenkinsConfiguration) GetJenkinsUrlObject() *web.MyUrl {
	myUrl := web.NewEmptyUrl()
	myUrl.SetBase(jc.JenkinsScheme, jc.JenkinsHost, jc.JenkinsPort)
	return &myUrl
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

func (je *JenkinsEndpoint) CheckJobExistence(crumb jenkins.Crumb, request jenkins.JobRequest) (*string, error) {
	utilities.LogInfo("Check Existence Jenkins Job -> " + request.GetJobFragmentUrl())
	checkUrl := je.buildCheckUrl(request)
	utilities.LogInfo("Check Existence Job URL: " + checkUrl)

	req, err := http.NewRequest(utilities.PostMethod, checkUrl, strings.NewReader(request.Contents))
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(req)
	addCrumbHeader(crumb, req)
	web.AddXmlHeader(req)

	result, err := web.ExecuteRequestAndReadStringBody(&je.client, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (je *JenkinsEndpoint) CreateUpdateJob(crumb jenkins.Crumb, request jenkins.JobRequest) (*string, error) {
	utilities.LogInfo("Create/Update Jenkins Job -> " + request.GetJobFragmentUrl())
	createUpdateUrl := je.buildCreateUpdateUrl(request)
	utilities.LogInfo("Create/Update Job URL: " + createUpdateUrl)

	req, err := http.NewRequest(utilities.PostMethod, createUpdateUrl, strings.NewReader(request.Contents))
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(req)
	addCrumbHeader(crumb, req)
	web.AddXmlHeader(req)

	result, err := web.ExecuteRequestAndReadStringBody(&je.client, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (je *JenkinsEndpoint) CreateFolder(crumb jenkins.Crumb, request jenkins.JobRequest) (*string, error) {
	utilities.LogInfo("Create Jenkins Folder -> " + request.GetJobFragmentUrl())
	folderUrl := je.buildCreateFolderUrl(request)
	utilities.LogInfo("Create Folder URL: " + folderUrl)

	folderTemplate, err := utilities.RunTemplateFromFile(je.configuration.FolderTemplatePath, utilities.Empty{})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(utilities.PostMethod, folderUrl, strings.NewReader(*folderTemplate))
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(req)
	addCrumbHeader(crumb, req)
	web.AddFormHeader(req)

	result, err := web.ExecuteRequestAndReadStringBody(&je.client, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (je *JenkinsEndpoint) DeleteJobOrFolder(crumb jenkins.Crumb, request jenkins.JobRequest) (*string, error) {
	utilities.LogInfo("Delete Jenkins Job/Folder -> " + request.GetJobFragmentUrl())
	deleteUrl := je.buildDeleteUrl(request)
	utilities.LogInfo("Delete Job/Folder URL: " + deleteUrl)

	req, err := http.NewRequest(utilities.PostMethod, deleteUrl, nil)
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(req)
	addCrumbHeader(crumb, req)
	web.AddXmlHeader(req)

	result, err := web.ExecuteRequestAndReadStringBody(&je.client, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (je *JenkinsEndpoint) GetJenkinsMetadata(crumb jenkins.Crumb) (*jenkins.JobMetadata, error) {
	utilities.LogInfo("Get Jenkins Job Metadata")
	metadataUrl := je.buildJobMetadataUrl()
	utilities.LogInfo("Metadata URL: " + metadataUrl)

	request, err := http.NewRequest(utilities.GetMethod, metadataUrl, nil)
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(request)
	addCrumbHeader(crumb, request)
	web.AddJsonHeader(request)

	var result jenkins.JobMetadata
	err = web.ExecuteRequestAndReadJsonBody(&je.client, request, &result)
	if err != nil {
		return nil, err
	}

	// give parent meaningful data
	result.Name = "Build Server"
	result.Url = je.configuration.GetJenkinsUrl()

	utilities.LogInfoMultiline(
		"Metadata retrieved: ",
		"Num Top-Level Jobs: "+strconv.Itoa(len(result.Jobs)))
	return &result, nil
}

func (je *JenkinsEndpoint) GetJenkinsCrumb() (*jenkins.Crumb, error) {
	utilities.LogInfo("Get Jenkins Crumb")
	crumbUrl := je.buildCrumbUrl()
	utilities.LogInfo("Crumb URL: " + crumbUrl)

	request, err := http.NewRequest(utilities.GetMethod, crumbUrl, nil)
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(request)
	web.AddJsonHeader(request)

	var result jenkins.Crumb
	err = web.ExecuteRequestAndReadJsonBody(&je.client, request, &result)
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
	r.SetBasicAuth(je.configuration.JenkinsUsername, je.configuration.JenkinsAccessToken)
}

func addCrumbHeader(crumb jenkins.Crumb, r *http.Request) {
	r.Header.Add(crumb.CrumbRequestField, crumb.Crumb)
}

func (je *JenkinsEndpoint) buildCreateFolderUrl(request jenkins.JobRequest) string {
	createFolderPath := je.configuration.GetJenkinsUrlObject()
	createFolderPath.AppendPathFragments(request.GetParentJobFragments())
	createFolderPath.AppendPathFragment("createItem")
	folderName := request.GetLastFragment()
	// this is the weirdest api call on the face of the planet
	// https://gist.github.com/stuart-warren/7786892
	createFolderPath.AppendQueryValue("name", folderName)
	createFolderPath.AppendQueryValue("mode", "com.cloudbees.hudson.plugins.folder.Folder")
	createFolderPath.AppendQueryValue("from", "")
	result, _ := json.Marshal(jenkins.NewFolderRequest(folderName))
	createFolderPath.AppendQueryValue("json", string(result))
	createFolderPath.AppendQueryValue("Submit", "OK")
	return createFolderPath.GetUrlStringValue()
}

// obsolete?
func (je *JenkinsEndpoint) buildCheckUrl(request jenkins.JobRequest) string {
	createUpdatePath := je.configuration.GetJenkinsUrlObject()
	createUpdatePath.AppendPathFragments(request.GetParentJobFragments())
	createUpdatePath.AppendPathFragment("checkJobName")
	createUpdatePath.AppendQueryValue("value", request.GetLastFragment())
	return createUpdatePath.GetUrlStringValue()
}

func (je *JenkinsEndpoint) buildCreateUpdateUrl(request jenkins.JobRequest) string {
	createUpdatePath := je.configuration.GetJenkinsUrlObject()
	createUpdatePath.AppendPathFragments(request.GetParentJobFragments())
	createUpdatePath.AppendPathFragment("createItem")
	createUpdatePath.AppendQueryValue("name", request.GetLastFragment())
	return createUpdatePath.GetUrlStringValue()
}

func (je *JenkinsEndpoint) buildDeleteUrl(request jenkins.JobRequest) string {
	deletePath := je.configuration.GetJenkinsUrlObject()
	deletePath.AppendPathFragments(request.GetJobFragments())
	deletePath.AppendPathFragment("doDelete")
	return deletePath.GetUrlStringValue()
}

func (je *JenkinsEndpoint) buildCrumbUrl() string {
	crumbPath := je.configuration.GetJenkinsUrlObject()
	crumbPath.AppendPathFragments([]string{"crumbIssuer", "api", "json"})
	return crumbPath.GetUrlStringValue()
}

func (je *JenkinsEndpoint) buildJobMetadataUrl() string {
	jobDepthString := strconv.Itoa(jenkins.MaximumJobDepth)
	metadataPath := je.configuration.GetJenkinsUrlObject()
	metadataPath.AppendPathFragments([]string{"api", "json"})
	metadataPath.AppendQueryValue("depth", jobDepthString)
	metadataPath.AppendQueryValue("pretty", "false")
	return metadataPath.GetUrlStringValue()
}
