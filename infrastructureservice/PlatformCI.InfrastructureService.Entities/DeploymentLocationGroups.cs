using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class DeploymentLocationGroups
    {
        public DeploymentLocationGroups()
        {
            DeploymentLocationGroupAssociations = new HashSet<DeploymentLocationGroupAssociations>();
        }

        public Guid DeploymentLocationGroupId { get; set; }
        public string GroupName { get; set; }
        public Guid EnvironmentId { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public Environments Environment { get; set; }
        public ICollection<DeploymentLocationGroupAssociations> DeploymentLocationGroupAssociations { get; set; }
    }
}