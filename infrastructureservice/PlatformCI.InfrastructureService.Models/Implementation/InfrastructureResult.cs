using System;
using System.Collections.Generic;
using System.Linq;
using PlatformCI.InfrastructureService.Models.Comparers;

namespace PlatformCI.InfrastructureService.Models.Implementation
{
    [Serializable]
    public class InfrastructureResult
    {
        public IList<InfrastructureDescriptor> Infrastructure { get; set; }
        public IList<EnvironmentDescriptor> Environments { get; set; }

        public IDictionary<string, EnvironmentInfrastructure> GetEnvironmentData()
        {
            var dict = new Dictionary<string, EnvironmentInfrastructure>();
            foreach (var environment in Environments)
            {
                var environmentName = environment.Environment;
                var iisSites = GetIisSites(environmentName);
                var iisApplications = GetIisApplications(environmentName);
                var memberPool = IisSiteApplicationMemberPool.InitPools(iisSites, iisApplications);
                var flattenedObj = new EnvironmentInfrastructure
                {
                    Environment = environmentName,
                    Servers = GetServerNames(environmentName),
                    IisApplicationPools = GetApplicationPools(environmentName),
                    SiteApplicationMembers = memberPool,
                    ScheduledTasks = GetScheduledTasks(environmentName),
                    WindowsServices = GetWindowsServices(environmentName)
                };
                dict.Add(environmentName, flattenedObj);
            }

            return dict;
        }

        private IList<WindowsService> GetWindowsServices(string environmentName)
        {
            var windowsServices = Infrastructure
                .SelectMany(x => x.WindowsServices)
                .Where(x => x.Environments.Contains(environmentName))
                .ToList();
            return windowsServices;
        }

        private IList<ScheduledTask> GetScheduledTasks(string environmentName)
        {
            var scheduledTasks = Infrastructure
                .SelectMany(x => x.ScheduledTasks)
                .Where(x => x.Environments.Contains(environmentName))
                .ToList();
            return scheduledTasks;
        }

        private IList<ServerNameType> GetServerNames(string environmentName)
        {
            return Environments
                .Where(x => x.Environment == environmentName)
                .SelectMany(x => x.ServerNames)
                .ToList();
        }

        private IList<IisApplication> GetIisApplications(string environmentName)
        {
            return Infrastructure
                .SelectMany(x => x.IisApplications)
                .Where(x => x.Sites.SelectMany(y => y.Environments)
                    .Contains(environmentName))
                .ToList();
        }

        private IList<IisSite> GetIisSites(string environmentName)
        {
            var iisSitesStandalone = Infrastructure
                .SelectMany(x => x.IisSites)
                .Where(x => x.Environments.Contains(environmentName))
                .ToList();
            var iisApplicationSites = Infrastructure
                .SelectMany(x => x.IisApplications)
                .SelectMany(x => x.Sites)
                .Where(x => x.Environments.Contains(environmentName));
            return iisSitesStandalone.Concat(iisApplicationSites)
                .Distinct(new IisSiteEqualityComparer())
                .ToList();
        }

        private IList<IisApplicationPool> GetApplicationPools(string environmentName)
        {
            var sitePools = Infrastructure
                .SelectMany(x => x.IisSites)
                .Where(x => x.Environments.Contains(environmentName))
                .Select(x => x.AppPool)
                .ToList();
            var appPools = Infrastructure
                .SelectMany(x => x.IisApplications)
                .Where(x => x.Sites.SelectMany(y => y.Environments)
                    .Contains(environmentName))
                .Select(x => x.AppPool);
            var applicationSiteAppPools = Infrastructure
                .SelectMany(x => x.IisApplications)
                .Where(x => x.Sites.SelectMany(y => y.Environments)
                    .Contains(environmentName))
                .SelectMany(x => x.Sites)
                .Where(x => x.Environments.Contains(environmentName))
                .Select(x => x.AppPool);
            var result = sitePools
                .Concat(appPools)
                .Concat(applicationSiteAppPools)
                .Distinct(new IisApplicationPoolEqualityComparer())
                .ToList();
            return result;
        }

