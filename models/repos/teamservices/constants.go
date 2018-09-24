package teamservices

import (
	"net/url"
)

const (
	ApiVersionParameter  string = "api-version"
	ApiVersion           string = "1.0"
	RepositoryApiSubpath string = "/_apis/git/repositories"
	RefsHeadsConstants   string = "refs/heads/"
	BlobConstant         string = "blob"
)

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
