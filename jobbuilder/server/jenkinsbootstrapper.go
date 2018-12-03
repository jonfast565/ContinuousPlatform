package server

import (
	"../../clients/jenkinsclient"
	"../../constants"
	"../../logging"
	"../../models/genmodel"
	"../../models/jenkinsmodel"
	"../../models/jobmodel"
	"../../stringutil"
	"github.com/ahmetb/go-linq"
	"sort"
)

func DeployJenkinsJobs(details *jobmodel.JobDetails) {
	defer func() {
		if r := recover(); r != nil {
			details.SetJobErrored()
			logging.LogPanicRecover(r)
		}
	}()

	scripts, err := GetScriptCache()
	if err != nil {
		panic(err)
	}

	jenkinsClient := jenkinsclient.NewJenkinsClient()
	jenkinsMetadata, err := jenkinsClient.GetJenkinsMetadata()
	if err != nil {
		panic(err)
	}

	myMetadataKeys := buildKeyListFromScripts(scripts)
	jenkinsInstanceMetadataKeys := jenkinsMetadata.GetFlattenedKeys()
	edits := buildEditList(&myMetadataKeys, jenkinsInstanceMetadataKeys)
	persistEditList(edits, jenkinsClient)
}

func persistEditList(edits jenkinsmodel.JenkinsEditList, jenkinsClient jenkinsclient.JenkinsClient) {
	for _, edit := range edits {
		jobRequest := edit.GetJobRequest()
		jobRequest.SanitizeSegments()
		switch edit.EditType {
		case jenkinsmodel.AddUpdateJob:
			logging.LogInfo("Create/Update Job: " + jobRequest.GetJobFragmentUrl())
			_, err := jenkinsClient.CreateUpdateJob(jobRequest)
			if err != nil {
				panic(err)
			}
			break
		case jenkinsmodel.AddFolder:
			logging.LogInfo("Create Folder Job: " + jobRequest.GetJobFragmentUrl())
			_, err := jenkinsClient.CreateFolder(jobRequest)
			if err != nil {
				panic(err)
			}
			break
		case jenkinsmodel.RemoveJobFolder:
			logging.LogInfo("Remove Job/Folder: " + jobRequest.GetJobFragmentUrl())
			_, err := jenkinsClient.DeleteJobOrFolder(jobRequest)
			if err != nil {
				panic(err)
			}
			break
		default:
			panic("not an option")
		}
	}
}

func buildKeyListFromScripts(scripts *genmodel.ScriptPackage) jenkinsmodel.JenkinsJobKeyList {
	myMetadataKeys := make(jenkinsmodel.JenkinsJobKeyList, 0)
	for _, script := range scripts.Scripts {
		isJenkinsScript := stringutil.StringArrayContains(script.ToolScope, constants.JenkinsToolName)
		if !isJenkinsScript {
			continue
		}
		keySet := script.GetJenkinsKeySet()
		for _, key := range keySet {
			if !myMetadataKeys.KeyAlreadyExists(key.Keys) {
				myMetadataKeys = append(myMetadataKeys, key)
			}
		}
	}
	sort.Sort(myMetadataKeys)
	return myMetadataKeys
}

func buildEditList(
	l1 *jenkinsmodel.JenkinsJobKeyList,
	l2 *jenkinsmodel.JenkinsJobKeyList) jenkinsmodel.JenkinsEditList {
	var results jenkinsmodel.JenkinsEditList
	for _, k1 := range *l1 {
		result := linq.From(*l2).FirstWithT(func(key jenkinsmodel.JenkinsJobKey) bool {
			return stringutil.StringArrayCompare(k1.Keys, key.Keys) && k1.Type == key.Type
		})

		if result != nil {
			// update
			resultKey := result.(jenkinsmodel.JenkinsJobKey)
			if resultKey.Type == jenkinsmodel.Folder {
				continue
			}
			results = append(results, jenkinsmodel.JenkinsEdit{
				Keys:     resultKey.Keys,
				Contents: "",
				EditType: jenkinsmodel.AddUpdateJob,
			})
		} else {
			// insert
			if k1.Type == jenkinsmodel.Folder {
				results = append(results, jenkinsmodel.JenkinsEdit{
					Keys:     k1.Keys,
					EditType: jenkinsmodel.AddFolder,
				})
			} else {
				results = append(results, jenkinsmodel.JenkinsEdit{
					Keys:     k1.Keys,
					Contents: "",
					EditType: jenkinsmodel.AddUpdateJob,
				})
			}
		}
	}

	for _, k2 := range *l2 {
		result := linq.From(*l2).FirstWithT(func(key jenkinsmodel.JenkinsJobKey) bool {
			return stringutil.StringArrayCompare(k2.Keys, key.Keys)
		})
		if result != nil {
			// delete
			resultKey := result.(jenkinsmodel.JenkinsJobKey)
			results = append(results, jenkinsmodel.JenkinsEdit{
				Keys:     resultKey.Keys,
				EditType: jenkinsmodel.RemoveJobFolder,
			})
		}
	}

	sort.Sort(results)
	return results
}
