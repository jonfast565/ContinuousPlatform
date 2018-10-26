using PlatformCI.MsBuildService.Models.Implementation;

namespace PlatformCI.MsBuildService.Driver.Interfaces
{
    public interface IMicrosoftBuildProviderEndpoint
    {
        MsBuildSolutionPrimitive GetSolutionFromFileBytes(FilePayload localPath);

        MsBuildProjectPrimitive GetProjectFromFileBytes(FilePayload localPath);

        MsBuildPublishProfilePrimitive GetPublishProfileFromFileBytes(FilePayload localPath);

        MsBuildProjectPrimitive GetProjectFromLocalPath(string localPath, string originalProjectName = null);

        MsBuildPublishProfilePrimitive GetPublishProfileFromLocalPath(string localPath,
            string originalPublishProfileName = null);

        MsBuildSolutionPrimitive GetSolutionFromLocalPath(string localPath, string originalSolutionName = null);
    }
}