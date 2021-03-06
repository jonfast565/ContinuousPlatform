package server

import (
	"github.com/ahmetb/go-linq"
	"github.com/jonfast565/continuous-platform/clients/jenkinsclient"
	"github.com/jonfast565/continuous-platform/constants"
	"github.com/jonfast565/continuous-platform/models/genmodel"
	"github.com/jonfast565/continuous-platform/models/jenkinsmodel"
	"github.com/jonfast565/continuous-platform/models/jobmodel"
	"github.com/jonfast565/continuous-platform/utilities/logging"
	"github.com/jonfast565/continuous-platform/utilities/stringutil"
	"sort"
)

// Deploys jobs to jenkins based on what scripts have been generated
func DeployJenkinsJobs(details *jobmodel.JobDetails) bool {
	defer func() {
		if r := recover(); r != nil {
			details.SetJobErrored()
			logging.LogPanicRecover(r)
		}
	}()

	details.ResetProgress()
	scripts, err := GetScriptCache()
	if err != nil {
		panic(err)
	}

	jenkinsClient := jenkinsclient.NewJenkinsClient()
	jenkinsMetadata, err := jenkinsClient.GetJenkinsMetadata()
	if err != nil {
		panic(err)
	}

	myKeys := buildKeyListFromScripts(scripts)
	jenkinsKeys := jenkinsMetadata.GetFlattenedKeys()
	edits := buildEditList(myKeys, jenkinsKeys, scripts)
	persistEditList(details, edits, jenkinsClient)

	return true
}

func persistEditList(
	details *jobmodel.JobDetails,
	edits jenkinsmodel.JenkinsEditList,
	jenkinsClient jenkinsclient.JenkinsClient) {
	details.SetTotalProgress(int64(len(edits)))
	for i, edit := range edits {
		details.SetProgress(int64(i + 1))
		jobRequest := edit.GetJobRequest()
		exists, err := jenkinsClient.CheckJobExists(jobRequest)
		if err != nil {
			panic(err)
		}
		switch edit.EditType {
		case jenkinsmodel.UpdateJob:
			if *exists {
				logging.LogInfo("Update Job: " + jobRequest.GetJobFragmentUrl())
				_, err = jenkinsClient.UpdateJob(jobRequest)
				if err != nil {
					panic(err)
				}
			} else {
				logging.LogInfo("Don't Update Job: " + jobRequest.GetJobFragmentUrl())
			}
			break
		case jenkinsmodel.AddJob:
			if !*exists {
				logging.LogInfo("Add Job: " + jobRequest.GetJobFragmentUrl())
				_, err := jenkinsClient.CreateJob(jobRequest)
				if err != nil {
					panic(err)
				}
			} else {
				logging.LogInfo("Update Job: " + jobRequest.GetJobFragmentUrl())
				_, err = jenkinsClient.UpdateJob(jobRequest)
				if err != nil {
					panic(err)
				}
			}
			break
		case jenkinsmodel.AddFolder:
			if !*exists {
				logging.LogInfo("Add Folder: " + jobRequest.GetJobFragmentUrl())
				_, err := jenkinsClient.CreateFolder(jobRequest)
				if err != nil {
					panic(err)
				}
			} else {
				logging.LogInfo("Don't Add Folder: " + jobRequest.GetJobFragmentUrl())
			}
			break
		case jenkinsmodel.RemoveJobFolder:
			if *exists {
				logging.LogInfo("Remove Job/Folder: " + jobRequest.GetJobFragmentUrl())
				_, err := jenkinsClient.DeleteJobOrFolder(jobRequest)
				if err != nil {
					panic(err)
				}
			} else {
				logging.LogInfo("Don't Remove Job/Folder: " + jobRequest.GetJobFragmentUrl())
			}
			break
		default:
			panic("not an option")
		}
	}
}

func buildKeyListFromScripts(scripts *genmodel.ScriptPackage) *jenkinsmodel.JenkinsJobKeyList {
	myMetadataKeys := make(jenkinsmodel.JenkinsJobKeyList, 0)
	for _, script := range scripts.Scripts {
		isJenkinsScript := stringutil.StringArrayContains(script.ToolScope, constants.JenkinsToolName)
		if !isJenkinsScript {
			continue
		}
		keySet := script.GetJenkinsKeyList()
		keySet.SanitizeKeyList()
		for _, key := range keySet {
			if !myMetadataKeys.KeyAlreadyExists(key.Keys) {
				myMetadataKeys = append(myMetadataKeys, key)
			}
		}
	}
	sort.Sort(myMetadataKeys)
	return &myMetadataKeys
}

func buildEditList(
	myKeys *jenkinsmodel.JenkinsJobKeyList,
	jenkinsKeys *jenkinsmodel.JenkinsJobKeyList,
	scripts *genmodel.ScriptPackage) jenkinsmodel.JenkinsEditList {

	addUpdateResults := getAddUpdateResults(myKeys, jenkinsKeys, scripts)
	deleteResults := getDeleteEdits(jenkinsKeys, myKeys)

	sort.Sort(addUpdateResults)
	sort.Sort(sort.Reverse(deleteResults))

	results := append(addUpdateResults, deleteResults...)
	return results
}

func getAddUpdateResults(
	myKeys *jenkinsmodel.JenkinsJobKeyList,
	jenkinsKeys *jenkinsmodel.JenkinsJobKeyList,
	scripts *genmodel.ScriptPackage) jenkinsmodel.JenkinsEditList {
	var addUpdateResults jenkinsmodel.JenkinsEditList
	for _, myKey := range *myKeys {
		found := false
		for _, jenkinsKey := range *jenkinsKeys {
			keyMatch := stringutil.StringArrayCompare(myKey.Keys, jenkinsKey.Keys) &&
				myKey.Type == jenkinsKey.Type
			if keyMatch {
				found = true
				break
			}
		}
		if found == false {
			if myKey.Type == jenkinsmodel.Folder {
				addUpdateResults = append(addUpdateResults, jenkinsmodel.JenkinsEdit{
					Keys:     myKey.Keys,
					EditType: jenkinsmodel.AddFolder,
				})
			} else {
				addUpdateResults = append(addUpdateResults, jenkinsmodel.JenkinsEdit{
					Keys:     myKey.Keys,
					Contents: *scripts.GetScriptContentsByKey(myKey),
					EditType: jenkinsmodel.AddJob,
				})
			}
		} else {
			if myKey.Type == jenkinsmodel.Folder {
				continue
			} else {
				addUpdateResults = append(addUpdateResults, jenkinsmodel.JenkinsEdit{
					Keys:     myKey.Keys,
					Contents: *scripts.GetScriptContentsByKey(myKey),
					EditType: jenkinsmodel.UpdateJob,
				})
			}
		}
	}
	return addUpdateResults
}

func getDeleteEdits(
	jenkinsKeys *jenkinsmodel.JenkinsJobKeyList,
	myKeys *jenkinsmodel.JenkinsJobKeyList) jenkinsmodel.JenkinsEditList {
	var deleteResults jenkinsmodel.JenkinsEditList
	for _, myDeleteKey := range *jenkinsKeys {
		result := linq.From(*myKeys).FirstWithT(func(key jenkinsmodel.JenkinsJobKey) bool {
			return stringutil.StringArrayCompare(myDeleteKey.Keys, key.Keys)
		})
		if result == nil && len(myDeleteKey.Keys) > 0 {
			deleteResults = append(deleteResults, jenkinsmodel.JenkinsEdit{
				Keys:     myDeleteKey.Keys,
				EditType: jenkinsmodel.RemoveJobFolder,
			})
		}
	}
	return deleteResults
}
