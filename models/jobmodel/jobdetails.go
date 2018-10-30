package jobmodel

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
