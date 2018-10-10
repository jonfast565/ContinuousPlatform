package jenkins

import (
	"../../utilities"
	"net/url"
	"strings"
)

func GetJenkinsJobFolderUrlFragment(baseUrl string, folderUrl string, existenceCheck bool) string {
	folderUrlNoBase := strings.Replace(folderUrl, baseUrl, "", -1)
	pathParser := new(utilities.PathParser)
	pathParser.SetActionSeries(folderUrlNoBase)
	pathParser.RemoveLastNActions(2)
	pathString := pathParser.GetPathString(false) + "/"
	folderJobName := utilities.GetLastPathComponent(folderUrlNoBase)
	var actionFragment string
	if existenceCheck {
		actionFragment = "checkJobName?value="
	} else {
		actionFragment = "createItem?name="
	}
	jenkinsFolderQuery := baseUrl + "/" + pathString + actionFragment + folderJobName
	return url.QueryEscape(jenkinsFolderQuery)
}
