using System;
using BuildSystem.Lib.FilesystemProvider.Enums;

namespace BuildSystem.Lib.FilesystemProvider.Implementation
{
    [Serializable]
    public class FilesystemMetadata
    {
        public string Path { get; set; }
        public FilesystemObjectType Type { get; set; }
        public string OptionalChangeId { get; set; }
    }
}