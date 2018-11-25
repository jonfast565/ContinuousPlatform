package importmodels

type InfrastructureImport struct {
	Applications    []AppImport
	IisApplications []IisApplicationImport
	IisSites        []IisSiteImport
	ScheduledTasks  []ScheduledTaskImport
	WindowsServices []WindowsServiceImport
	Environments    []EnvironmentImport
}
