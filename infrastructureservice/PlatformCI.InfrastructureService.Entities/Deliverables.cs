using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class Deliverables
    {
        public Deliverables()
        {
            DeliverableGroupAssociations = new HashSet<DeliverableGroupAssociations>();
            ResourceAssociations = new HashSet<ResourceAssociations>();
        }

        public Guid DeliverableId { get; set; }
        public string RepositoryName { get; set; }
        public string SolutionName { get; set; }
        public string ProjectName { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public ICollection<DeliverableGroupAssociations> DeliverableGroupAssociations { get; set; }
        public ICollection<ResourceAssociations> ResourceAssociations { get; set; }
    }
}