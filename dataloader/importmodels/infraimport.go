package importmodels

type InfraImport struct {
	Applications        []AppRecord
	IisApplications     []IisApplicationRecord
	IisApplicationPools []IisApplicationPoolRecord
	IisSites            []IisSiteRecord
	ScheduledTasks      []ScheduledTaskRecord
	WindowsServices     []WindowsServiceRecord
	Environments        []EnvironmentRecord
	Servers             []ServerRecord
}
