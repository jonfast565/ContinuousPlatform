package main

import "net/http"

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

func (je *JenkinsEndpoint) CreateUpdateJob() {

}

func (je *JenkinsEndpoint) CreateFolder() {

}

func (je *JenkinsEndpoint) DeleteJobOrFolder() {

}

func (je *JenkinsEndpoint) GetJenkinsMetadata() {

}

func (je *JenkinsEndpoint) GetJenkinsCrumb() {

}
