package jenkinsmodel

type JenkinsJobMetadata struct {
	Name string
	Url  string
	Jobs []JenkinsJobMetadata
}
