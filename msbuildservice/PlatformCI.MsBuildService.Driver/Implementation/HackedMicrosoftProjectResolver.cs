using PlatformCI.MsBuildService.Driver.Interfaces;
using PlatformCI.MsBuildService.Driver.Statics;
using PlatformCI.MsBuildService.Models.Implementation;

namespace PlatformCI.MsBuildService.Driver.Implementation
{
    public class HackedMicrosoftProjectResolver : IMicrosoftProjectResolver
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