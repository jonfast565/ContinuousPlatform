package server

import (
	"../../models/jobmodel"
)

func GenerateScripts(details *jobmodel.JobDetails) {
	defer func() {
		if r := recover(); r != nil {
			details.SetJobErrored()
		}
	}()

}
