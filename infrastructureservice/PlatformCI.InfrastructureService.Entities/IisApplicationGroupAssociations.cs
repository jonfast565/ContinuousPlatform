using System;

namespace PlatformCI.InfrastructureService.Entities
{
    public class IisApplicationGroupAssociations
    {
        public Guid IisApplicationGroupId { get; set; }
        public Guid IisApplicationId { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public IisApplications IisApplication { get; set; }
        public IisApplicationGroups IisApplicationGroup { get; set; }
    }
}