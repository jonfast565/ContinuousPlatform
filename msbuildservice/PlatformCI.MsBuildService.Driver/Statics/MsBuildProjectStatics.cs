using System;
using System.IO;
using System.Linq;
using BuildSystem.Lib.Models.Deliverable.Implementation;
using BuildSystem.Lib.Oplog.Enums;
using BuildSystem.Lib.Oplog.Interfaces;
using Microsoft.Build.Evaluation;
using Microsoft.Build.Exceptions;

namespace BuildSystem.Lib.MicrosoftBuildProvider.Statics
{
    internal static class MsBuildProjectStatics
    {
        internal static MsBuildPublishProfilePrimitive GetPublicProfileWithMsBuild(
            string localPath,
            string originalPublishProfileName,
            IOplog opLog)
        {
            try
            {
                ProjectCollection.GlobalProjectCollection.UnloadAllProjects();

                var publishProfile = new Project(localPath);
                var parser = new PathParser.Implementation.PathParser();
                var publishProfileName = parser.GetLastItemFromPath(localPath);
                var publishUrl = MsBuildPropertyStatics.GetPublishUrl(publishProfile);
                var originalPublishProfileNameExpr = originalPublishProfileName != null
                    ? "<-" + originalPublishProfileName
                    : string.Empty;

                opLog.Log(LogOperationType.Info,
                    $"Found publish profile {publishProfileName} {originalPublishProfileNameExpr}");

                return new MsBuildPublishProfilePrimitive
                {
                    Name = publishProfileName,
                    PublishUrl = publishUrl
                };
            }
            catch (Exception e) when
            (e is InvalidProjectFileException
             || e is InvalidOperationException
             || e is IOException)
            {
                opLog.Log(LogOperationType.Error, e);
                return new MsBuildPublishProfilePrimitive
                {
                    Failed = true,
                    Error = e
                };
            }
        }

        internal static MsBuildProjectPrimitive GetProjectWithMsBuild(
            string localPath,
            string originalProjectName,
            IOplog opLog)
        {
            try
            {
                ProjectCollection.GlobalProjectCollection.UnloadAllProjects();

                var project = new Project(localPath);
                var parser = new PathParser.Implementation.PathParser();
                var projectName = parser.GetLastItemFromPath(localPath);
                var targetFrameworks = MsBuildPropertyStatics.GetTargetFrameworksIfApplicable(project);
                var relativePaths = MsBuildPropertyStatics.GetRelativeProjectReferencePaths(project);
                var defaultNamespace = MsBuildPropertyStatics.GetDefaultNamespace(project);
                var assemblyName = MsBuildPropertyStatics.GetAssemblyName(project);
                var originalProjectNameExpr = originalProjectName != null
                    ? "<-" + originalProjectName
                    : string.Empty;

                opLog.Log(LogOperationType.Info,
                    $"Found project {projectName} {originalProjectNameExpr}");

                return new MsBuildProjectPrimitive
                {
                    Name = projectName,
                    TargetFrameworks = targetFrameworks?.ToList(),
                    DefaultNamespace = defaultNamespace,
                    AssemblyName = assemblyName,
                    RelativeProjectReferencePaths = relativePaths.ToList()
                };
            }
            catch (Exception e) when
            (e is InvalidProjectFileException
             || e is InvalidOperationException
             || e is IOException)
            {
                opLog.Log(LogOperationType.Error, e);
                return new MsBuildProjectPrimitive
                {
                    Failed = true,
                    Error = e
                };
            }
        }
    }
}