        public IList<FlattenedBuildInfrastructure> GetFlattenedData(
            InfrastructureRequestFilter projectKey)
        {
            var applicableDeliverables = GetApplicableDeliverables(projectKey);
            var result = new List<FlattenedBuildInfrastructure>();
            foreach (var environmentName in applicableDeliverables
                .SelectMany(x => x.ApplicableEnvironments))
            foreach (var serverName in Environments
                .Where(x => x.Environment == environmentName)
                .SelectMany(x => x.ServerNames)
                // TODO: Do not allow server type name to be a magic string
                .Where(x => x.ServerType == "Web" && HasWebTargets(projectKey)
                            || x.ServerType == "Application" && HasAppTargets(projectKey)))
            {
                var flattenedObj = new FlattenedBuildInfrastructure
                {
                    ServerName = serverName.ServerName,
                    ServerGroup = environmentName,
                    DeploymentLocations = GetApplicableDeploymentLocations(
                        projectKey,
                        environmentName),
                    AppPoolNames = GetApplicationPools(projectKey, environmentName),
                    ServiceNames = applicableDeliverables
                        .SelectMany(x => x.WindowsServices)
                        .Where(x => x.Environments.Contains(environmentName))
                        .Distinct(new WindowsServiceEqualityComparer())
                        .Select(x => x.ServiceName)
                        .ToList(),
                    TaskNames = applicableDeliverables
                        .SelectMany(x => x.ScheduledTasks)
                        .Where(x => x.Environments.Contains(environmentName))
                        .Distinct(new ScheduledTaskEqualityComparer())
                        .Select(x => x.TaskName)
                        .ToList()
                };
                result.Add(flattenedObj);
            }

            return result;
        }

        private IList<string> GetApplicationPools(
            InfrastructureRequestFilter projectKey,
            string environmentName)
        {
            var deliverables = GetApplicableDeliverables(projectKey);
            if (deliverables
                .SelectMany(x => x.IisSites)
                .Any())
                return deliverables
                    .SelectMany(x => x.IisSites)
                    .Where(x => x.Environments.Contains(environmentName))
                    .Select(x => x.AppPool)
                    .Select(x => x.AppPoolName)
                    .ToList();
            // ReSharper disable once ConvertIfStatementToReturnStatement
            if (deliverables
                .SelectMany(x => x.IisApplications)
                .Any())
                return deliverables
                    .SelectMany(x => x.IisApplications)
                    .Where(x => x.Sites
                        .SelectMany(y => y.Environments)
                        .Contains(environmentName))
                    .Select(x => x.AppPool)
                    .Select(x => x.AppPoolName)
                    .ToList();
            return new List<string>();
        }

        private IList<string> GetApplicableDeploymentLocations(
            InfrastructureRequestFilter projectKey,
            string environmentName)
        {
            var deliverables = GetApplicableDeliverables(projectKey);
            if (deliverables
                .SelectMany(x => x.IisSites)
                .Any())
            {
                var result = deliverables
                    .SelectMany(x => x.IisSites)
                    .Where(x => x.Environments.Contains(environmentName))
                    .Select(x => x.PhysicalPath)
                    .ToList();
                return result;
            }

            if (deliverables
                .SelectMany(x => x.IisApplications)
                .Any())
            {
                var iisSiteForEnvironment = deliverables
                    .SelectMany(x => x.IisApplications)
                    .SelectMany(x => x.Sites)
                    .First(x => x.Environments.Contains(environmentName))
                    .SiteName;
                var result = deliverables
                    .SelectMany(x => x.IisApplications)
                    .Where(x => x.Sites
                        .SelectMany(y => y.Environments)
                        .Contains(environmentName))
                    .Select(x => x.PhysicalPath.Replace("{{SiteName}}", iisSiteForEnvironment))
                    .ToList();
                return result;
            }

            if (deliverables
                .SelectMany(x => x.ScheduledTasks)
                .Any())
            {
                var result = deliverables
                    .SelectMany(x => x.ScheduledTasks)
                    .Select(x => x.BinaryPath)
                    .ToList();
                return result;
            }

            if (deliverables
                .SelectMany(x => x.WindowsServices)
                .Any())
            {
                var result = deliverables
                    .SelectMany(x => x.WindowsServices)
                    .Select(x => x.BinaryPath)
                    .ToList();
                return result;
            }

            throw new Exception(
                $"Project key {projectKey.ProjectName} " + //TODO: fill this in
                "has no associated infrastructure data, cannot continue");
        }

        private IQueryable<InfrastructureDescriptor> GetApplicableDeliverables(
            InfrastructureRequestFilter projectKey)
        {
            return Infrastructure
                .Where(x => x.SolutionName == projectKey.SolutionName
                            && x.ProjectName == projectKey.ProjectName
                            && x.RepositoryName == projectKey.RepositoryName)
                .AsQueryable();
        }

        private bool HasWebTargets(InfrastructureRequestFilter projectKey)
        {
            return GetApplicableDeliverables(projectKey)
                       .SelectMany(x => x.IisSites)
                       .Any()
                   || GetApplicableDeliverables(projectKey)
                       .SelectMany(x => x.IisApplications)
                       .Any()
                   || GetApplicableDeliverables(projectKey)
                       .SelectMany(x => x.IisApplicationPools)
                       .Any();
        }

        private bool HasAppTargets(InfrastructureRequestFilter projectKey)
        {
            return GetApplicableDeliverables(projectKey)
                       .SelectMany(x => x.WindowsServices)
                       .Any()
                   || GetApplicableDeliverables(projectKey)
                       .SelectMany(x => x.ScheduledTasks)
                       .Any();
        }
    }
}