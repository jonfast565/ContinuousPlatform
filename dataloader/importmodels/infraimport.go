package importmodels

type InfraImport struct {
	Applications    []AppRecord
	IisApplications []IisApplicationRecord
	IisSites        []IisSiteRecord
	ScheduledTasks  []ScheduledTaskRecord
	WindowsServices []WindowsServiceRecord
	Environments    []EnvironmentRecord
}
