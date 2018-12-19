package jobmodel

type JobController struct {
	JobList JobDetailsList
}

func NewJobController(jobs JobDetailsList) JobController {
	return JobController{
		JobList: jobs,
	}
}

func (jc *JobController) TriggerStartingJob() {
	jc.JobList[0].Trigger = true
}

func (jc JobController) RunSequence() {
	for _, job := range jc.JobList {
		result := job.RunJob()
		if !result {
			break
		}
	}
}
