using BuildSystem.Lib.Models.Deliverable.Implementation;

namespace BuildSystem.Lib.MicrosoftBuildProvider.Interfaces
{
    public interface IMicrosoftProjectResolver
    {
        MsBuildSolutionPrimitive TryParseSolution(string localPath, string originalSolutionName);
        MsBuildProjectPrimitive TryParseProject(string localPath, string originalProjectName);
        MsBuildPublishProfilePrimitive TryParsePublishProfile(string localPath, string originalPublishProfileName);
    }
}