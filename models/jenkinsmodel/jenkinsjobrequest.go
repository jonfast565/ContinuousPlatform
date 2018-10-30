package jenkinsmodel

import "../../webutil"

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
