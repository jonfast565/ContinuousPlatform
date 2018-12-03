package jobmodel

type JobController struct {
	DetectChanges      JobDetails
	BuildDeliverables  JobDetails
	GenerateScripts    JobDetails
	DeployDebugScripts JobDetails
	DeployJenkinsJobs  JobDetails
}

func NewJobController() JobController {
	return JobController{
		DetectChanges:      JobDetails{Trigger: false, Status: Stopped},
		BuildDeliverables:  JobDetails{Trigger: false, Status: Stopped},
		GenerateScripts:    JobDetails{Trigger: false, Status: Stopped},
		DeployDebugScripts: JobDetails{Trigger: false, Status: Stopped},
		DeployJenkinsJobs:  JobDetails{Trigger: false, Status: Stopped},
	}
}

func (jc *JobController) TriggerStartingJob() {
	jc.DetectChanges.Trigger = true
}
