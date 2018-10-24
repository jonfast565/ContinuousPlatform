using System;

namespace PlatformCI.InfrastructureService.Entities
{
    public class DeliverableGroupAssociations
    {
        public Guid DeliverableId { get; set; }
        public Guid DeliverableGroupId { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public Deliverables Deliverable { get; set; }
        public DeliverableGroups DeliverableGroup { get; set; }
    }
}