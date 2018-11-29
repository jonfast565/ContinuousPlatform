package jenkinsmodel

type JenkinsJobRecordType string

var (
	FreestyleJob JenkinsJobRecordType = "hudson.model.FreeStyleProject"
	PipelineJob  JenkinsJobRecordType = "org.jenkinsci.plugins.workflow.job.WorkflowJob"
	Folder       JenkinsJobRecordType = "com.cloudbees.hudson.plugins.folder.Folder"
)

type JenkinsJobRecord struct {
	KeyElements []string
	Type        JenkinsJobRecordType
}
