using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class DeploymentLocations
    {
        public DeploymentLocations()
        {
            DeploymentLocationGroupAssociations = new HashSet<DeploymentLocationGroupAssociations>();
            ResourceAssociations = new HashSet<ResourceAssociations>();
        }

        public Guid DeploymentLocationId { get; set; }
        public string FriendlyName { get; set; }
        public string PhysicalPath { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public ICollection<DeploymentLocationGroupAssociations> DeploymentLocationGroupAssociations { get; set; }
        public ICollection<ResourceAssociations> ResourceAssociations { get; set; }
    }
}