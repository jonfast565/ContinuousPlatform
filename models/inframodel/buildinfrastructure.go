package inframodel

type InfrastructureDescriptor struct {
	IisSites               []IisSite
	IisApplications        []IisApplication
	IisApplicationPools    []IisApplicationPool
	WindowsServices        []WindowsService
	ScheduledTasks         []ScheduledTask
	ApplicableEnvironments []string
	RepositoryName         string
	SolutionName           string
	ProjectName            string
}

func (i InfrastructureDescriptor) AnyPhysicalInfrastructure() bool {
	return len(i.IisSites) > 0 ||
		len(i.IisApplications) > 0 ||
		len(i.WindowsServices) > 0 ||
		len(i.ScheduledTasks) > 0
}
