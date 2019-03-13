package constants

import "net/url"

// Gets the file paths for the team services API
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
	// TODO: Change functionality to use MyURL struct
	// myUrl := webutil.NewEmptyUrl()
	// myUrl.AppendPathFragments([]string {projectName, "_apis", "git", "repositories", repositoryId, "items"})
	return collectionUrl + "/" + projectName + RepositoryApiSubpath + "/" +
		repositoryId + "/items" + "?" + "scopePath=" + url.QueryEscape(scopePath) +
		"&recursionLevel=Full&includeContentMetadata=true" +
		"&versionType=branch&version=" + url.QueryEscape(branchName) + "&" + GetApiVersionParams()
}

// Get the branch path for a repository from the Team Services API
func GetBranchApiPath(collectionUrl string, projectName string, repositoryId string) string {
	// TODO: Change functionality to use MyURL struct
	return collectionUrl + "/" + projectName + RepositoryApiSubpath + "/" +
		repositoryId + "/refs" + "?" + GetApiVersionParams()
}

// Get the repository path for a
func GetRepositoryApiPath(collectionUrl string, projectName string) string {
	// TODO: Change functionality to use MyURL struct
	return collectionUrl + "/" + projectName + RepositoryApiSubpath + "?" + GetApiVersionParams()
}

// Get the API version parameter string
// TODO: Why is this method so small? Replace with MyURL struct in parent
func GetApiVersionParams() string {
	// TODO: Change functionality to use MyURL struct
	return ApiVersionParameter + "=" + ApiVersion
}
