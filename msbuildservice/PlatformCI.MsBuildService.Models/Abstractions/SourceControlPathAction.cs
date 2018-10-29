using PlatformCI.MsBuildService.Models.Enums;

namespace PlatformCI.MsBuildService.Models.Abstractions
{
    public abstract class SourceControlPathAction
    {
        protected SourceControlPathAction(string nextDirectory, PathActionType action)
        {
            NextDirectory = nextDirectory;
            Action = action;
        }

        public string NextDirectory { get; }

        public PathActionType Action { get; }
    }
}