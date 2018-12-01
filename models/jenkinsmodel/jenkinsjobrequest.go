package jenkinsmodel

import (
	"../../webutil"
	"net/url"
)

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
		if i >= len(jr.FolderSegments)-1 {
			result = append(result, segment)
			return result
		} else {
			result = append(result, []string{segment, "job"}...)
		}
	}
	return result
}

func (jr *JenkinsJobRequest) GetParentJobFragments() []string {
	fragments := jr.GetJobFragments()
	shavedFragments := fragments[:len(fragments)-2]
	return shavedFragments
}

func (jr *JenkinsJobRequest) GetParentJobFragmentUrl() string {
	myUrl := webutil.NewEmptyUrl()
	myUrl.AppendPathFragments(jr.GetParentJobFragments())
	return myUrl.GetFragmentValue()
}

func (jr *JenkinsJobRequest) GetLastFragment() string {
	return jr.FolderSegments[len(jr.FolderSegments)-1]
}

func (jr *JenkinsJobRequest) SanitizeSegments() {
	var result []string
	for _, frag := range jr.FolderSegments {
		result = append(result, url.QueryEscape(frag))
	}
	jr.FolderSegments = result
}
