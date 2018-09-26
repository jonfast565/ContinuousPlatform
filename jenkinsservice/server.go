package main

import (
	"../models/jenkins"
	"net/http"
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

func (je *JenkinsEndpoint) CreateUpdateJob() error {

}

func (je *JenkinsEndpoint) CreateFolder() error {

}

func (je *JenkinsEndpoint) DeleteJobOrFolder() error {

}

func (je *JenkinsEndpoint) GetJenkinsMetadata() (*jenkins.JobMetadata, error) {

}

func (je *JenkinsEndpoint) GetJenkinsCrumb() (*jenkins.Crumb, error) {

}

func (je *JenkinsEndpoint) addAuthHeader(r *http.Request) {
	r.SetBasicAuth(je.configuration.Username, je.configuration.AccessToken)
}

func (je *JenkinsEndpoint) buildCrumbUrl() string {
	return je.configuration.JenkinsUrl + "/crumbIssuer/api/json"
}

func (je *JenkinsEndpoint) buildJobUrl() string {
	return je.configuration.JenkinsUrl + "/api/json?depth=" +
		string(jenkins.MaximumJobDepth) + "&pretty=false"
}
