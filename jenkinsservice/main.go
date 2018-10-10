package main

import (
	"../models/jenkins"
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
	var jobRequest jenkins.JobRequest
	err := utilities.DecodeJsonFromBody(r, &jobRequest)
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}

	crumb, err := endpoint.GetJenkinsCrumb()
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}

	_, err = endpoint.CreateUpdateJob(*crumb, jobRequest)
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}

	w.WriteHeader(200)
}

func createFolder(w http.ResponseWriter, r *http.Request) {
	var jobRequest jenkins.JobRequest
	err := utilities.DecodeJsonFromBody(r, &jobRequest)
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}

	crumb, err := endpoint.GetJenkinsCrumb()
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}

	_, err = endpoint.CreateFolder(*crumb, jobRequest)
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}

	w.WriteHeader(200)
}

func deleteJobOrFolder(w http.ResponseWriter, r *http.Request) {
	var jobRequest jenkins.JobRequest
	err := utilities.DecodeJsonFromBody(r, &jobRequest)
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}

	crumb, err := endpoint.GetJenkinsCrumb()
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}

	_, err = endpoint.DeleteJobOrFolder(*crumb, jobRequest)
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}

	w.WriteHeader(200)
}

func getJenkinsMetadata(w http.ResponseWriter, r *http.Request) {
	crumb, err := endpoint.GetJenkinsCrumb()
	if err != nil {
		w.WriteHeader(500)
		utilities.LogError(err)
		return
	}

	result, err := endpoint.GetJenkinsMetadata(*crumb)
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

func getJenkinsCrumb(w http.ResponseWriter, r *http.Request) {
	result, err := endpoint.GetJenkinsCrumb()
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
