package jenkinsmodel

import (
	"github.com/jonfast565/continuous-platform/utilities/stringutil"
)

type JenkinsEdit struct {
	Keys     []string
	Contents string
	EditType JenkinsEditType
}

func (je JenkinsEdit) GetJobRequest() JenkinsJobRequest {
	return JenkinsJobRequest{
		FolderSegments: je.Keys,
		Contents:       je.Contents,
	}
}

type JenkinsEditList []JenkinsEdit

func (jel JenkinsEditList) Len() int {
	return len(jel)
}
func (jel JenkinsEditList) Swap(i, j int) {
	jel[i], jel[j] = jel[j], jel[i]
}
func (jel JenkinsEditList) Less(i, j int) bool {
	return len(jel[i].Keys) < len(jel[j].Keys) &&
		stringutil.StringArrayCompareNumeric(jel[i].Keys, jel[j].Keys) < 0
}
