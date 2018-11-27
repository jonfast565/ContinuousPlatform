package jobmodel

import (
	"fmt"
	"sync/atomic"
)

type JobDetails struct {
	Trigger       bool
	Status        JobStatus
	Progress      int64
	TotalProgress int64
}

func (jd *JobDetails) SetProgress(newValue int64) {
	atomic.StoreInt64(&jd.Progress, newValue)
}

func (jd *JobDetails) ResetProgress() {
	atomic.StoreInt64(&jd.Progress, 0)
	atomic.StoreInt64(&jd.TotalProgress, 0)
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
