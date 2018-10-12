package jenkins

import "../../utilities/web"

const (
	MaximumJobDepth int = 4
)

type EditType int

const (
	RemoveJobFolder EditType = 0
	AddUpdateJob    EditType = 1
	AddFolder       EditType = 2
)

type Crumb struct {
	Crumb             string
	CrumbRequestField string
}

type JobMetadata struct {
	Name string
	Url  string
	Jobs []JobMetadata
}

type JobRecord struct {
	TopLevelFolder    string
	MidLevel1Folder   string
	MidLevel2Folder   string
	BottomLevelFolder string
}

type JobRequest struct {
	FolderSegments []string
	Contents       string
}

func (jr *JobRequest) GetJobFragmentUrl() string {
	myUrl := web.NewEmptyUrl()
	myUrl.AppendPathFragments(jr.GetJobFragments())
	return myUrl.GetFragmentValue()
}

func (jr *JobRequest) GetJobFragments() []string {
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

func (jr *JobRequest) GetParentJobFragments() []string {
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

func (jr *JobRequest) GetLastFragment() string {
	return jr.FolderSegments[len(jr.FolderSegments)-1]
}

type Edit struct {
	Name      string
	Url       string
	EditType  EditType
	JobRecord JobRecord
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
