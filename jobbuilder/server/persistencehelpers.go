package server

import (
	"../../clients/persistenceclient"
	"../../logging"
	"../../models/repomodel"
	"encoding/json"
)

func SetRepositoriesCache(repoPackage repomodel.RepositoryPackage) error {
	client := persistenceclient.NewPersistenceClient()
	packageBytes, err := json.Marshal(repoPackage)
	if err != nil {
		return err
	}
	logging.LogInfo("Persisting repositories...")
	err = client.SetKeyValueCache("Repositories", packageBytes)
	if err != nil {
		return err
	}
	return nil
}

func GetRepositoriesCache() (*repomodel.RepositoryPackage, error) {
	client := persistenceclient.NewPersistenceClient()
	logging.LogInfo("Getting repositories...")
	repositoriesBytes, err := client.GetKeyValueCache("Repositories")
	if err != nil {
		return nil, err
	}

	var value repomodel.RepositoryPackage
	err = json.Unmarshal(repositoriesBytes, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}
