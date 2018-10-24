using System;

namespace PlatformCI.InfrastructureService.Models.Implementation
{
    [Serializable]
    public class ServerNameType
    {
        public string ServerName { get; set; }
        public string ServerType { get; set; }
    }
}