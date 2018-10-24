using System;

namespace PlatformCI.InfrastructureService.Entities
{
    public class ScheduledTaskGroupAssociations
    {
        public Guid ScheduledTaskId { get; set; }
        public Guid ScheduledTaskGroupId { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public ScheduledTasks ScheduledTask { get; set; }
        public ScheduledTaskGroups ScheduledTaskGroup { get; set; }
    }
}