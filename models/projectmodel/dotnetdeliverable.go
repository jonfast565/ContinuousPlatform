package projectmodel

type DotNetDeliverable struct {
	Repository string
	Branch     string
	Solution   *MsBuildSolution
	Project    *MsBuildProject
}
