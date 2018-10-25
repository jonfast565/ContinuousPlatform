using BuildSystem.Lib.PathParser.Abstractions;
using BuildSystem.Lib.PathParser.Enums;

namespace BuildSystem.Lib.PathParser.Implementation
{
    public class PathActionGoAhead : SourceControlPathAction
    {
        public PathActionGoAhead(string nextDirectory) : base(nextDirectory, PathActionType.GoAhead)
        {
        }
    }
}