package server

import (
	"../../constants"
	"../../jsonutil"
	"../../logging"
	"../../models/jenkinsmodel"
	"../../templating"
	"../../webutil"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

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

func (je *JenkinsEndpoint) CheckJobExistence(
	crumb jenkinsmodel.JenkinsCrumb,
	request jenkinsmodel.JenkinsJobRequest) (*string, error) {
	logging.LogInfo("Check Existence Jenkins Job -> " + request.GetJobFragmentUrl())
	checkUrl := je.buildCheckUrl(request)
	logging.LogInfo("Check Existence Job URL: " + checkUrl)

	req, err := http.NewRequest(constants.PostMethod, checkUrl, strings.NewReader(request.Contents))
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(req)
	addCrumbHeader(crumb, req)
	webutil.AddXmlHeader(req)

	result, err := webutil.ExecuteRequestAndReadStringBody(&je.client, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (je *JenkinsEndpoint) CreateUpdateJob(
	crumb jenkinsmodel.JenkinsCrumb,
	request jenkinsmodel.JenkinsJobRequest) (*string, error) {
	logging.LogInfo("Create/Update Jenkins Job -> " + request.GetJobFragmentUrl())
	createUpdateUrl := je.buildCreateUpdateUrl(request)
	logging.LogInfo("Create/Update Job URL: " + createUpdateUrl)

	req, err := http.NewRequest(constants.PostMethod, createUpdateUrl, strings.NewReader(request.Contents))
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(req)
	addCrumbHeader(crumb, req)
	webutil.AddXmlHeader(req)

	result, err := webutil.ExecuteRequestAndReadStringBody(&je.client, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (je *JenkinsEndpoint) CreateFolder(
	crumb jenkinsmodel.JenkinsCrumb,
	request jenkinsmodel.JenkinsJobRequest) (*string, error) {
	logging.LogInfo("Create Jenkins Folder -> " + request.GetParentJobFragmentUrl())
	folderUrl := je.buildCreateFolderUrl(request)
	logging.LogInfo("Create Folder URL: " + folderUrl)

	folderTemplate, err := templating.RunTemplateFromFile(je.configuration.FolderTemplatePath, jsonutil.Empty{})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(constants.PostMethod, folderUrl, strings.NewReader(*folderTemplate))
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(req)
	addCrumbHeader(crumb, req)
	webutil.AddFormHeader(req)

	result, err := webutil.ExecuteRequestAndReadStringBody(&je.client, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (je *JenkinsEndpoint) DeleteJobOrFolder(
	crumb jenkinsmodel.JenkinsCrumb,
	request jenkinsmodel.JenkinsJobRequest) (*string, error) {
	logging.LogInfo("Delete Jenkins Job/Folder -> " + request.GetJobFragmentUrl())
	deleteUrl := je.buildDeleteUrl(request)
	logging.LogInfo("Delete Job/Folder URL: " + deleteUrl)

	req, err := http.NewRequest(constants.PostMethod, deleteUrl, nil)
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(req)
	addCrumbHeader(crumb, req)
	webutil.AddXmlHeader(req)

	result, err := webutil.ExecuteRequestAndReadStringBody(&je.client, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (je *JenkinsEndpoint) GetJenkinsMetadata(
	crumb jenkinsmodel.JenkinsCrumb) (*jenkinsmodel.JenkinsJobMetadata, error) {
	logging.LogInfo("Get Jenkins Job Metadata")
	metadataUrl := je.buildJobMetadataUrl()
	logging.LogInfo("Metadata URL: " + metadataUrl)

	request, err := http.NewRequest(constants.GetMethod, metadataUrl, nil)
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(request)
	addCrumbHeader(crumb, request)
	webutil.AddJsonHeader(request)

	var result jenkinsmodel.JenkinsJobMetadata
	err = webutil.ExecuteRequestAndReadJsonBody(&je.client, request, &result)
	if err != nil {
		return nil, err
	}

	// give parent meaningful data
	result.Name = "Build Server"
	result.Url = je.configuration.GetJenkinsUrl()

	logging.LogInfoMultiline(
		"Metadata retrieved: ",
		"Num Top-Level Jobs: "+strconv.Itoa(len(result.Jobs)))
	return &result, nil
}

func (je *JenkinsEndpoint) GetJenkinsCrumb() (*jenkinsmodel.JenkinsCrumb, error) {
	logging.LogInfo("Get Jenkins Crumb")
	crumbUrl := je.buildCrumbUrl()
	logging.LogInfo("Crumb URL: " + crumbUrl)

	request, err := http.NewRequest(constants.GetMethod, crumbUrl, nil)
	if err != nil {
		return nil, err
	}

	je.addAuthHeader(request)
	webutil.AddJsonHeader(request)

	var result jenkinsmodel.JenkinsCrumb
	err = webutil.ExecuteRequestAndReadJsonBody(&je.client, request, &result)
	if err != nil {
		return nil, err
	}

	logging.LogInfoMultiline(
		"Crumb reads: ",
		"Request Field: "+result.CrumbRequestField,
		"Crumb: "+result.Crumb)

	return &result, nil
}

func (je *JenkinsEndpoint) addAuthHeader(r *http.Request) {
	if je.configuration.JenkinsCredentialsRequired {
		r.SetBasicAuth(je.configuration.JenkinsUsername, je.configuration.JenkinsAccessToken)
	}
}

func addCrumbHeader(crumb jenkinsmodel.JenkinsCrumb, r *http.Request) {
	r.Header.Add(crumb.CrumbRequestField, crumb.Crumb)
}

func (je *JenkinsEndpoint) buildCreateFolderUrl(request jenkinsmodel.JenkinsJobRequest) string {
	createFolderPath := je.configuration.GetJenkinsUrlObject()
	createFolderPath.AppendPathFragments(request.GetParentJobFragments())
	createFolderPath.AppendPathFragment("createItem")
	folderName := request.GetLastFragment()
	// this is the weirdest api call on the face of the planet
	// https://gist.github.com/stuart-warren/7786892
	createFolderPath.AppendQueryValue("name", folderName)
	createFolderPath.AppendQueryValue("mode", "com.cloudbees.hudson.plugins.folder.Folder")
	createFolderPath.AppendQueryValue("from", "")
	result, _ := json.Marshal(jenkinsmodel.NewFolderRequest(folderName))
	createFolderPath.AppendQueryValue("json", string(result))
	createFolderPath.AppendQueryValue("Submit", "OK")
	return createFolderPath.GetForciblyEncodedUrlStringValue()
}

func (je *JenkinsEndpoint) buildCheckUrl(request jenkinsmodel.JenkinsJobRequest) string {
	createUpdatePath := je.configuration.GetJenkinsUrlObject()
	createUpdatePath.AppendPathFragments(request.GetParentJobFragments())
	createUpdatePath.AppendPathFragment("checkJobName")
	createUpdatePath.AppendQueryValue("value", request.GetLastFragment())
	return createUpdatePath.GetForciblyEncodedUrlStringValue()
}

func (je *JenkinsEndpoint) buildCreateUpdateUrl(request jenkinsmodel.JenkinsJobRequest) string {
	createUpdatePath := je.configuration.GetJenkinsUrlObject()
	createUpdatePath.AppendPathFragments(request.GetParentJobFragments())
	createUpdatePath.AppendPathFragment("createItem")
	createUpdatePath.AppendQueryValue("name", request.GetLastFragment())
	return createUpdatePath.GetForciblyEncodedUrlStringValue()
}

func (je *JenkinsEndpoint) buildDeleteUrl(request jenkinsmodel.JenkinsJobRequest) string {
	deletePath := je.configuration.GetJenkinsUrlObject()
	deletePath.AppendPathFragments(request.GetJobFragments())
	deletePath.AppendPathFragment("doDelete")
	return deletePath.GetForciblyEncodedUrlStringValue()
}

func (je *JenkinsEndpoint) buildCrumbUrl() string {
	crumbPath := je.configuration.GetJenkinsUrlObject()
	crumbPath.AppendPathFragments([]string{"crumbIssuer", "api", "json"})
	return crumbPath.GetUrlStringValue()
}

func (je *JenkinsEndpoint) buildJobMetadataUrl() string {
	jobDepthString := strconv.Itoa(jenkinsmodel.JenkinsMaximumJobDepth)
	metadataPath := je.configuration.GetJenkinsUrlObject()
	metadataPath.AppendPathFragments([]string{"api", "json"})
	metadataPath.AppendQueryValue("depth", jobDepthString)
	metadataPath.AppendQueryValue("pretty", "false")
	return metadataPath.GetUrlStringValue()
}
