package server

import (
	"github.com/jonfast565/continuous-platform/jobbuilder/server/builders"
	"github.com/jonfast565/continuous-platform/models/jobmodel"
	"github.com/jonfast565/continuous-platform/models/projectmodel"
	"github.com/jonfast565/continuous-platform/models/repomodel"
	"github.com/jonfast565/continuous-platform/utilities/logging"
	"runtime"
	"sync"
)

func BuildDeliverables(details *jobmodel.JobDetails) bool {
	defer func() {
		if r := recover(); r != nil {
			details.SetJobErrored()
			logging.LogPanicRecover(r)
		}
	}()

	details.ResetProgress()
	repositories, err := GetRepositoriesCache()
	if err != nil {
		panic(err)
	}

	resultsChannel := make(chan projectmodel.Deliverable)
	var wg sync.WaitGroup
	var results []projectmodel.Deliverable
	metadataLength := len(repositories.Metadata)
	wg.Add(metadataLength)
	details.SetTotalProgress(int64(metadataLength))

	for _, repository := range repositories.Metadata {
		go buildDeliverable(repository, &wg, details, resultsChannel)
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

	details.CompleteProgress()
	deliverablePackage := projectmodel.DeliverablePackage{Deliverables: results}
	runtime.GC()

	err = SetDeliverablesCache(deliverablePackage)
	if err != nil {
		panic(err)
	}

	return true
}

func buildDeliverable(repository repomodel.RepositoryMetadata,
	wg *sync.WaitGroup,
	details *jobmodel.JobDetails,
	resultsChan chan projectmodel.Deliverable) {
	// log both the failure of this goroutine (job) and the failure of the whole job
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			details.SetJobErrored()
			logging.LogPanicRecover(r)
		}
	}()

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

	details.IncrementProgress()
	go func() { resultsChan <- deliverable }()
}
