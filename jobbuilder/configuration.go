package main

type Configuration struct {
	Port              int
	RunJobsOnStartup  bool
	CyclicalRuns      bool
	CycleRateLimiting bool
	CycleRateLimit    int
	DebugBasePath     string
}
