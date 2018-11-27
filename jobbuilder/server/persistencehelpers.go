package server

import (
	"../../clients/persistenceclient"
	"../../logging"
	"../../models/genmodel"
	"../../models/projectmodel"
	"../../models/repomodel"
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func SetRepositoriesCache(repoPackage repomodel.RepositoryPackage) error {
	client := persistenceclient.NewPersistenceClient()
	var packageBuffer bytes.Buffer
	enc := gob.NewEncoder(&packageBuffer)
	err := enc.Encode(repoPackage)
	if err != nil {
		return err
	}
	logging.LogInfo("Persisting repositories...")
	err = client.SetKeyValueCache("Repositories", packageBuffer.Bytes(), true)
	if err != nil {
		return err
	}
	return nil
}

func GetRepositoriesCache() (*repomodel.RepositoryPackage, error) {
	client := persistenceclient.NewPersistenceClient()
	logging.LogInfo("Getting repositories...")
	repositoriesBytes, err := client.GetKeyValueCache("Repositories", true)
	if err != nil {
		return nil, err
	}

	var value repomodel.RepositoryPackage
	packageBuffer := bytes.NewBuffer(repositoriesBytes)
	dec := gob.NewDecoder(packageBuffer)
	err = dec.Decode(&value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func SetDeliverablesCache(deliverablePackage projectmodel.DeliverablePackage) error {
	client := persistenceclient.NewPersistenceClient()
	logging.LogInfo("Persisting deliverables...")
	packageBuffer, err := json.Marshal(deliverablePackage)
	if err != nil {
		panic(err)
	}
	err = client.SetKeyValueCache("Deliverables", packageBuffer, true)
	if err != nil {
		panic(err)
	}
	return nil
}

func GetDeliverablesCache() (*projectmodel.DeliverablePackage, error) {
	client := persistenceclient.NewPersistenceClient()
	logging.LogInfo("Getting deliverables...")
	packageBytes, err := client.GetKeyValueCache("Deliverables", true)
	if err != nil {
		return nil, err
	}

	var value projectmodel.DeliverablePackage
	err = json.Unmarshal(packageBytes, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func SetScriptCache(scriptPackage genmodel.ScriptPackage) error {
	client := persistenceclient.NewPersistenceClient()
	logging.LogInfo("Persisting scripts...")
	packageBuffer, err := json.Marshal(scriptPackage)
	if err != nil {
		panic(err)
	}
	err = client.SetKeyValueCache("Scripts", packageBuffer, true)
	if err != nil {
		panic(err)
	}
	return nil
}
