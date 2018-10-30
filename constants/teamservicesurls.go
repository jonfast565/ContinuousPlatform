package constants

import "net/url"

// TODO: Change functionality to use MyURL struct
func GetApiFilesPath(collectionUrl string,
	projectName string,
	repositoryId string,
	branchName string,
	optionalScopePath string) string {
	var scopePath = ""
	if optionalScopePath == "" {
		scopePath = "/"
	} else {
		scopePath = optionalScopePath
	}
	//myUrl := webutil.NewEmptyUrl()
	//myUrl.AppendPathFragments([]string {projectName, "_apis", "git", "repositories", repositoryId, "items"})
	return collectionUrl + "/" + projectName + RepositoryApiSubpath + "/" +
		repositoryId + "/items" + "?" + "scopePath=" + url.QueryEscape(scopePath) +
		"&recursionLevel=Full&includeContentMetadata=true" +
		"&versionType=branch&version=" + url.QueryEscape(branchName) + "&" + GetApiVersionParams()
}

func GetBranchApiPath(collectionUrl string, projectName string, repositoryId string) string {
	return collectionUrl + "/" + projectName + RepositoryApiSubpath + "/" +
		repositoryId + "/refs" + "?" + GetApiVersionParams()
}

func GetRepositoryApiPath(collectionUrl string, projectName string) string {
	return collectionUrl + "/" + projectName + RepositoryApiSubpath + "?" + GetApiVersionParams()
}

func GetApiVersionParams() string {
	return ApiVersionParameter + "=" + ApiVersion
}
