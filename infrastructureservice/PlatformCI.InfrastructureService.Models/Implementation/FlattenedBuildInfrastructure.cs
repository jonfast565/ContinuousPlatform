using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Models.Implementation
{
    [Serializable]
    public class FlattenedBuildInfrastructure
    {
        public string ServerName { get; set; }
        public string ServerGroup { get; set; }
        public IList<string> DeploymentLocations { get; set; }
        public IList<string> AppPoolNames { get; set; }
        public IList<string> ServiceNames { get; set; }
        public IList<string> TaskNames { get; set; }
    }
}