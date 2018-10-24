using BuildSystem.Lib.MicrosoftBuildProvider.Interfaces;
using BuildSystem.Lib.MicrosoftBuildProvider.Statics;
using BuildSystem.Lib.Models.Deliverable.Implementation;
using BuildSystem.Lib.Oplog.Interfaces;

namespace BuildSystem.Lib.MicrosoftBuildProvider.Implementation
{
    internal class DefaultMicrosoftProjectResolver : IMicrosoftProjectResolver
    {
        public DefaultMicrosoftProjectResolver(IOplog opLog)
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
            return MsBuildProjectStatics
                .GetProjectWithMsBuild(localPath, originalProjectName, _opLog);
        }

        public MsBuildPublishProfilePrimitive TryParsePublishProfile(string localPath,
            string originalPublishProfileName)
        {
            return MsBuildProjectStatics
                .GetPublicProfileWithMsBuild(localPath, originalPublishProfileName, _opLog);
        }
    }
}