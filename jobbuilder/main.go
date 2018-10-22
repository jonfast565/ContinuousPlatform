package main

import (
	"../constants"
	"../jsonutil"
	"../logging"
	"../models/jobmodel"
	"../networking"
	"./server"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
	controller = jobmodel.NewJobController()
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				logging.LogInfo("Beginning job cycle")

				if controller.DetectChanges.Trigger {
					logging.LogInfo("Detect changes")
					controller.DetectChanges.UnsetTriggerBeginRun()
					server.DetectChanges(&controller.DetectChanges)
					controller.DetectChanges.SetJobStoppedOrErrored()
					continue
				}

				if controller.BuildDeliverables.Trigger {
					logging.LogInfo("Build deliverables")
					controller.BuildDeliverables.UnsetTriggerBeginRun()
					server.BuildDeliverables(&controller.BuildDeliverables)
					controller.BuildDeliverables.SetJobStoppedOrErrored()
					continue
				}

				if controller.GenerateScripts.Trigger {
					logging.LogInfo("Generate scripts")
					controller.GenerateScripts.UnsetTriggerBeginRun()
					server.GenerateScripts(&controller.GenerateScripts)
					controller.GenerateScripts.SetJobStoppedOrErrored()
					continue
				}

				if controller.DeployJenkinsJobs.Trigger {
					logging.LogInfo("Deploy Jenkins jobs")
					controller.DeployJenkinsJobs.UnsetTriggerBeginRun()
					server.DeployJenkinsJobs(&controller.DeployJenkinsJobs)
					controller.DeployJenkinsJobs.SetJobStoppedOrErrored()
					continue
				}

				if configuration.CyclicalRuns {
					logging.LogInfo("Cyclical run enabled. Triggering starting job.")
					controller.TriggerStartingJob()
				}
				time.Sleep(2 * time.Second)
			}
		}
	}()

	router := mux.NewRouter()
	router.HandleFunc("/Daemon/GetJobDetails", getJobDetails).Methods(constants.PostMethod)

	localPort := networking.GetLocalPort(configuration.Port)
	logging.LogContentService(localPort)
	log.Fatal(http.ListenAndServe(localPort, router))

	quit <- true
	logging.LogApplicationEnd()
}

func getJobDetails(w http.ResponseWriter, r *http.Request) {
	resultBytes, err := jsonutil.EncodeJsonToBytes(&controller)
	if err != nil {
		w.WriteHeader(500)
		logging.LogError(err)
		return
	}

	w.Write(*resultBytes)
}
