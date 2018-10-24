package inframodel

type InfrastructureMetadata struct {
	Infrastructure []InfrastructureDescriptor
	Environments   []EnvironmentDescriptor
}

type EnvironmentInfrastructure struct {
	Environment            string
	IisApplicationPools    []IisApplicationPool
	ScheduledTasks         []ScheduledTask
	WindowsServices        []WindowsService
	SiteApplicationMembers []IisSiteApplicationMemberPool
	Servers                []ServerNameType
}

type FlattenedBuildInfrastructure struct {
	ServerName          string
	ServerGroup         string
	DeploymentLocations []string
	AppPoolNames        []string
	ServiceNames        []string
	TaskNames           []string
}

type EnvironmentDescriptor struct {
	Environment string
	ServerNames []ServerNameType
}

type ServerNameType struct {
	ServerName string
	ServerType string
}

type IisApplication struct {
	Sites           []IisSite
	ApplicationName string
	PhysicalPath    string
	AppPool         IisApplicationPool
	ApplicationGuid string
}

type IisApplicationPool struct {
	AppPoolName             string
	AppPoolType             string
	AppPoolFrameworkVersion string
	AppPoolGuid             string
}

type IisSite struct {
	SiteName     string
	PhysicalPath string
	AppPool      IisApplicationPool
	SiteGuid     string
	Environments []string
}

type IisSiteApplicationMemberPool struct {
	ParentSite        IisSite
	ChildApplications []IisApplication
	// TODO: Implement InitPools
}

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

type ScheduledTask struct {
	TaskName                  string
	BinaryPath                string
	BinaryExecutableName      string
	BinaryExecutableArguments string
	ScheduleType              string
	RepeatInterval            int64 // TODO: Deal with these appropriately
	RepetitionDuration        int64
	ExecutionTimeLimit        int64
	Priority                  int
	TaskGuid                  string
	Environments              []string
}

type WindowsService struct {
	ServiceName               string
	BinaryPath                string
	BinaryExecutableName      string
	BinaryExecutableArguments string
	ServiceGuid               string
	Environments              []string
}

func (i InfrastructureDescriptor) AnyPhysicalInfrastructure() bool {
	return len(i.IisSites) > 0 ||
		len(i.IisApplications) > 0 ||
		len(i.WindowsServices) > 0 ||
		len(i.ScheduledTasks) > 0
}
