package main

import (
	"../constants"
	"../jsonutil"
	"../logging"
	"../models/jenkinsmodel"
	"../networking"
	"./server"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	configuration server.JenkinsConfiguration
	endpoint      server.JenkinsEndpoint
)

func main() {
	logging.CreateLog()
	logging.LogHeader("Jenkins Service")
	logging.LogApplicationStart()

	jsonutil.DecodeJsonFromFile("./appsettings.json", &configuration)

	router := mux.NewRouter()
	endpoint = server.NewJenkinsEndpoint(configuration)
	router.HandleFunc("/Daemon/GetJenkinsMetadata", getJenkinsMetadata).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/GetJenkinsCrumb", getJenkinsCrumb).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/CreateJob", createJob).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/UpdateJob", updateJob).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/CheckJob", checkJob).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/CreateFolder", createFolder).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/DeleteJobOrFolder", deleteJobOrFolder).Methods(constants.PostMethod)

	localPort := networking.GetLocalPort(configuration.Port)
	logging.LogContentService(localPort)
	logging.LogFatal(http.ListenAndServe(localPort, router))
	logging.LogApplicationEnd()
}

func checkJob(w http.ResponseWriter, r *http.Request) {
	var jobRequest jenkinsmodel.JenkinsJobRequest
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

func createJob(w http.ResponseWriter, r *http.Request) {
	var jobRequest jenkinsmodel.JenkinsJobRequest
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

	_, err = endpoint.CreateJob(*crumb, jobRequest)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	w.WriteHeader(200)
}

func updateJob(w http.ResponseWriter, r *http.Request) {
	var jobRequest jenkinsmodel.JenkinsJobRequest
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

	_, err = endpoint.UpdateJob(*crumb, jobRequest)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	w.WriteHeader(200)
}

func createFolder(w http.ResponseWriter, r *http.Request) {
	var jobRequest jenkinsmodel.JenkinsJobRequest
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
	var jobRequest jenkinsmodel.JenkinsJobRequest
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

	_, err = w.Write(*resultBytes)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
	}
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

	_, err = w.Write(*resultBytes)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
	}
}
