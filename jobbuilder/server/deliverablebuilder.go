package server

import (
	"../../logging"
	"../../models/jobmodel"
	"fmt"
)

func BuildDeliverables(details *jobmodel.JobDetails) {
	defer func() {
		if r := recover(); r != nil {
			details.SetJobErrored()
			logging.LogPanicRecover(r)
		}
	}()

	repositories, err := GetRepositoriesCache()
	if err != nil {
		panic(err)
	}

	for _, repository := range repositories.Metadata {
		// debug for now
		fmt.Println(repository)
	}
}
