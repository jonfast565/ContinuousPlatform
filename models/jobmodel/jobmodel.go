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

func (jd *JobDetails) UnsetTriggerBeginRun() {
	jd.Trigger = false
	jd.Status = Running
}

func (jd *JobDetails) SetJobErrored() {
	jd.Status = Errored
}

func (jd *JobDetails) SetJobStoppedOrErrored() {
	if jd.Status != Errored {
		jd.Status = Stopped
	}
}

func (jd *JobDetails) TriggerJob() {
	jd.Trigger = true
}

func (jd *JobDetails) Errored() bool {
	return jd.Status == Errored
}

type JobController struct {
	DetectChanges     JobDetails
	BuildDeliverables JobDetails
	GenerateScripts   JobDetails
	DeployJenkinsJobs JobDetails
}

func NewJobController() JobController {
	return JobController{
		DetectChanges:     JobDetails{Trigger: false, Status: Stopped},
		BuildDeliverables: JobDetails{Trigger: false, Status: Stopped},
		GenerateScripts:   JobDetails{Trigger: false, Status: Stopped},
		DeployJenkinsJobs: JobDetails{Trigger: false, Status: Stopped},
	}
}

func (jc *JobController) TriggerStartingJob() {
	jc.DetectChanges.Trigger = true
}
