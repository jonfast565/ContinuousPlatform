using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class IisApplications
    {
        public IisApplications()
        {
            IisApplicationGroupAssociations = new HashSet<IisApplicationGroupAssociations>();
            ResourceAssociations = new HashSet<ResourceAssociations>();
        }

        public Guid IisApplicationId { get; set; }
        public Guid IisApplicationPoolId { get; set; }
        public string ApplicationName { get; set; }
        public string PhysicalPath { get; set; }
        public string ApplicationInternalAliasName { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public IisApplicationPools IisApplicationPool { get; set; }
        public ICollection<IisApplicationGroupAssociations> IisApplicationGroupAssociations { get; set; }
        public ICollection<ResourceAssociations> ResourceAssociations { get; set; }
    }
}