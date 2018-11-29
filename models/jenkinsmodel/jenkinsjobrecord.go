package jenkinsmodel

type JenkinsJobRecordType string

var (
	File   JenkinsJobRecordType = "File"
	Folder JenkinsJobRecordType = "Folder"
)

type JenkinsJobRecord struct {
	KeyElements []string
	Type        JenkinsJobRecordType
}
