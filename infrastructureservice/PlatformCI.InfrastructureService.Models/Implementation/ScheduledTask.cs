using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Models.Implementation
{
    [Serializable]
    public class ScheduledTask
    {
        public string TaskName { get; set; }
        public string BinaryPath { get; set; }
        public string BinaryExecutableName { get; set; }
        public string BinaryExecutableArguments { get; set; }
        public string ScheduleType { get; set; }
        public TimeSpan? RepeatInterval { get; set; }
        public TimeSpan? RepetitionDuration { get; set; }
        public TimeSpan? ExecutionTimeLimit { get; set; }
        public int? Priority { get; set; }
        public Guid TaskGuid { get; set; }
        public IList<string> Environments { get; set; }
    }
}