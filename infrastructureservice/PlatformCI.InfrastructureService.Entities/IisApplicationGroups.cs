using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class IisApplicationGroups
    {
        public IisApplicationGroups()
        {
            IisApplicationGroupAssociations = new HashSet<IisApplicationGroupAssociations>();
            IisSiteApplicationGroupAssociations = new HashSet<IisSiteApplicationGroupAssociations>();
        }

        public Guid IisApplicationGroupId { get; set; }
        public string GroupName { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public ICollection<IisApplicationGroupAssociations> IisApplicationGroupAssociations { get; set; }
        public ICollection<IisSiteApplicationGroupAssociations> IisSiteApplicationGroupAssociations { get; set; }
    }
}