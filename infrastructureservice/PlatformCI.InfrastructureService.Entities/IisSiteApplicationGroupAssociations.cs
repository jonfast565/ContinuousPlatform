using System;

namespace PlatformCI.InfrastructureService.Entities
{
    public class IisSiteApplicationGroupAssociations
    {
        public Guid IisSiteId { get; set; }
        public Guid IisApplicationGroupId { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public IisApplicationGroups IisApplicationGroup { get; set; }
        public IisSites IisSite { get; set; }
    }
}