using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Models.Implementation
{
    [Serializable]
    public class IisApplication
    {
        public IList<IisSite> Sites { get; set; }
        public string ApplicationName { get; set; }
        public string PhysicalPath { get; set; }
        public IisApplicationPool AppPool { get; set; }
        public Guid ApplicationGuid { get; set; }
    }
}