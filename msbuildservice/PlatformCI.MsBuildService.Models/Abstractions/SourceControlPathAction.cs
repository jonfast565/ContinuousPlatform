using BuildSystem.Lib.PathParser.Enums;
using BuildSystem.Lib.PathParser.Interfaces;

namespace BuildSystem.Lib.PathParser.Abstractions
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