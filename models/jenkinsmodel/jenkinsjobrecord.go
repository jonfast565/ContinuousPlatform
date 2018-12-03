package jenkinsmodel

type JenkinsJobType string

var (
	FreestyleJob JenkinsJobType = "hudson.model.FreeStyleProject"
	PipelineJob  JenkinsJobType = "org.jenkinsci.plugins.workflow.job.WorkflowJob"
	Folder       JenkinsJobType = "com.cloudbees.hudson.plugins.folder.Folder"
	BuildServer  JenkinsJobType = "hudson.model.Hudson"
)

type JenkinsJobRecord struct {
	KeyElements []string
	Type        JenkinsJobType
}
