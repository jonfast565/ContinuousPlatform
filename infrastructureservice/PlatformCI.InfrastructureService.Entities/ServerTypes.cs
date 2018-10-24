using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class ServerTypes
    {
        public ServerTypes()
        {
            Servers = new HashSet<Servers>();
        }

        public Guid ServerTypeId { get; set; }
        public string ServerTypeName { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public ICollection<Servers> Servers { get; set; }
    }
}