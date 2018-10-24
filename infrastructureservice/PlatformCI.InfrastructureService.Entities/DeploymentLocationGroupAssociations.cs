using System;

namespace PlatformCI.InfrastructureService.Entities
{
    public class DeploymentLocationGroupAssociations
    {
        public Guid DeploymentLocationId { get; set; }
        public Guid DeploymentLocationGroupId { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public DeploymentLocations DeploymentLocation { get; set; }
        public DeploymentLocationGroups DeploymentLocationGroup { get; set; }
    }
}