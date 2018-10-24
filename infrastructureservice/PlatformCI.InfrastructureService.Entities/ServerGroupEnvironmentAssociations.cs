using System;

namespace PlatformCI.InfrastructureService.Entities
{
    public class ServerGroupEnvironmentAssociations
    {
        public Guid EnvironmentId { get; set; }
        public Guid ServerGroupId { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public Environments Environment { get; set; }
        public ServerGroups ServerGroup { get; set; }
    }
}