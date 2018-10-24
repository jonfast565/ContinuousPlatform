using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class IisSites
    {
        public IisSites()
        {
            IisSiteApplicationGroupAssociations = new HashSet<IisSiteApplicationGroupAssociations>();
            IisSiteBindings = new HashSet<IisSiteBindings>();
            IisSiteGroupAssociations = new HashSet<IisSiteGroupAssociations>();
            ResourceAssociations = new HashSet<ResourceAssociations>();
        }

        public Guid IisSiteId { get; set; }
        public Guid IisSiteApplicationPoolId { get; set; }
        public string SiteName { get; set; }
        public string PhysicalPath { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public IisApplicationPools IisSiteApplicationPool { get; set; }
        public ICollection<IisSiteApplicationGroupAssociations> IisSiteApplicationGroupAssociations { get; set; }
        public ICollection<IisSiteBindings> IisSiteBindings { get; set; }
        public ICollection<IisSiteGroupAssociations> IisSiteGroupAssociations { get; set; }
        public ICollection<ResourceAssociations> ResourceAssociations { get; set; }
    }
}