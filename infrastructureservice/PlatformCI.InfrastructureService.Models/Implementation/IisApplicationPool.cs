using System;

namespace PlatformCI.InfrastructureService.Models.Implementation
{
    [Serializable]
    public class IisApplicationPool
    {
        public string AppPoolName { get; set; }
        public string AppPoolType { get; set; }
        public string AppPoolFrameworkVersion { get; set; }
        public Guid AppPoolGuid { get; set; }

        public string RealAppPoolFrameworkVersion =>
            AppPoolFrameworkVersion == "No Managed Code" ? string.Empty : AppPoolFrameworkVersion;
    }
}