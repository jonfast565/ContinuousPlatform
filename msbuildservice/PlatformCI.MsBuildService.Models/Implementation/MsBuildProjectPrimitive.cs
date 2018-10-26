using System;
using System.Collections.Generic;

namespace PlatformCI.MsBuildService.Models.Implementation
{
    [Serializable]
    public class MsBuildProjectPrimitive
    {
        public IList<string> TargetFrameworks { get; set; }
        public string DefaultNamespace { get; set; }
        public string AssemblyName { get; set; }

        public ICollection<string> RelativeProjectReferencePaths { get; set; }
            = new List<string>();

        public bool IsNetCoreProject { get; set; }
        public string Name { get; set; }
        public bool Failed { get; set; }
        public Exception Error { get; set; }
    }
}