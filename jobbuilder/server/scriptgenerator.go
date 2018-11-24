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

	/*
		persistenceClient := persistenceclient.NewPersistenceClient()
		// TODO: Generate scripts here

		buildInfrastructure, err := persistenceClient.GetBuildInfrastructure(key)
		if err != nil {
			panic(err)
		}
	*/

}
