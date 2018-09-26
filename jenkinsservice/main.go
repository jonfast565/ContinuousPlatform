package main

import (
	"../utilities"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var configuration JenkinsConfiguration
var endpoint JenkinsEndpoint

func main() {
	utilities.CreateLog()
	utilities.LogHeader("Jenkins Service")
	utilities.LogApplicationStart()

	utilities.DecodeJsonFromFile("./appsettings.json", &configuration)

	router := mux.NewRouter()
	endpoint = NewJenkinsEndpoint(configuration)
	router.HandleFunc("/Daemon/CreateUpdateJob", createUpdateJob).Methods(utilities.PostMethod)
	router.HandleFunc("/Daemon/CreateFolder", createFolder).Methods(utilities.PostMethod)
	router.HandleFunc("/Daemon/DeleteJobOrFolder", deleteJobOrFolder).Methods(utilities.PostMethod)
	router.HandleFunc("/Daemon/GetJenkinsMetadata", getJenkinsMetadata).Methods(utilities.PostMethod)
	router.HandleFunc("/Daemon/GetJenkinsCrumb", getJenkinsCrumb).Methods(utilities.PostMethod)

	localPort := utilities.GetLocalPort(configuration.Port)
	utilities.LogContentService(localPort)
	log.Fatal(http.ListenAndServe(localPort, router))
	utilities.LogApplicationEnd()
}

func createUpdateJob(w http.ResponseWriter, r *http.Request) {
	result, err := endpoint.CreateUpdateJob()
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}
}

func createFolder(w http.ResponseWriter, r *http.Request) {
	result, err := endpoint.CreateFolder()
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}
}

func deleteJobOrFolder(w http.ResponseWriter, r *http.Request) {
	result, err := endpoint.DeleteJobOrFolder()
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}
}

func getJenkinsMetadata(w http.ResponseWriter, r *http.Request) {
	result, err := endpoint.GetJenkinsMetadata()
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}
}

func getJenkinsCrumb(w http.ResponseWriter, r *http.Request) {
	result, err := endpoint.GetJenkinsCrumb()
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}
	w.Write(result)
}
