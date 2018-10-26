using PlatformCI.MsBuildService.Models.Abstractions;
using PlatformCI.MsBuildService.Models.Enums;

namespace PlatformCI.MsBuildService.Models.Implementation
{
    public class PathActionGoBack : SourceControlPathAction
    {
        public PathActionGoBack() : base(null, PathActionType.GoBack)
        {
        }
    }
}