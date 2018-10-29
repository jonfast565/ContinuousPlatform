package server

import (
	"../../logging"
	"../../models/jobmodel"
	"../../models/projectmodel"
	"./builders"
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

	// TODO: Needs to be a concurrent routine
	var dotNetDeliverables []projectmodel.DotNetDeliverable
	for _, repository := range repositories.Metadata {
		// only build graph once for multiple runs
		logging.LogInfo("Building repo graph for " + repository.Name + " b. " + repository.Branch)
		graph := repository.BuildGraph()

		// support .NET deliverables at this time
		logging.LogInfo("Building deliverables for " + repository.Name + " b. " + repository.Branch)
		deliverables, err := builders.BuildDotNetDeliverables(repository, *graph)
		if err != nil {
			panic(err)
		}
		dotNetDeliverables = append(dotNetDeliverables, deliverables...)
	}
}
