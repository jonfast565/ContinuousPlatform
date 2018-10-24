using System;

namespace PlatformCI.InfrastructureService.Entities
{
    public class ServerGroupAssociations
    {
        public Guid ServerId { get; set; }
        public Guid ServerGroupId { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public Servers Server { get; set; }
        public ServerGroups ServerGroup { get; set; }
    }
}