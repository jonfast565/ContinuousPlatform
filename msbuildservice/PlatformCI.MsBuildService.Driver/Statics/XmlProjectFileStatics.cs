using System;
using System.Collections.Generic;
using System.Linq;
using System.Xml;
using BuildSystem.Lib.Models.Deliverable.Implementation;

namespace BuildSystem.Lib.MicrosoftBuildProvider.Statics
{
    internal static class XmlProjectFileStatics
    {
        internal static MsBuildProjectPrimitive GetProjectFromXmlFile(
            string localPath,
            string originalProjectName)
        {
            try
            {
                var document = new XmlDocument();
                document.Load(localPath);
                var namespaceValue = document.CreateLegacyMsBuildNamespace();
                string defaultNamespace = null, assemblyName;
                List<string> targetFrameworks = null, relativePaths;
                bool isNetCoreProject;

                if (namespaceValue.NsManager != null)
                {
                    // legacy .NET Framework
                    defaultNamespace = document.GetSingleNodeInnerText("//ns:Project/ns:PropertyGroup/ns:RootNamespace",
                        namespaceValue);
                    assemblyName = document.GetSingleNodeInnerText("//ns:Project/ns:PropertyGroup/ns:AssemblyName",
                        namespaceValue);
                    relativePaths = document.GetNodesPropertyValue("//ns:Project/ns:ItemGroup/ns:ProjectReference",
                        "Include", namespaceValue)?.ToList();
                    isNetCoreProject = false;
                }
                else
                {
                    // .NET Core and newer
                    assemblyName =
                        document.GetSingleNodeInnerText("//Project/PropertyGroup/AssemblyName", namespaceValue);
                    targetFrameworks = document
                        .GetSingleNodeInnerText("//Project/PropertyGroup/TargetFramework", namespaceValue)?.Split(';')
                        ?.ToList();
                    relativePaths = document
                        .GetNodesPropertyValue("//Project/ItemGroup/ProjectReference", "Include", namespaceValue)
                        ?.ToList();
                    isNetCoreProject = true;
                }

                document = null;

                var parser = new PathParser.Implementation.PathParser();
                var projectName = parser.GetLastItemFromPath(localPath);
                var originalProjectNameExpr = originalProjectName != null
                    ? "<-" + originalProjectName
                    : string.Empty;
                // opLog.Log(LogOperationType.Info,
                    // $"Found project {projectName} {originalProjectNameExpr}");

                return new MsBuildProjectPrimitive
                {
                    Name = projectName,
                    TargetFrameworks = targetFrameworks,
                    DefaultNamespace = defaultNamespace,
                    AssemblyName = assemblyName,
                    RelativeProjectReferencePaths = relativePaths,
                    IsNetCoreProject = isNetCoreProject
                };
            }
            catch (Exception e)
            {
                return new MsBuildProjectPrimitive
                {
                    Failed = true,
                    Error = e
                };
            }
        }

        internal static MsBuildPublishProfilePrimitive GetPublishProfileFromXmlFile(
            string localPath,
            string originalPublishProfileName)
        {
            try
            {
                var document = new XmlDocument();
                document.Load(localPath);
                var namespaceValue = document.CreateLegacyMsBuildNamespace();
                var publishUrl =
                    document.GetSingleNodeInnerText("//ns:Project/ns:PropertyGroup/ns:publishUrl", namespaceValue);
                publishUrl = publishUrl.Replace("$(ProjectDir)", string.Empty);

                document = null;

                var parser = new PathParser.Implementation.PathParser();
                var publishProfileName = parser.GetLastItemFromPath(localPath);
                var originalPublishProfileNameExpr = originalPublishProfileName != null
                    ? "<-" + originalPublishProfileName
                    : string.Empty;
                // opLog.Log(LogOperationType.Info,
                    // $"Found publish profile {publishProfileName} {originalPublishProfileNameExpr}");

                return new MsBuildPublishProfilePrimitive
                {
                    Name = publishProfileName,
                    PublishUrl = publishUrl
                };
            }
            catch (Exception e)
            {
                return new MsBuildPublishProfilePrimitive
                {
                    Failed = true,
                    Error = e
                };
            }
        }
    }
}