using BuildSystem.Lib.MicrosoftBuildProvider.Interfaces;
using BuildSystem.Lib.MicrosoftBuildProvider.Statics;
using BuildSystem.Lib.Models.Deliverable.Implementation;

namespace BuildSystem.Lib.MicrosoftBuildProvider.Implementation
{
    internal class DefaultMicrosoftProjectResolver : IMicrosoftProjectResolver
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