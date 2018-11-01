package main

import (
	"../constants"
	"../jsonutil"
	"../logging"
	"../models/repomodel"
	"../networking"
	"./server"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	configuration server.TeamServicesConfiguration
	endpoint      *server.TeamServicesEndpoint
)

func main() {
	logging.CreateLog()
	logging.LogHeader("Repo Service")
	logging.LogApplicationStart()

	jsonutil.DecodeJsonFromFile("./appsettings.json", &configuration)
	endpoint = server.NewTeamServicesEndpoint(configuration)

	router := mux.NewRouter()
	router.HandleFunc("/Daemon/GetRepositories", getRepositories).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/GetFile", getFile).Methods(constants.PostMethod)

	localPort := networking.GetLocalPort(configuration.Port)
	logging.LogContentService(localPort)
	logging.LogFatal(http.ListenAndServe(localPort, router))
	logging.LogApplicationEnd()
}

func getRepositories(w http.ResponseWriter, r *http.Request) {
	result, err := endpoint.GetRepositories()
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}
	resultBytes, err := jsonutil.EncodeJsonToBytes(&result)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}
	w.Write(*resultBytes)
}

func getFile(w http.ResponseWriter, r *http.Request) {
	var repositoryFileMetadata repomodel.RepositoryFileMetadata
	err := jsonutil.DecodeJsonFromBody(r, &repositoryFileMetadata)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}
	result, err := endpoint.GetFile(repositoryFileMetadata)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}
	resultBytes, err := jsonutil.EncodeJsonToBytes(&result)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}
	w.Write(*resultBytes)
}
