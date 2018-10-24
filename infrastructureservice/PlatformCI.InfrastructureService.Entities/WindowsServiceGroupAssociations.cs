using System;

namespace PlatformCI.InfrastructureService.Entities
{
    public class WindowsServiceGroupAssociations
    {
        public Guid WindowsServiceId { get; set; }
        public Guid WindowsServiceGroupId { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public WindowsServices WindowsService { get; set; }
        public WindowsServiceGroups WindowsServiceGroup { get; set; }
    }
}