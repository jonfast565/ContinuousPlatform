using System;
using PlatformCI.MsBuildService.Models.Enums;

namespace PlatformCI.MsBuildService.Models.Implementation
{
    [Serializable]
    public class FilesystemMetadata
    {
        public string Path { get; set; }
        public FilesystemObjectType Type { get; set; }
        public string OptionalChangeId { get; set; }
    }
}