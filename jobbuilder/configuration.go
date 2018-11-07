package main

type Configuration struct {
	Port                    int
	RunJobsOnStartup        bool
	CyclicalRuns            bool
	ChangeRateLimiting      bool
	ChangeRateLimit         int
	BetweenJobWait          int
	ProceedDespiteNoChanges bool
}
