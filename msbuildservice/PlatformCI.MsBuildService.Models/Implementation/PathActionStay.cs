using BuildSystem.Lib.PathParser.Abstractions;
using BuildSystem.Lib.PathParser.Enums;

namespace BuildSystem.Lib.PathParser.Implementation
{
    public class PathActionStay : SourceControlPathAction
    {
        public PathActionStay() : base(null, PathActionType.Stay)
        {
        }
    }
}