using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class BusinessLines
    {
        public BusinessLines()
        {
            ResourceAmalgamations = new HashSet<ResourceAmalgamations>();
        }

        public Guid BusinessLineId { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public ICollection<ResourceAmalgamations> ResourceAmalgamations { get; set; }
    }
}