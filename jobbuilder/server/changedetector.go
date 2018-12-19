package server

import (
	"../../clients/repoclient"
	"../../logging"
	"../../models/jobmodel"
	"../../models/repomodel"
	"github.com/ahmetb/go-linq"
	"strconv"
)

func DetectChanges(details *jobmodel.JobDetails) bool {
	defer func() {
		if r := recover(); r != nil {
			details.SetJobErrored()
			logging.LogPanicRecover(r)
		}
	}()

	details.ResetProgress()
	oldPackage, err := GetRepositoriesCache()
	if err != nil {
		logging.LogInfo("Got nothing back, starting from scratch")
		oldPackage = repomodel.NewRepositoryPackage()
	}

	repoClient := repoclient.NewRepoClient()
	newPackage, err := repoClient.GetRepositories()
	if err != nil {
		panic(err)
	}

	changeFlag := false
	logging.LogInfo("Checking changes...")
	if len(oldPackage.Metadata) != len(newPackage.Metadata) {
		logging.LogInfoMultiline("Repo Count Mismatch",
			"Old Package Ct.: "+strconv.Itoa(len(oldPackage.Metadata)),
			"New Package Ct.: "+strconv.Itoa(len(newPackage.Metadata)))
		changeFlag = true
	}

	if changeFlag {
		logging.LogInfo("Repo count changed")
		err = SetRepositoriesCache(*newPackage)
		if err != nil {
			panic(err)
		}
		return true
	}

	details.SetTotalProgress(int64(len(newPackage.Metadata)))
	for _, newRepo := range newPackage.Metadata {
		logging.LogInfo("Process repo " +
			newRepo.Name + " b. " +
			newRepo.Branch)

		oldRepo := linq.From(oldPackage.Metadata).FirstWithT(
			func(r repomodel.RepositoryMetadata) bool {
				return r.Name == newRepo.Name &&
					r.Branch == newRepo.Branch
			})

		if oldRepo == nil {
			logging.LogInfo("Repo " +
				newRepo.Name + " b. " +
				newRepo.Branch + " not found")
			changeFlag = true
			details.CompleteProgress()
			break
		} else {
			details.IncrementProgress()
		}
	}

	if changeFlag {
		logging.LogInfo("Branch structure changed")
		err = SetRepositoriesCache(*newPackage)
		if err != nil {
			panic(err)
		}
		return true
	}

	return false
}
