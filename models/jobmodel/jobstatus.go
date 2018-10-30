package jobmodel

type JobStatus int

const (
	Stopped JobStatus = 0
	Running JobStatus = 1
	Errored JobStatus = 2
)
