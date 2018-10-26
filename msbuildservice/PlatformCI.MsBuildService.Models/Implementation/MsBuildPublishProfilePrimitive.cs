using System;

namespace PlatformCI.MsBuildService.Models.Implementation
{
    [Serializable]
    public class MsBuildPublishProfilePrimitive
    {
        public string PublishUrl { get; set; }
        public string Name { get; set; }
        public bool Failed { get; set; }
        public Exception Error { get; set; }
    }
}