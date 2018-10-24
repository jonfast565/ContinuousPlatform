using System;

namespace PlatformCI.InfrastructureService.Entities
{
    public class Jobs
    {
        public Guid JobId { get; set; }
        public string JobName { get; set; }
        public string MachineName { get; set; }
        public Guid JobStateId { get; set; }
        public bool JobTrigger { get; set; }
        public int JobOrder { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public JobStates JobState { get; set; }
    }
}