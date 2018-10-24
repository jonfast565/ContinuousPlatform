using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Entities
{
    public class IisApplicationPools
    {
        public IisApplicationPools()
        {
            IisApplications = new HashSet<IisApplications>();
            IisSites = new HashSet<IisSites>();
        }

        public Guid IisApplicationPoolId { get; set; }
        public string PoolName { get; set; }
        public string PoolType { get; set; }
        public string PoolFrameworkVersion { get; set; }
        public string CreatedBy { get; set; }
        public DateTime CreatedDateTime { get; set; }
        public string LastModifiedBy { get; set; }
        public DateTime LastModifiedDateTime { get; set; }

        public ICollection<IisApplications> IisApplications { get; set; }
        public ICollection<IisSites> IisSites { get; set; }
    }
}