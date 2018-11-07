using System;
using System.Collections.Generic;

namespace PlatformCI.MsBuildService.Models.Implementation
{
    [Serializable]
    public class MsBuildSolutionPrimitive
    {
        public ICollection<string> Configurations { get; set; }
        public ICollection<string> RelativeProjectPaths { get; set; }

        public string Name { get; set; }
        public bool Failed { get; set; }
        public Exception Error { get; set; }
    }
}