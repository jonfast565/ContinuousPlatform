using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Models.Implementation
{
    [Serializable]
    public class EnvironmentDescriptor
    {
        public string Environment { get; set; }
        public IList<ServerNameType> ServerNames { get; set; }
    }
}