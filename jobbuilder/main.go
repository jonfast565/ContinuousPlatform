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
	controller = jobmodel.NewJobController()
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

				if controller.DetectChanges.Trigger {
					logging.LogInfo("Detect changes")
					controller.DetectChanges.UnsetTriggerBeginRun()
					changesExist := server.SourceControlChangesExist(&controller.DetectChanges)
					controller.DetectChanges.SetJobStoppedOrErrored()

					if !controller.DetectChanges.Errored() {
						if changesExist {
							logging.LogInfo("Changes detected. Building deliverables")
							controller.BuildDeliverables.TriggerJob()
						} else {
							if configuration.ProceedDespiteNoChanges {
								logging.LogInfo("Change override set. Proceeding despite no changes")
								controller.BuildDeliverables.TriggerJob()
							} else {
								logging.LogInfo("No changes detected. Initiating change cycle again")
								controller.DetectChanges.TriggerJob()

								if configuration.ChangeRateLimiting {
									logging.LogInfo("Rate limit wait time: " +
										strconv.Itoa(configuration.ChangeRateLimit) + "s")
									rateLimitDuration := time.Duration(configuration.ChangeRateLimit)
									time.Sleep(rateLimitDuration * time.Second)
								}
							}
						}
					}
					continue
				}

				if controller.BuildDeliverables.Trigger {
					logging.LogInfo("Build deliverables")
					controller.BuildDeliverables.UnsetTriggerBeginRun()
					server.BuildDeliverables(&controller.BuildDeliverables)
					controller.BuildDeliverables.SetJobStoppedOrErrored()

					if !controller.BuildDeliverables.Errored() {
						controller.GenerateScripts.TriggerJob()
					}
					continue
				}

				if controller.GenerateScripts.Trigger {
					logging.LogInfo("Generate scripts")
					controller.GenerateScripts.UnsetTriggerBeginRun()
					server.GenerateScripts(&controller.GenerateScripts)
					controller.GenerateScripts.SetJobStoppedOrErrored()

					if !controller.GenerateScripts.Errored() {
						controller.DeployDebugScripts.TriggerJob()
					}
					continue
				}

				if controller.DeployDebugScripts.Trigger {
					logging.LogInfo("Deploy scripts for debugging")
					controller.DeployDebugScripts.UnsetTriggerBeginRun()
					server.DeployScriptsForDebugging(configuration.DebugBasePath, &controller.DeployDebugScripts)
					controller.DeployDebugScripts.SetJobStoppedOrErrored()

					if !controller.DeployDebugScripts.Errored() {
						controller.DeployJenkinsJobs.TriggerJob()
					}
					continue
				}

				if controller.DeployJenkinsJobs.Trigger {
					logging.LogInfo("Deploy Jenkins jobs")
					controller.DeployJenkinsJobs.UnsetTriggerBeginRun()
					// server.DeployJenkinsJobs(&controller.DeployJenkinsJobs)
					controller.DeployJenkinsJobs.SetJobStoppedOrErrored()

					if !controller.DeployJenkinsJobs.Errored() &&
						configuration.CyclicalRuns {
						logging.LogInfo("Cyclical run enabled. Triggering starting job.")
						controller.TriggerStartingJob()
					}
					continue
				}

				logging.LogInfo("Between job wait: " +
					strconv.Itoa(configuration.BetweenJobWait) + "s")
				betweenJobDuration := time.Duration(configuration.BetweenJobWait)
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
