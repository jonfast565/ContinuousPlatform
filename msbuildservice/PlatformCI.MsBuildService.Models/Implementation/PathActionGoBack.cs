using BuildSystem.Lib.PathParser.Abstractions;
using BuildSystem.Lib.PathParser.Enums;

namespace BuildSystem.Lib.PathParser.Implementation
{
    public class PathActionGoBack : SourceControlPathAction
    {
        public PathActionGoBack() : base(null, PathActionType.GoBack)
        {
        }
    }
}