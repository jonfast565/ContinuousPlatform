package main

import (
	"../constants"
	"../jsonutil"
	"../logging"
	"../models/jobmodel"
	"../networking"
	"./server"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

var (
	configuration Configuration
	controller    jobmodel.JobController
)

func main() {
	logging.CreateLog()
	logging.LogHeader("Job Builder")
	logging.LogApplicationStart()

	jsonutil.DecodeJsonFromFile("./appsettings.json", &configuration)

	jobs := initializeJobDetailsList()

	controller = jobmodel.NewJobController(jobs)
	if configuration.RunJobsOnStartup {
		controller.TriggerStartingJob()
	}

	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				logging.LogInfo("Beginning job cycle")
				controller.RunSequence()
				logging.LogInfo("Between job wait: " +
					strconv.Itoa(configuration.BetweenCycleWait) + "s")
				betweenJobDuration := time.Duration(configuration.BetweenCycleWait)
				time.Sleep(betweenJobDuration * time.Second)
			}
		}
	}()

	router := mux.NewRouter()
	router.HandleFunc("/Daemon/GetJobDetails", getJobDetails).Methods(constants.PostMethod)
	router.HandleFunc("/Daemon/TriggerJob", triggerJob).Methods(constants.PostMethod)

	localPort := networking.GetLocalPort(configuration.Port)
	logging.LogContentService(localPort)
	logging.LogFatal(http.ListenAndServe(localPort, router))

	quit <- true
	logging.LogApplicationEnd()
}

func initializeJobDetailsList() jobmodel.JobDetailsList {
	deployJenkinsJobsJob := jobmodel.NewJobDetails(
		"DeployJenkinsJobs",
		"Deploy Jenkins jobs",
		nil,
		server.DeployJenkinsJobs)
	deployDebugScriptsJob := jobmodel.NewJobDetails(
		"DeployDebugScripts",
		"Deploy debug scripts",
		deployJenkinsJobsJob,
		server.DeployScriptsForDebugging)
	generateScriptsJob := jobmodel.NewJobDetails(
		"GenerateScripts",
		"Generate scripts",
		deployDebugScriptsJob,
		server.GenerateScripts)
	buildDeliverablesJob := jobmodel.NewJobDetails(
		"BuildDeliverables",
		"Build deliverables",
		generateScriptsJob,
		server.BuildDeliverables)
	changeDetectorJob := jobmodel.NewJobDetails(
		"DetectChanges",
		"Detect changes",
		buildDeliverablesJob,
		server.DetectChanges)
	jobs := jobmodel.JobDetailsList{
		changeDetectorJob,
		buildDeliverablesJob,
		generateScriptsJob,
		deployDebugScriptsJob,
		deployJenkinsJobsJob,
	}
	return jobs
}

func triggerJob(w http.ResponseWriter, r *http.Request) {
	var model jobmodel.JobTrigger

	err := jsonutil.DecodeJsonFromBody(r, &model)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	switch model.JobName {
	case "DetectChanges":
		controller.DetectChanges.TriggerJob()
		break
	case "BuildDeliverables":
		controller.BuildDeliverables.TriggerJob()
		break
	case "GenerateScripts":
		controller.GenerateScripts.TriggerJob()
		break
	case "DeployDebugScripts":
		controller.DeployDebugScripts.TriggerJob()
		break
	case "DeployJenkinsJobs":
		controller.DeployJenkinsJobs.TriggerJob()
		break
	default:
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}

func getJobDetails(w http.ResponseWriter, r *http.Request) {
	resultBytes, err := jsonutil.EncodeJsonToBytes(&controller)
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
