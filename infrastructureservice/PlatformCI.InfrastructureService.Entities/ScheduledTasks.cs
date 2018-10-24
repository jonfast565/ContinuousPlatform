using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class ScheduledTasks
    {
        public ScheduledTasks()
        {
            ResourceAssociations = new HashSet<ResourceAssociations>();
            ScheduledTaskGroupAssociations = new HashSet<ScheduledTaskGroupAssociations>();
        }

        public Guid ScheduledTaskId { get; set; }
        public string TaskName { get; set; }
        public string BinaryPath { get; set; }
        public string BinaryExecutableName { get; set; }
        public string BinaryExecutableArguments { get; set; }
        public string ScheduleType { get; set; }
        public long? RepeatInterval { get; set; }
        public long? RepetitionDuration { get; set; }
        public long? ExecutionTimeLimit { get; set; }
        public int? Priority { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public ICollection<ResourceAssociations> ResourceAssociations { get; set; }
        public ICollection<ScheduledTaskGroupAssociations> ScheduledTaskGroupAssociations { get; set; }
    }
}