package models

import "../webutil"

const (
	JenkinsMaximumJobDepth int = 4
)

type JenkinsEditType int

const (
	RemoveJobFolder JenkinsEditType = 0
	AddUpdateJob    JenkinsEditType = 1
	AddFolder       JenkinsEditType = 2
)

type JenkinsCrumb struct {
	Crumb             string
	CrumbRequestField string
}

type JenkinsJobMetadata struct {
	Name string
	Url  string
	Jobs []JenkinsJobMetadata
}

type JenkinsJobRecord struct {
	TopLevelFolder    string
	MidLevel1Folder   string
	MidLevel2Folder   string
	BottomLevelFolder string
}

type JenkinsJobRequest struct {
	FolderSegments []string
	Contents       string
}

func (jr *JenkinsJobRequest) GetJobFragmentUrl() string {
	myUrl := webutil.NewEmptyUrl()
	myUrl.AppendPathFragments(jr.GetJobFragments())
	return myUrl.GetFragmentValue()
}

func (jr *JenkinsJobRequest) GetJobFragments() []string {
	result := make([]string, 0)
	result = append(result, "job")
	for i, segment := range jr.FolderSegments {
		if i == len(jr.FolderSegments)-1 {
			result = append(result, segment)
			return result
		} else {
			result = append(result, []string{segment, "job"}...)
		}
	}
	return result
}

func (jr *JenkinsJobRequest) GetParentJobFragments() []string {
	result := make([]string, 0)
	result = append(result, "job")
	for i, segment := range jr.FolderSegments {
		if i == len(jr.FolderSegments)-2 {
			result = append(result, segment)
			return result
		} else {
			result = append(result, []string{segment, "job"}...)
		}
	}
	return result
}

func (jr *JenkinsJobRequest) GetLastFragment() string {
	return jr.FolderSegments[len(jr.FolderSegments)-1]
}

type JenkinsEdit struct {
	Name      string
	Url       string
	EditType  JenkinsEditType
	JobRecord JenkinsJobRecord
}

type JenkinsFolderRequest struct {
	Name   string `json:"name"`
	Mode   string `json:"mode"`
	From   string `json:"from"`
	Submit string `json:"Submit"`
}

func NewFolderRequest(folderName string) JenkinsFolderRequest {
	return JenkinsFolderRequest{
		Name:   folderName,
		Mode:   "com.cloudbees.hudson.plugins.folder.Folder",
		From:   "",
		Submit: "OK",
	}
}
