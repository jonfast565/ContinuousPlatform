using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class ScheduledTaskGroups
    {
        public ScheduledTaskGroups()
        {
            ScheduledTaskGroupAssociations = new HashSet<ScheduledTaskGroupAssociations>();
        }

        public Guid ScheduledTaskGroupId { get; set; }
        public string TaskGroupName { get; set; }
        public Guid EnvironmentId { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public Environments Environment { get; set; }
        public ICollection<ScheduledTaskGroupAssociations> ScheduledTaskGroupAssociations { get; set; }
    }
}