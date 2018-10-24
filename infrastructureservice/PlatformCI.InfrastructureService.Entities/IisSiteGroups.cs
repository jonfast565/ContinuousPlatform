using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class IisSiteGroups
    {
        public IisSiteGroups()
        {
            IisSiteGroupAssociations = new HashSet<IisSiteGroupAssociations>();
        }

        public Guid IisSiteGroupId { get; set; }
        public string GroupName { get; set; }
        public Guid EnvironmentId { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public Environments Environment { get; set; }
        public ICollection<IisSiteGroupAssociations> IisSiteGroupAssociations { get; set; }
    }
}