package server

import (
	"../../models/jobmodel"
)

func BuildDeliverables(details *jobmodel.JobDetails) {
	defer func() {
		if r := recover(); r != nil {
			details.Status = jobmodel.Errored
		}
	}()
}
