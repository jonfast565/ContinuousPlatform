package main

import (
	"../utilities"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	utilities.CreateLog()
	utilities.LogHeader("Jenkins Service")
	utilities.LogApplicationStart()

	// utilities.DecodeJsonFromFile("./appsettings.json", &configuration)

	router := mux.NewRouter()
	router.HandleFunc("/Daemon/CreateUpdateJob", createUpdateJob).Methods(utilities.PostMethod)
	router.HandleFunc("/Daemon/DeleteJob", deleteJob).Methods(utilities.PostMethod)

	localPort := utilities.GetLocalPort(1212) // TODO: Replace
	utilities.LogContentService(localPort)
	log.Fatal(http.ListenAndServe(localPort, router))
	utilities.LogApplicationEnd()
}

func createUpdateJob(w http.ResponseWriter, r *http.Request) {

}

func deleteJob(w http.ResponseWriter, r *http.Request) {

}
