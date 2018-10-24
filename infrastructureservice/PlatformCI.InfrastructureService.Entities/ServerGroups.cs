using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class ServerGroups
    {
        public ServerGroups()
        {
            ServerGroupAssociations = new HashSet<ServerGroupAssociations>();
            ServerGroupEnvironmentAssociations = new HashSet<ServerGroupEnvironmentAssociations>();
        }

        public Guid ServerGroupId { get; set; }
        public string GroupName { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public ICollection<ServerGroupAssociations> ServerGroupAssociations { get; set; }
        public ICollection<ServerGroupEnvironmentAssociations> ServerGroupEnvironmentAssociations { get; set; }
    }
}