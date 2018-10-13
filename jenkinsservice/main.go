package main

import (
	"../constants"
	"../jsonutil"
	"../logging"
	"../models"
	"../networking"
	"./server"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var configuration server.JenkinsConfiguration
var endpoint server.JenkinsEndpoint

func main() {
	logging.CreateLog()
	logging.LogHeader("Jenkins Service")
	logging.LogApplicationStart()

	jsonutil.DecodeJsonFromFile("./appsettings.json", &configuration)

	router := mux.NewRouter()
	endpoint = server.NewJenkinsEndpoint(configuration)
	router.HandleFunc("/Daemon/GetJenkinsMetadata", getJenkinsMetadata).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/GetJenkinsCrumb", getJenkinsCrumb).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/CreateUpdateJob", createUpdateJob).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/CheckJob", checkJob).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/CreateFolder", createFolder).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/DeleteJobOrFolder", deleteJobOrFolder).Methods(constants.PostMethod)

	localPort := networking.GetLocalPort(configuration.Port)
	logging.LogContentService(localPort)
	log.Fatal(http.ListenAndServe(localPort, router))
	logging.LogApplicationEnd()
}

func checkJob(w http.ResponseWriter, r *http.Request) {
	var jobRequest models.JenkinsJobRequest
	err := jsonutil.DecodeJsonFromBody(r, &jobRequest)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	crumb, err := endpoint.GetJenkinsCrumb()
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	_, err = endpoint.CheckJobExistence(*crumb, jobRequest)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	w.WriteHeader(200)
}

func createUpdateJob(w http.ResponseWriter, r *http.Request) {
	var jobRequest models.JenkinsJobRequest
	err := jsonutil.DecodeJsonFromBody(r, &jobRequest)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	crumb, err := endpoint.GetJenkinsCrumb()
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	_, err = endpoint.CreateUpdateJob(*crumb, jobRequest)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	w.WriteHeader(200)
}

func createFolder(w http.ResponseWriter, r *http.Request) {
	var jobRequest models.JenkinsJobRequest
	err := jsonutil.DecodeJsonFromBody(r, &jobRequest)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	crumb, err := endpoint.GetJenkinsCrumb()
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	_, err = endpoint.CreateFolder(*crumb, jobRequest)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	w.WriteHeader(200)
}

func deleteJobOrFolder(w http.ResponseWriter, r *http.Request) {
	var jobRequest models.JenkinsJobRequest
	err := jsonutil.DecodeJsonFromBody(r, &jobRequest)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	crumb, err := endpoint.GetJenkinsCrumb()
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	_, err = endpoint.DeleteJobOrFolder(*crumb, jobRequest)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	w.WriteHeader(200)
}

func getJenkinsMetadata(w http.ResponseWriter, r *http.Request) {
	crumb, err := endpoint.GetJenkinsCrumb()
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	result, err := endpoint.GetJenkinsMetadata(*crumb)
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

func getJenkinsCrumb(w http.ResponseWriter, r *http.Request) {
	result, err := endpoint.GetJenkinsCrumb()
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
