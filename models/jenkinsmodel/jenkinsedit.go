package jenkinsmodel

type JenkinsEdit struct {
	Name      string
	Url       string
	EditType  JenkinsEditType
	JobRecord JenkinsJobRecord
}
