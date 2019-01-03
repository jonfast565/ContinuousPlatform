package main

import (
	"../constants"
	"../jsonutil"
	"../limitutil"
	"../logging"
	"../models/repomodel"
	"../networking"
	"./server"
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
	"net/http"
)

var (
	configuration server.TeamServicesConfiguration
	endpoint      *server.TeamServicesEndpoint
	rateLimiter   *rate.Limiter
)

func main() {
	logging.CreateLog()
	logging.LogHeader("Repo Service")
	logging.LogApplicationStart()

	jsonutil.DecodeJsonFromFile("./appsettings.json", &configuration)
	endpoint = server.NewTeamServicesEndpoint(configuration)
	rateLimiter = limitutil.NewRateLimiter(2, 10)

	router := mux.NewRouter()
	router.HandleFunc("/Daemon/GetRepositories", getRepositories).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/GetFile", getFile).Methods(constants.PostMethod)

	localPort := networking.GetLocalPort(configuration.Port)
	logging.LogContentService(localPort)
	logging.LogFatal(http.ListenAndServe(localPort, router))
	logging.LogApplicationEnd()
}

func getRepositories(w http.ResponseWriter, r *http.Request) {
	limitutil.WaitForAllow(rateLimiter)

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

	_, err = w.Write(*resultBytes)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
	}
}

func getFile(w http.ResponseWriter, r *http.Request) {
	limitutil.WaitForAllow(rateLimiter)

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

	_, err = w.Write(*resultBytes)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
	}
}
