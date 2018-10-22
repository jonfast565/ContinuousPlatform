package main

import (
	"../jsonutil"
	"../logging"
	"../models/jobmodel"
	"./server"
	"fmt"
	_ "github.com/gizak/termui"
	"time"
)

var (
	configuration Configuration
)

func main() {
	logging.CreateLog()
	logging.LogHeader("Job Builder")
	logging.LogApplicationStart()

	jsonutil.DecodeJsonFromFile("./appsettings.json", &configuration)
	controller := jobmodel.NewJobController()
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
					controller.DetectChanges.Status = jobmodel.Running
					server.DetectChanges(&controller.DetectChanges)
					if controller.BuildDeliverables.Status != jobmodel.Errored {
						controller.BuildDeliverables.Status = jobmodel.Stopped
					}
				}

				if controller.BuildDeliverables.Trigger {
					logging.LogInfo("Build deliverables")
					controller.BuildDeliverables.Status = jobmodel.Running
					server.BuildDeliverables(&controller.BuildDeliverables)
					if controller.BuildDeliverables.Status != jobmodel.Errored {
						controller.BuildDeliverables.Status = jobmodel.Stopped
					}
				}

				if controller.GenerateScripts.Trigger {
					logging.LogInfo("Generate scripts")
					controller.GenerateScripts.Status = jobmodel.Running
					server.GenerateScripts(&controller.GenerateScripts)
					if controller.GenerateScripts.Status != jobmodel.Errored {
						controller.GenerateScripts.Status = jobmodel.Stopped
					}
				}

				if controller.DeployJenkinsJobs.Trigger {
					logging.LogInfo("Deploy Jenkins jobs")
					controller.DeployJenkinsJobs.Status = jobmodel.Running
					server.DeployJenkinsJobs(&controller.DeployJenkinsJobs)
					if controller.DeployJenkinsJobs.Status != jobmodel.Errored {
						controller.DeployJenkinsJobs.Status = jobmodel.Stopped
					}
				}

				if configuration.CyclicalRuns {
					logging.LogInfo("Cyclical run")
				}
				time.Sleep(2 * time.Second)
			}
		}
	}()

	// replace with web methods for getting statuses and
	// triggering jobs to run and stopping jobs
	logging.LogInfo("Waiting for exiting key press...")
	fmt.Scanln()
	quit <- true
	logging.LogApplicationEnd()
}
