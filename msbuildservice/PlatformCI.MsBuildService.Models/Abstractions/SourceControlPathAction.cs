using PlatformCI.MsBuildService.Models.Enums;
using PlatformCI.MsBuildService.Models.Interfaces;

namespace PlatformCI.MsBuildService.Models.Abstractions
{
    public abstract class SourceControlPathAction : ISourceControlPathAction
    {
        public SourceControlPathAction(string nextDirectory, PathActionType action)
        {
            NextDirectory = nextDirectory;
            Action = action;
        }

        public string NextDirectory { get; }

        public PathActionType Action { get; }
    }
}