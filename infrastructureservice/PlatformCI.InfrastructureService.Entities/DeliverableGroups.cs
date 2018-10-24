using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class DeliverableGroups
    {
        public DeliverableGroups()
        {
            DeliverableGroupAssociations = new HashSet<DeliverableGroupAssociations>();
        }

        public Guid DeliverableGroupId { get; set; }
        public string GroupName { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public ICollection<DeliverableGroupAssociations> DeliverableGroupAssociations { get; set; }
    }
}