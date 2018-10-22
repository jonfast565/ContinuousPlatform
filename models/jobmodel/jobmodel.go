package jobmodel

type JobStatus int

const (
	Stopped JobStatus = 0
	Running JobStatus = 1
	Errored JobStatus = 2
)

type JobDetails struct {
	Trigger bool
	Status  JobStatus
}

type JobController struct {
	StopJobs          bool
	DetectChanges     JobDetails
	BuildDeliverables JobDetails
	GenerateScripts   JobDetails
	DeployJenkinsJobs JobDetails
}

func NewJobController() JobController {
	return JobController{
		StopJobs:          false,
		DetectChanges:     JobDetails{Trigger: false, Status: Stopped},
		BuildDeliverables: JobDetails{Trigger: false, Status: Stopped},
		GenerateScripts:   JobDetails{Trigger: false, Status: Stopped},
		DeployJenkinsJobs: JobDetails{Trigger: false, Status: Stopped},
	}
}
