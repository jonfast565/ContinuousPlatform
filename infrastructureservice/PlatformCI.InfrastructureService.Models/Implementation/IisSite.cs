using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Models.Implementation
{
    [Serializable]
    public class IisSite
    {
        public string SiteName { get; set; }
        public string PhysicalPath { get; set; }
        public IisApplicationPool AppPool { get; set; }
        public Guid SiteGuid { get; set; }
        public IList<string> Environments { get; set; }
    }
}