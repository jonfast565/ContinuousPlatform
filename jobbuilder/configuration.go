package main

// configuration of jobs
type JobConfiguration struct {
	Port              int
	RunJobsOnStartup  bool
	CyclicalRuns      bool
	CycleRateLimiting bool
	CycleRateLimit    int
	DebugBasePath     string
}
