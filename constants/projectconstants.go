package constants

var ValidProjectExtensions = []string{
	`^.*\.csproj$`,
	`^.*\.fsproj$`,
	`^.*\.vbproj$`,
}

var ValidSolutionExtensions = []string{
	`^.*\.sln$`,
}

var ValidPublishProfileExtensions = []string{
	`^.*\.pubxml$`,
}

var PublishProfilePathExclusionList = []string{
	":",
}

var PublishProfileSanitizationMap = map[string]string{
	"$(ProjectDir)":  ".",
	"$(SolutionDir)": ".",
}
