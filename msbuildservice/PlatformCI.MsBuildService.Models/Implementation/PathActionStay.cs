using PlatformCI.MsBuildService.Models.Abstractions;
using PlatformCI.MsBuildService.Models.Enums;

namespace PlatformCI.MsBuildService.Models.Implementation
{
    public class PathActionStay : SourceControlPathAction
    {
        public PathActionStay() : base(null, PathActionType.Stay)
        {
        }
    }
}