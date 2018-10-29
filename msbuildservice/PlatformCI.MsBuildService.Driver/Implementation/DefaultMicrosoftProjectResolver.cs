using PlatformCI.MsBuildService.Driver.Interfaces;
using PlatformCI.MsBuildService.Driver.Statics;
using PlatformCI.MsBuildService.Models.Implementation;

namespace PlatformCI.MsBuildService.Driver.Implementation
{
    public class DefaultMicrosoftProjectResolver : IMicrosoftProjectResolver
    {
        public MsBuildSolutionPrimitive TryParseSolution(string localPath, string originalSolutionName)
        {
            return MsBuildSolutionStatics
                .GetSolutionWithMsBuild(localPath, originalSolutionName);
        }

        public MsBuildProjectPrimitive TryParseProject(string localPath, string originalProjectName)
        {
            return MsBuildProjectStatics
                .GetProjectWithMsBuild(localPath, originalProjectName);
        }

        public MsBuildPublishProfilePrimitive TryParsePublishProfile(string localPath,
            string originalPublishProfileName)
        {
            return MsBuildProjectStatics
                .GetPublicProfileWithMsBuild(localPath, originalPublishProfileName);
        }
    }
}