using BuildSystem.Lib.MicrosoftBuildProvider.Interfaces;
using BuildSystem.Lib.MicrosoftBuildProvider.Statics;
using BuildSystem.Lib.Models.Deliverable.Implementation;
using BuildSystem.Lib.Oplog.Interfaces;

namespace BuildSystem.Lib.MicrosoftBuildProvider.Implementation
{
    internal class HackedMicrosoftProjectResolver : IMicrosoftProjectResolver
    {
        public HackedMicrosoftProjectResolver(IOplog opLog)
        {
            _opLog = opLog;
        }

        public IOplog _opLog { get; }

        public MsBuildSolutionPrimitive TryParseSolution(string localPath, string originalSolutionName)
        {
            return MsBuildSolutionStatics
                .GetSolutionWithMsBuild(localPath, originalSolutionName, _opLog);
        }

        public MsBuildProjectPrimitive TryParseProject(string localPath, string originalProjectName)
        {
            return XmlProjectFileStatics
                .GetProjectFromXmlFile(localPath, originalProjectName, _opLog);
        }

        public MsBuildPublishProfilePrimitive TryParsePublishProfile(string localPath,
            string originalPublishProfileName)
        {
            return XmlProjectFileStatics
                .GetPublishProfileFromXmlFile(localPath, originalPublishProfileName, _opLog);
        }
    }
}