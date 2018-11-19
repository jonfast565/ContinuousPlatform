package server

import (
	"../../clients/infraclient"
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

	infraClient := infraclient.NewInfraClient()
	metadata, err := infraClient.GetInfrastructure()
	if err != nil {
		panic(err)
	}

	// TODO: Generate scripts here

}
