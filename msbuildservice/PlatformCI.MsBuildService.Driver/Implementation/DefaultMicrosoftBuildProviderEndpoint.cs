using System;
using PlatformCI.MsBuildService.Driver.Interfaces;
using PlatformCI.MsBuildService.Models.Implementation;
using PlatformCI.MsBuildService.Models.Interfaces;

namespace PlatformCI.MsBuildService.Driver.Implementation
{
    public class DefaultMicrosoftBuildProviderEndpoint : IMicrosoftBuildProviderEndpoint
    {
        private readonly IFilesystemProvider _provider;
        private readonly IMicrosoftProjectResolver _resolver;

        public DefaultMicrosoftBuildProviderEndpoint(IFilesystemProvider provider, IMicrosoftProjectResolver resolver)
        {
            _provider = provider;
            _resolver = resolver;
        }

        public MsBuildProjectPrimitive GetProjectFromLocalPath(
            string localPath,
            string originalProjectName = null)
        {
            return _resolver.TryParseProject(localPath, originalProjectName);
        }

        public MsBuildPublishProfilePrimitive GetPublishProfileFromLocalPath(
            string localPath,
            string originalPublishProfileName = null)
        {
            return _resolver.TryParsePublishProfile(localPath, originalPublishProfileName);
        }

        public MsBuildSolutionPrimitive GetSolutionFromLocalPath(
            string localPath,
            string originalSolutionName = null)
        {
            return _resolver.TryParseSolution(localPath, originalSolutionName);
        }

        public MsBuildProjectPrimitive GetProjectFromFileBytes(FilePayload localPath)
        {
            Console.WriteLine($"Caching {localPath.Name}");
            var projectPath = WriteMsBuildItemToDisk(localPath);
            var project = GetProjectFromLocalPath(projectPath, localPath.Name);
            DeleteMsBuildItemFromDisk(projectPath);
            project.Name = localPath.Name;
            return project;
        }

        public MsBuildPublishProfilePrimitive GetPublishProfileFromFileBytes(FilePayload localPath)
        {
            Console.WriteLine($"Caching {localPath.Name}");
            var publishProfilePath = WriteMsBuildItemToDisk(localPath);
            var publishProfile = GetPublishProfileFromLocalPath(publishProfilePath, localPath.Name);
            DeleteMsBuildItemFromDisk(publishProfilePath);
            publishProfile.Name = localPath.Name;
            return publishProfile;
        }

        public MsBuildSolutionPrimitive GetSolutionFromFileBytes(FilePayload localPath)
        {
            Console.WriteLine($"Caching {localPath.Name}");
            var solutionPath = WriteMsBuildItemToDisk(localPath);
            var solution = GetSolutionFromLocalPath(solutionPath, localPath.Name);
            DeleteMsBuildItemFromDisk(solutionPath);
            solution.Name = localPath.Name;
            return solution;
        }

        private string WriteMsBuildItemToDisk(FilePayload payload)
        {
            var path = _provider.GetTemporarySystemPath();
            _provider.WriteFile(path, payload.Bytes);
            return path;
        }

        private void DeleteMsBuildItemFromDisk(string localPath)
        {
            _provider.DeleteFile(localPath);
        }
    }
}