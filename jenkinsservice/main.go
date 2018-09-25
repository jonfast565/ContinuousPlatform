package main

import (
	"../utilities"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var configuration JenkinsConfiguration

func main() {
	utilities.CreateLog()
	utilities.LogHeader("Jenkins Service")
	utilities.LogApplicationStart()

	utilities.DecodeJsonFromFile("./appsettings.json", &configuration)

	router := mux.NewRouter()
	router.HandleFunc("/Daemon/CreateUpdateJob", createUpdateJob).Methods(utilities.PostMethod)
	router.HandleFunc("/Daemon/CreateFolder", createFolder).Methods(utilities.PostMethod)
	router.HandleFunc("/Daemon/DeleteJobOrFolder", deleteJobOrFolder).Methods(utilities.PostMethod)
	router.HandleFunc("/Daemon/GetJenkinsMetadata", getJenkinsMetadata).Methods(utilities.PostMethod)
	router.HandleFunc("/Daemon/GetJenkinsCrumb", getJenkinsCrumb).Methods(utilities.PostMethod)

	localPort := utilities.GetLocalPort(1212) // TODO: Replace
	utilities.LogContentService(localPort)
	log.Fatal(http.ListenAndServe(localPort, router))
	utilities.LogApplicationEnd()
}

func createUpdateJob(w http.ResponseWriter, r *http.Request) {

}

func createFolder(w http.ResponseWriter, r *http.Request) {

}

func deleteJobOrFolder(w http.ResponseWriter, r *http.Request) {

}

func getJenkinsMetadata(w http.ResponseWriter, r *http.Request) {

}

func getJenkinsCrumb(w http.ResponseWriter, r *http.Request) {

}
