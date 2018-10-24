using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class ResourceAmalgamations
    {
        public ResourceAmalgamations()
        {
            ResourceAssociations = new HashSet<ResourceAssociations>();
        }

        public Guid ResourceAmalgamationId { get; set; }
        public Guid BusinessLineId { get; set; }
        public string Name { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public BusinessLines BusinessLine { get; set; }
        public ICollection<ResourceAssociations> ResourceAssociations { get; set; }
    }
}