using System;

namespace BuildSystem.Lib.Models.Deliverable.Implementation
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