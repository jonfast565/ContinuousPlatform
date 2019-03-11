package main

import (
	"github.com/gorilla/mux"
	"github.com/jonfast565/continuous-platform/constants"
	"github.com/jonfast565/continuous-platform/models/inframodel"
	"github.com/jonfast565/continuous-platform/models/loggingmodel"
	"github.com/jonfast565/continuous-platform/models/persistmodel"
	"github.com/jonfast565/continuous-platform/persistenceservice/server"
	"github.com/jonfast565/continuous-platform/utilities/jsonutil"
	"github.com/jonfast565/continuous-platform/utilities/logging"
	"github.com/jonfast565/continuous-platform/utilities/networking"
	"net/http"
)

var (
	configuration server.PersistenceServiceConfiguration
	endpoint      *server.PersistenceServiceEndpoint
)

func main() {
	logging.CreateLog()
	logging.LogHeader("Persistence Service")
	logging.LogApplicationStart()

	jsonutil.DecodeJsonFromFile("./appsettings.json", &configuration)
	endpoint = server.NewPersistenceServiceEndpoint(configuration)

	router := mux.NewRouter()
	router.HandleFunc("/Daemon/GetKeyValueCache", getKeyValueCache).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/SetKeyValueCache", setKeyValueCache).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/SetLogRecord", setLogRecord).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/GetBuildInfrastructure", getBuildInfrastructure).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/GetResourceList", getResourceList).Methods(constants.PostMethod)

	localPort := networking.GetLocalPort(configuration.Port)
	logging.LogContentService(localPort)
	logging.LogFatal(http.ListenAndServe(localPort, router))
	logging.LogApplicationEnd()
}

func getKeyValueCache(w http.ResponseWriter, r *http.Request) {
	var model persistmodel.KeyRequest

	err := jsonutil.DecodeJsonFromBody(r, &model)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	result, err := endpoint.GetKeyValueCache(&model)
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

func setKeyValueCache(w http.ResponseWriter, r *http.Request) {
	var model persistmodel.KeyValueRequest

	err := jsonutil.DecodeJsonFromBody(r, &model)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	err = endpoint.SetKeyValueCache(&model)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	w.WriteHeader(200)
}

func getBuildInfrastructure(w http.ResponseWriter, r *http.Request) {
	var model inframodel.ResourceKey

	err := jsonutil.DecodeJsonFromBody(r, &model)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	result, err := endpoint.GetBuildInfrastructure(model)
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

func getResourceList(w http.ResponseWriter, r *http.Request) {
	result, err := endpoint.GetResourceCache()
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

func setLogRecord(w http.ResponseWriter, r *http.Request) {
	var model loggingmodel.LogRecord

	err := jsonutil.DecodeJsonFromBody(r, &model)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	err = endpoint.SetLogRecord(&model)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}
}
