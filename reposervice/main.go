package main

import (
	"github.com/gorilla/mux"
	"github.com/jonfast565/continuous-platform/constants"
	"github.com/jonfast565/continuous-platform/jsonutil"
	"github.com/jonfast565/continuous-platform/limitutil"
	"github.com/jonfast565/continuous-platform/logging"
	"github.com/jonfast565/continuous-platform/models/repomodel"
	"github.com/jonfast565/continuous-platform/networking"
	"github.com/jonfast565/continuous-platform/reposervice/server"
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
