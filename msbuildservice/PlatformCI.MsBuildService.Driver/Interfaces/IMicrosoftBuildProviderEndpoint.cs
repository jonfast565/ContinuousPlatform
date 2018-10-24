using BuildSystem.Lib.Interfaces.Generic.Implementation;
using BuildSystem.Lib.Models.Deliverable.Implementation;

namespace BuildSystem.Lib.MicrosoftBuildProvider.Interfaces
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