package server

import (
	"../../clients/persistenceclient"
	"../../clients/repoclient"
	"../../logging"
	"../../models/jobmodel"
	"../../models/repomodel"
	"encoding/json"
	"github.com/ahmetb/go-linq"
	"strconv"
)

func SourceControlChangesExist(details *jobmodel.JobDetails) bool {
	defer func() {
		if r := recover(); r != nil {
			details.SetJobErrored()
			logging.LogPanicRecover(r)
		}
	}()

	oldPackage, err := GetRepositoriesCache()
	if err != nil {
		oldPackage = &repomodel.RepositoryPackage{}
	}

	repoClient := repoclient.NewRepoClient()
	newPackage, err := repoClient.GetRepositories()

	if len(oldPackage.Metadata) != len(newPackage.Metadata) {
		logging.LogInfoMultiline("Repo Count Mismatch",
			"Old Package Ct.: "+strconv.Itoa(len(oldPackage.Metadata)),
			"New Package Ct.: "+strconv.Itoa(len(newPackage.Metadata)))
		return true
	}

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
			return true
		}
	}

	return false
}

func GetRepositoriesCache() (*repomodel.RepositoryPackage, error) {
	client := persistenceclient.NewPersistenceClient()
	repositoriesBytes, err := client.GetKeyValueCache("Repositories")
	if err != nil {
		return nil, err
	}

	var value repomodel.RepositoryPackage
	err = json.Unmarshal(repositoriesBytes, value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}
