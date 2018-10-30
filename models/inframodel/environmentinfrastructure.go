package inframodel

type EnvironmentInfrastructure struct {
	Environment            string
	IisApplicationPools    []IisApplicationPool
	ScheduledTasks         []ScheduledTask
	WindowsServices        []WindowsService
	SiteApplicationMembers []IisSiteApplicationMemberPool
	Servers                []ServerNameType
}
