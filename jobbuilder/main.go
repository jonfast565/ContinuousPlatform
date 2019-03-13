// job builder main
package main

import (
	"github.com/ahmetb/go-linq"
	"github.com/gorilla/mux"
	"github.com/jonfast565/continuous-platform/constants"
	"github.com/jonfast565/continuous-platform/jobbuilder/server"
	"github.com/jonfast565/continuous-platform/models/jobmodel"
	"github.com/jonfast565/continuous-platform/utilities/jsonutil"
	"github.com/jonfast565/continuous-platform/utilities/logging"
	"github.com/jonfast565/continuous-platform/utilities/networking"
	"net/http"
	"strconv"
	"time"
)

var (
	configuration JobConfiguration
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
				if configuration.CycleRateLimiting {
					logging.LogInfo("Cycle rate limit: " +
						strconv.Itoa(configuration.CycleRateLimit) + "s")
					betweenJobDuration := time.Duration(configuration.CycleRateLimit)
					time.Sleep(betweenJobDuration * time.Second)
				}
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

	result := linq.From(controller.JobList).FirstWithT(func(jobDetail *jobmodel.JobDetails) bool {
		return jobDetail.Name == model.JobName
	})

	if result != nil {
		resultPtr := result.(*jobmodel.JobDetails)
		resultPtr.TriggerJob()
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
	}
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
