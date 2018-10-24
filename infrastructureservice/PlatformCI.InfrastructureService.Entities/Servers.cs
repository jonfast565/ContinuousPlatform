using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class Servers
    {
        public Servers()
        {
            ServerGroupAssociations = new HashSet<ServerGroupAssociations>();
        }

        public Guid ServerId { get; set; }
        public string ServerName { get; set; }
        public Guid ServerTypeId { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public ServerTypes ServerType { get; set; }
        public ICollection<ServerGroupAssociations> ServerGroupAssociations { get; set; }
    }
}