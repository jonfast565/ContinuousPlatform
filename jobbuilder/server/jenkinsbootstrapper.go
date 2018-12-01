package server

import (
	"../../clients/jenkinsclient"
	"../../constants"
	"../../logging"
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
	jenkinsKeys := jenkinsMetadata.GetFlattenedKeys()
	if err != nil {
		panic(err)
	}

	myMetadataKeys := make(jenkinsmodel.JenkinsJobKeyList, 0)
	for _, script := range scripts.Scripts {
		isJenkinsScript := stringutil.StringArrayContains(script.ToolScope, constants.JenkinsToolName)
		if !isJenkinsScript {
			continue
		}
		var tempKeys []string
		for i, key := range script.KeyElements {
			tempKeys = append(tempKeys, key)
			currentKeys := make([]string, 0)
			currentKeys = append(currentKeys, tempKeys...)
			if i == len(script.KeyElements)-1 {
				if !myMetadataKeys.KeyAlreadyExists(currentKeys) {
					myMetadataKeys = append(myMetadataKeys, jenkinsmodel.JenkinsJobKey{
						Keys: script.KeyElements,
						Type: string(jenkinsmodel.PipelineJob),
					})
				}
			} else {
				if !jenkinsKeys.PartialKeyAlreadyExists(currentKeys) &&
					!myMetadataKeys.KeyAlreadyExists(currentKeys) {
					myMetadataKeys = append(myMetadataKeys, jenkinsmodel.JenkinsJobKey{
						Keys: currentKeys,
						Type: string(jenkinsmodel.Folder),
					})
				}
			}
		}
	}

	sort.Sort(myMetadataKeys)
	jenkinsInstanceMetadataKeys := jenkinsMetadata.GetFlattenedKeys()
	edits := buildEditList(&myMetadataKeys, jenkinsInstanceMetadataKeys)
	sort.Sort(edits)

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

func buildEditList(
	l1 *jenkinsmodel.JenkinsJobKeyList,
	l2 *jenkinsmodel.JenkinsJobKeyList) jenkinsmodel.JenkinsEditList {
	var results jenkinsmodel.JenkinsEditList
	for _, k1 := range *l1 {
		result := linq.From(*l2).FirstWithT(func(key jenkinsmodel.JenkinsJobKey) bool {
			return stringutil.StringArrayCompare(k1.Keys, key.Keys)
		})

		if result != nil {
			// update
			resultKey := result.(jenkinsmodel.JenkinsJobKey)
			if resultKey.Type == string(jenkinsmodel.Folder) {
				continue
			}
			results = append(results, jenkinsmodel.JenkinsEdit{
				Keys:     resultKey.Keys,
				Contents: "",
				EditType: jenkinsmodel.AddUpdateJob,
			})
		} else {
			// insert
			if k1.Type == string(jenkinsmodel.Folder) {
				results = append(results, jenkinsmodel.JenkinsEdit{
					Keys:     k1.Keys,
					Contents: "",
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
				Contents: "",
				EditType: jenkinsmodel.RemoveJobFolder,
			})
		}
	}

	return results
}
