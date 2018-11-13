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

	_, err := GetDeliverablesCache()
	if err != nil {
		panic(err)
	}

	// TODO: Get infrastructure data here

	// TODO: Generate scripts here
}
