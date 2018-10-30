package server

import "../../webutil"

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
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(jc.JenkinsScheme, jc.JenkinsHost, jc.JenkinsPort)
	return myUrl.GetBasePath()
}

func (jc *JenkinsConfiguration) GetJenkinsUrlObject() *webutil.MyUrl {
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(jc.JenkinsScheme, jc.JenkinsHost, jc.JenkinsPort)
	return &myUrl
}
