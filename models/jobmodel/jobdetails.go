package jobmodel

import (
	"../../logging"
	"fmt"
	"sync/atomic"
)

type JobDetails struct {
	Name           string
	Trigger        bool
	Status         JobStatus
	Progress       int64
	TotalProgress  int64
	jobDescription string
	nextJob        *JobDetails
	runnable       func(jobDetails *JobDetails) bool
}

type JobDetailsList []*JobDetails

func NewJobDetails(
	name string,
	jobDescription string,
	nextJob *JobDetails,
	runnable func(jobDetails *JobDetails) bool) *JobDetails {
	return &JobDetails{
		Name:           name,
		Trigger:        false,
		Status:         Stopped,
		Progress:       0,
		TotalProgress:  0,
		jobDescription: jobDescription,
		nextJob:        nextJob,
		runnable:       runnable,
	}
}

func (jd *JobDetails) SetProgress(newValue int64) {
	atomic.StoreInt64(&jd.Progress, newValue)
}

func (jd *JobDetails) ResetProgress() {
	atomic.StoreInt64(&jd.Progress, 0)
	atomic.StoreInt64(&jd.TotalProgress, 0)
}

func (jd *JobDetails) IncrementTotalProgressBy(newValue int64) {
	atomic.AddInt64(&jd.TotalProgress, newValue)
}

func (jd *JobDetails) IncrementTotalProgress() {
	atomic.AddInt64(&jd.TotalProgress, 1)
}

func (jd *JobDetails) IncrementProgress() {
	atomic.AddInt64(&jd.Progress, 1)
}

func (jd *JobDetails) SetTotalProgress(newValue int64) {
	atomic.StoreInt64(&jd.TotalProgress, newValue)
}

func (jd *JobDetails) CompleteProgress() {
	totalProgress := atomic.LoadInt64(&jd.TotalProgress)
	atomic.StoreInt64(&jd.TotalProgress, totalProgress)
}

func (jd *JobDetails) GetCompletionPercentage() string {
	progress := atomic.LoadInt64(&jd.Progress)
	totalProgress := atomic.LoadInt64(&jd.TotalProgress)
	percentValue := int64((float64(progress) / float64(totalProgress)) * 100)
	return fmt.Sprintf("%d%%", percentValue)
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

func (jd *JobDetails) RunJob() bool {
	result := false
	if jd.Trigger {
		logging.LogInfo(jd.jobDescription)
		jd.UnsetTriggerBeginRun()
		result = jd.runnable(jd)
		jd.SetJobStoppedOrErrored()
		if !jd.Errored() {
			if jd.nextJob != nil {
				jd.nextJob.TriggerJob()
			}
			return result
		}
	}
	return false
}
