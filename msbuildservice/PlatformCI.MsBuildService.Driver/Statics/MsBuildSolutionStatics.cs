using System;
using System.IO;
using Microsoft.Build.Construction;
using Microsoft.Build.Evaluation;
using Microsoft.Build.Exceptions;
using PlatformCI.MsBuildService.Models.Implementation;

namespace PlatformCI.MsBuildService.Driver.Statics
{
    internal static class MsBuildSolutionStatics
    {
        internal static MsBuildSolutionPrimitive GetSolutionWithMsBuild(
            string localPath,
            string originalSolutionName)
        {
            try
            {
                ProjectCollection.GlobalProjectCollection.UnloadAllProjects();

                var solution = SolutionFile.Parse(localPath);
                var parser = new PathParser();
                var solutionName = parser.GetLastItemFromPath(localPath);
                var configurations = MsBuildPropertyStatics.GetConfigurations(solution);
                var relativePaths = MsBuildPropertyStatics.GetProjectRelativePaths(solution);
                var originalSolutionNameExpr = originalSolutionName != null
                    ? "<-" + originalSolutionName
                    : string.Empty;

                Console.WriteLine($"Found solution {solutionName} {originalSolutionNameExpr}");

                return new MsBuildSolutionPrimitive
                {
                    Name = solutionName,
                    Configurations = configurations,
                    RelativeProjectPaths = relativePaths
                };
            }
            catch (Exception e) when
            (e is InvalidProjectFileException
             || e is InvalidOperationException
             || e is IOException)
            {
                return new MsBuildSolutionPrimitive
                {
                    Error = e,
                    Failed = true
                };
            }
        }
    }
}