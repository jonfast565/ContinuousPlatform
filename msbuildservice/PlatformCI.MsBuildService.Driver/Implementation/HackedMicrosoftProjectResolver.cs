using BuildSystem.Lib.MicrosoftBuildProvider.Interfaces;
using BuildSystem.Lib.MicrosoftBuildProvider.Statics;
using BuildSystem.Lib.Models.Deliverable.Implementation;

namespace BuildSystem.Lib.MicrosoftBuildProvider.Implementation
{
    internal class HackedMicrosoftProjectResolver : IMicrosoftProjectResolver
    {
        public MsBuildSolutionPrimitive TryParseSolution(string localPath, string originalSolutionName)
        {
            return MsBuildSolutionStatics
                .GetSolutionWithMsBuild(localPath, originalSolutionName);
        }

        public MsBuildProjectPrimitive TryParseProject(string localPath, string originalProjectName)
        {
            return XmlProjectFileStatics
                .GetProjectFromXmlFile(localPath, originalProjectName);
        }

        public MsBuildPublishProfilePrimitive TryParsePublishProfile(string localPath,
            string originalPublishProfileName)
        {
            return XmlProjectFileStatics
                .GetPublishProfileFromXmlFile(localPath, originalPublishProfileName);
        }
    }
}