using System;
using System.Collections.Generic;

namespace PlatformCI.InfrastructureService.Models.Implementation
{
    [Serializable]
    public class WindowsService
    {
        public string ServiceName { get; set; }
        public string BinaryPath { get; set; }
        public string BinaryExecutableName { get; set; }
        public string BinaryExecutableArguments { get; set; }
        public Guid ServiceGuid { get; set; }
        public IList<string> Environments { get; set; }
    }
}