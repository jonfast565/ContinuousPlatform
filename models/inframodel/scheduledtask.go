package inframodel

type ScheduledTask struct {
	TaskName                  string
	BinaryPath                string
	BinaryExecutableName      string
	BinaryExecutableArguments string
	ScheduleType              string
	RepeatInterval            int64 // TODO: Deal with these appropriately
	RepetitionDuration        int64
	ExecutionTimeLimit        int64
	Priority                  int64
	TaskGuid                  string
}
