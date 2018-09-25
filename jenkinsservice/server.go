package main

type JenkinsConfiguration struct {
	Port               int    `json:"port"`
	JenkinsUrl         string `json:"jenkinsUrl"`
	Username           string `json:"username"`
	AccessToken        string `json:"accessToken"`
	FolderTemplatePath string `json:"folderTemplatePath"`
}
