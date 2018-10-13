package main

import (
	"../constants"
	"../jsonutil"
	"../logging"
	"../networking"
	"./server"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var configuration server.PersistenceServiceConfiguration
var endpoint *server.PersistenceServiceEndpoint

func main() {
	logging.CreateLog()
	logging.LogHeader("Persistence Service")
	logging.LogApplicationStart()

	jsonutil.DecodeJsonFromFile("./appsettings.json", &configuration)
	endpoint = server.NewPersistenceServiceEndpoint(configuration)

	router := mux.NewRouter()
	router.HandleFunc("/Daemon/GetStuff", getStuff).Methods(constants.PostMethod)

	localPort := networking.GetLocalPort(configuration.Port)
	logging.LogContentService(localPort)
	log.Fatal(http.ListenAndServe(localPort, router))
	logging.LogApplicationEnd()
}

func getStuff(w http.ResponseWriter, r *http.Request) {
	/*
		var model models.Model
		err := jsonutil.DecodeJsonFromBody(r, &model)
		if err != nil {
			w.WriteHeader(500)
			logging.LogError(err)
			return
		}
		result, err := endpoint.something(something)
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
	*/
}
