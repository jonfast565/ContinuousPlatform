using System;

namespace PlatformCI.InfrastructureService.Entities
{
    public class IisSiteGroupAssociations
    {
        public Guid IisSiteGroupId { get; set; }
        public Guid IisSiteId { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public IisSites IisSite { get; set; }
        public IisSiteGroups IisSiteGroup { get; set; }
    }
}