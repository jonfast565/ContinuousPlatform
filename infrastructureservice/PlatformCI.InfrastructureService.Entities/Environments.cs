using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class Environments
    {
        public Environments()
        {
            DeploymentLocationGroups = new HashSet<DeploymentLocationGroups>();
            IisSiteGroups = new HashSet<IisSiteGroups>();
            ScheduledTaskGroups = new HashSet<ScheduledTaskGroups>();
            ServerGroupEnvironmentAssociations = new HashSet<ServerGroupEnvironmentAssociations>();
            WindowsServiceGroups = new HashSet<WindowsServiceGroups>();
        }

        public Guid EnvironmentId { get; set; }
        public string EnvironmentName { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public ICollection<DeploymentLocationGroups> DeploymentLocationGroups { get; set; }
        public ICollection<IisSiteGroups> IisSiteGroups { get; set; }
        public ICollection<ScheduledTaskGroups> ScheduledTaskGroups { get; set; }
        public ICollection<ServerGroupEnvironmentAssociations> ServerGroupEnvironmentAssociations { get; set; }
        public ICollection<WindowsServiceGroups> WindowsServiceGroups { get; set; }
    }
}