package importmodels

type ScheduledTaskRecord struct {
	Name                      string
	BinaryPath                string
	BinaryExecutableName      string
	BinaryExecutableArguments string
	ScheduleType              string
	RepeatInterval            int64
	RepetitionDuration        int64
	ExecutionTimeLimit        int64
	Priority                  int64
}
