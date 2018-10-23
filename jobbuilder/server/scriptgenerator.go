package server

import (
	"../../logging"
	"../../models/jobmodel"
)

func GenerateScripts(details *jobmodel.JobDetails) {
	defer func() {
		if r := recover(); r != nil {
			details.SetJobErrored()
			logging.LogPanicRecover(r)
		}
	}()
}
