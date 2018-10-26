using PlatformCI.MsBuildService.Models.Enums;

namespace PlatformCI.MsBuildService.Models.Interfaces
{
    public interface ISourceControlPathAction
    {
        string NextDirectory { get; }
        PathActionType Action { get; }
    }
}