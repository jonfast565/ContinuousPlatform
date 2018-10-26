using PlatformCI.MsBuildService.Models.Implementation;

namespace PlatformCI.MsBuildService.Driver.Interfaces
{
    public interface IMicrosoftProjectResolver
    {
        MsBuildSolutionPrimitive TryParseSolution(string localPath, string originalSolutionName);
        MsBuildProjectPrimitive TryParseProject(string localPath, string originalProjectName);
        MsBuildPublishProfilePrimitive TryParsePublishProfile(string localPath, string originalPublishProfileName);
    }
}