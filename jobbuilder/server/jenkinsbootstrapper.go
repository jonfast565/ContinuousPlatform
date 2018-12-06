package server

import (
	"../../clients/jenkinsclient"
	"../../constants"
	"../../logging"
	"../../models/genmodel"
	"../../models/jenkinsmodel"
	"../../models/jobmodel"
	"../../stringutil"
	"encoding/json"
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

	myKeys := buildKeyListFromScripts(scripts)
	jenkinsKeys := jenkinsMetadata.GetFlattenedKeys()
	edits := buildEditList(&myKeys, jenkinsKeys, scripts)
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
		keySet := script.GetJenkinsKeyList()
		keySet.SanitizeKeyList()
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
	myKeys *jenkinsmodel.JenkinsJobKeyList,
	jenkinsKeys *jenkinsmodel.JenkinsJobKeyList,
	scripts *genmodel.ScriptPackage) jenkinsmodel.JenkinsEditList {
	var results jenkinsmodel.JenkinsEditList
	debugLogKeys(myKeys, jenkinsKeys)
	for _, myKey := range *myKeys {
		found := false
		for _, jenkinsKey := range *jenkinsKeys {
			keyMatch := stringutil.StringArrayCompare(myKey.Keys, jenkinsKey.Keys) && myKey.Type == jenkinsKey.Type
			if keyMatch {
				found = true
				break
			}
		}
		if found == false {
			if myKey.Type == jenkinsmodel.Folder {
				results = append(results, jenkinsmodel.JenkinsEdit{
					Keys:     myKey.Keys,
					EditType: jenkinsmodel.AddFolder,
				})
			} else {
				results = append(results, jenkinsmodel.JenkinsEdit{
					Keys:     myKey.Keys,
					Contents: *scripts.GetScriptContentsByKey(myKey),
					EditType: jenkinsmodel.AddUpdateJob,
				})
			}
		} else {
			if myKey.Type == jenkinsmodel.Folder {
				continue
			} else {
				results = append(results, jenkinsmodel.JenkinsEdit{
					Keys:     myKey.Keys,
					Contents: *scripts.GetScriptContentsByKey(myKey),
					EditType: jenkinsmodel.AddUpdateJob,
				})
			}
		}
	}

	/*
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
	*/

	sort.Sort(results)
	return results
}

func debugLogKeys(myKeys *jenkinsmodel.JenkinsJobKeyList, jenkinsKeys *jenkinsmodel.JenkinsJobKeyList) {
	json1, _ := json.Marshal(jenkinsKeys)
	logging.LogInfo("Jenkins Keys")
	logging.LogInfo(string(json1))
	json2, _ := json.Marshal(myKeys)
	logging.LogInfo("My Keys")
	logging.LogInfo(string(json2))
}
