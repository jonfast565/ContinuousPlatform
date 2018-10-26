using PlatformCI.MsBuildService.Models.Abstractions;
using PlatformCI.MsBuildService.Models.Enums;

namespace PlatformCI.MsBuildService.Models.Implementation
{
    public class PathActionGoAhead : SourceControlPathAction
    {
        public PathActionGoAhead(string nextDirectory) : base(nextDirectory, PathActionType.GoAhead)
        {
        }
    }
}