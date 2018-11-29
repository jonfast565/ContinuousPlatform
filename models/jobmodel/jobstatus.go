package jobmodel

type JobStatus string

const (
	Stopped JobStatus = "Stopped"
	Running JobStatus = "Running"
	Errored JobStatus = "Errored"
)
