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
	metadata, err := jenkinsClient.GetJenkinsMetadata()
	if err != nil {
		panic(err)
	}

	var myMetadataKeys jenkinsmodel.JenkinsJobKeyList
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
				if !myMetadataKeys.KeyAlreadyExists(currentKeys) {
					myMetadataKeys = append(myMetadataKeys, jenkinsmodel.JenkinsJobKey{
						Keys: currentKeys,
						Type: string(jenkinsmodel.Folder),
					})
				}
			}
		}
	}

	sort.Sort(myMetadataKeys)
	jenkinsInstanceMetadataKeys := metadata.GetFlattenedKeys()
	edits := buildEditList(&myMetadataKeys, jenkinsInstanceMetadataKeys)
}

func buildEditList(
	l1 *jenkinsmodel.JenkinsJobKeyList,
	l2 *jenkinsmodel.JenkinsJobKeyList) []jenkinsmodel.JenkinsEdit {
	var results []jenkinsmodel.JenkinsEdit
	for _, k1 := range *l1 {
		result := linq.From(l2).FirstWithT(func(key jenkinsmodel.JenkinsJobKey) bool {
			return stringutil.StringArrayCompare(k1.Keys, key.Keys)
		})

		if result != nil {
			// update
			results = append(results, jenkinsmodel.JenkinsEdit{
				Name:      "",
				Url:       "",
				EditType:  0,
				JobRecord: JenkinsJobRecord{},
			})
		} else {
			// insert
			results = append(results, jenkinsmodel.JenkinsEdit{
				Name:      "",
				Url:       "",
				EditType:  0,
				JobRecord: JenkinsJobRecord{},
			})
		}
	}

	for _, k2 := range *l2 {
		result := linq.From(l2).FirstWithT(func(key jenkinsmodel.JenkinsJobKey) bool {
			return stringutil.StringArrayCompare(k2.Keys, key.Keys)
		})
		if result != nil {
			// delete
			results = append(results, jenkinsmodel.JenkinsEdit{
				Name:      "",
				Url:       "",
				EditType:  0,
				JobRecord: JenkinsJobRecord{},
			})
		}
	}

	return nil
}
