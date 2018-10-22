package server

import (
	"../../models/jobmodel"
)

func DetectChanges(details *jobmodel.JobDetails) {
	defer func() {
		if r := recover(); r != nil {
			details.Status = jobmodel.Errored
		}
	}()
}
