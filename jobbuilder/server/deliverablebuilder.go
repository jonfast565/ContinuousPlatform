package server

import (
	"../../logging"
	"../../models/jobmodel"
	"../../models/projectmodel"
	"../../models/repomodel"
	"./builders"
	"runtime"
	"sync"
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

	resultsChannel := make(chan projectmodel.Deliverable)
	var wg sync.WaitGroup
	var results []projectmodel.Deliverable
	wg.Add(len(repositories.Metadata))

	for _, repository := range repositories.Metadata {
		go buildDeliverable(repository, &wg, resultsChannel)
	}

	wg.Wait()
	for {
		noMore := false
		select {
		case msg := <-resultsChannel:
			results = append(results, msg)
		default:
			logging.LogInfo("No more deliverables received")
			noMore = true
		}
		if noMore {
			break
		}
	}

	deliverablePackage := projectmodel.DeliverablePackage{Deliverables: results}
	runtime.GC()
	SetDeliverablesCache(deliverablePackage)
}

func buildDeliverable(repository repomodel.RepositoryMetadata,
	wg *sync.WaitGroup,
	resultsChan chan projectmodel.Deliverable) {
	defer wg.Done()

	logging.LogInfo("Building deliverables for " + repository.Name + " b. " + repository.Branch)

	// support .NET deliverables at this time
	deliverables, err := builders.BuildDotNetDeliverables(repository)
	if err != nil {
		panic(err)
	}

	// TODO: Support Golang deliverables, like this project
	// TODO: Support Ruby deliverables, like mobile projects
	// TODO: Support NodeJS/NPM deliverables, for the status dashboard, etc.
	// TODO: Support other deliverable contexts if necessary (Haskell, etc.)

	deliverable := projectmodel.Deliverable{
		DotNetDeliverables: deliverables,
	}

	go func() { resultsChan <- deliverable }()
}
