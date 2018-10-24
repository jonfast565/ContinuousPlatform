using System;
using System.IO;
using BuildSystem.Lib.Models.Deliverable.Implementation;
using BuildSystem.Lib.Oplog.Enums;
using BuildSystem.Lib.Oplog.Interfaces;
using Microsoft.Build.Construction;
using Microsoft.Build.Evaluation;
using Microsoft.Build.Exceptions;

namespace BuildSystem.Lib.MicrosoftBuildProvider.Statics
{
    internal static class MsBuildSolutionStatics
    {
        internal static MsBuildSolutionPrimitive GetSolutionWithMsBuild(
            string localPath,
            string originalSolutionName,
            IOplog opLog)
        {
            try
            {
                ProjectCollection.GlobalProjectCollection.UnloadAllProjects();

                var solution = SolutionFile.Parse(localPath);
                var parser = new PathParser.Implementation.PathParser();
                var solutionName = parser.GetLastItemFromPath(localPath);
                var configurations = MsBuildPropertyStatics.GetConfigurations(solution);
                var relativePaths = MsBuildPropertyStatics.GetProjectRelativePaths(solution);
                var originalSolutionNameExpr = originalSolutionName != null
                    ? "<-" + originalSolutionName
                    : string.Empty;

                opLog.Log(LogOperationType.Info,
                    $"Found solution {solutionName} {originalSolutionNameExpr}");

                return new MsBuildSolutionPrimitive
                {
                    Name = solutionName,
                    Configurations = configurations,
                    ProjectRelativePaths = relativePaths
                };
            }
            catch (Exception e) when
            (e is InvalidProjectFileException
             || e is InvalidOperationException
             || e is IOException)
            {
                opLog.Log(LogOperationType.Error, e);
                return new MsBuildSolutionPrimitive
                {
                    Error = e,
                    Failed = true
                };
            }
        }
    }
}