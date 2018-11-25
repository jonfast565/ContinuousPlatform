package importmodels

type ScheduledTaskImport struct {
	Names                     []string
	BinaryPath                string
	BinaryExecutableName      string
	BinaryExecutableArguments string
	ScheduleType              string
	RepeatInterval            int64
	RepetitionDuration        int64
	ExecutionTimeLimit        int64
	Priority                  int64
	LoadBalanced              bool
	Environments              []EnvironmentImportPart
}
