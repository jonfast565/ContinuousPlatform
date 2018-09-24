package main

import (
	"../models/repos"
	"../utilities"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var configuration TeamServicesConfiguration
var endpoint *TeamServicesEndpoint

func main() {
	utilities.CreateLog()
	utilities.LogHeader("RepoService")
	utilities.LogApplicationStart()

	utilities.DecodeJsonFromFile("./appsettings.json", &configuration)
	endpoint = NewTeamServicesEndpoint(configuration)

	router := mux.NewRouter()
	router.HandleFunc("/Daemon/GetRepositories", getRepositories).Methods(utilities.PostMethod)
	router.HandleFunc("/Daemon/GetFile", getFile).Methods(utilities.PostMethod)

	localPort := utilities.GetLocalPort(configuration.Port)
	utilities.LogContentService(localPort)
	log.Fatal(http.ListenAndServe(localPort, router))
	utilities.LogApplicationEnd()
}

func getRepositories(w http.ResponseWriter, r *http.Request) {
	result, err := endpoint.GetRepositories()
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}
	resultBytes, err := utilities.EncodeJsonToBytes(&result)
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}
	w.Write(*resultBytes)
}

func getFile(w http.ResponseWriter, r *http.Request) {
	var repositoryFileMetadata repos.RepositoryFileMetadata
	err := utilities.DecodeJsonFromBody(r, &repositoryFileMetadata)
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}
	result, err := endpoint.GetFile(repositoryFileMetadata)
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}
	resultBytes, err := utilities.EncodeJsonToBytes(&result)
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}
	w.Write(*resultBytes)
}
