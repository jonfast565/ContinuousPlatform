using System.Collections.Generic;
using System.Linq;
using Microsoft.Build.Construction;
using Microsoft.Build.Evaluation;

namespace PlatformCI.MsBuildService.Driver.Statics
{
    internal static class MsBuildPropertyStatics
    {
        internal static string GetPublishUrl(Project publishProfile)
        {
            return publishProfile.Properties
                .FirstOrDefault(x => x.Name == "publishUrl")
                ?.EvaluatedValue;
        }

        internal static ICollection<string> GetConfigurations(SolutionFile solution)
        {
            return solution.SolutionConfigurations
                .Select(x => x.ConfigurationName)
                .Distinct()
                .ToList();
        }

        internal static ICollection<string> GetProjectRelativePaths(SolutionFile solution)
        {
            return solution.ProjectsInOrder
                .Select(x => x.RelativePath)
                .ToList();
        }

        internal static IEnumerable<string> GetRelativeProjectReferencePaths(Project project)
        {
            return project.Items
                .Where(x => x.ItemType == "ProjectReference")
                .Select(x => x.EvaluatedInclude)
                .ToList();
        }

        internal static string GetDefaultNamespace(Project project)
        {
            return project.Properties
                .Where(x => x.Name == "RootNamespace")
                .Select(x => x.EvaluatedValue)
                .FirstOrDefault();
        }

        internal static string GetAssemblyName(Project project)
        {
            return project.Properties
                .Where(x => x.Name == "AssemblyName")
                .Select(x => x.EvaluatedValue)
                .FirstOrDefault();
        }

        internal static IEnumerable<string> GetTargetFrameworksIfApplicable(Project project)
        {
            return project.Properties
                .Where(x => x.Name == "TargetFramework")
                .Select(x => x.EvaluatedValue)
                .FirstOrDefault()?.Split(';');
        }
    }
}