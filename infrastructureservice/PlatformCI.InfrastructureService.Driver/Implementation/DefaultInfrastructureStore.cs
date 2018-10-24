using System;
using System.Collections.Generic;
using System.Linq;
using Microsoft.EntityFrameworkCore;
using PlatformCI.InfrastructureService.Driver.Interfaces;
using PlatformCI.InfrastructureService.Entities;
using PlatformCI.InfrastructureService.Entities.EfContextBuilder;
using PlatformCI.InfrastructureService.Models.Implementation;

namespace PlatformCI.InfrastructureService.Driver.Implementation
{
    public class DefaultInfrastructureStore : IInfrastructureStore
    {
        public DefaultInfrastructureStore(EfContextBuilder buildSystemContextBuilder)
        {
            BuildSystemContextBuilder = buildSystemContextBuilder;
        }

        public EfContextBuilder BuildSystemContextBuilder { get; }

        public InfrastructureResult GetInfrastructureMetadata(InfrastructureRequestFilter requestFilter)
        {
            using (var context = BuildSystemContextBuilder.GetNewContext())
            {
                var optionalEnvironmentId = requestFilter != null
                    ? context.Environments.FirstOrDefault(x =>
                        x.EnvironmentName == requestFilter.OptionalEnvironmentName)?.EnvironmentId
                    : null;

                var environmentDescriptorList = GetEnvironmentMap(optionalEnvironmentId);

                var descriptorList = context.Deliverables
                    .Where(x => (requestFilter.RepositoryName == null
                                 || x.RepositoryName == requestFilter.RepositoryName)
                                && (requestFilter.SolutionName == null
                                    || x.SolutionName == requestFilter.SolutionName)
                                && (requestFilter.ProjectName == null
                                    || x.ProjectName == requestFilter.ProjectName))
                    .ToList();

                var results = new List<InfrastructureDescriptor>();
                foreach (var descriptor in descriptorList)
                {
                    var iisSites = GetIisSites(descriptor.DeliverableId, optionalEnvironmentId, context);
                    var iisApplications = GetIisApplications(descriptor.DeliverableId, optionalEnvironmentId, context);
                    var iisAppPools = GetIisApplicationPools(descriptor.DeliverableId, optionalEnvironmentId, context);
                    var scheduledTasks = GetScheduledTasks(descriptor.DeliverableId, optionalEnvironmentId, context);
                    var windowsServices = GetWindowsServices(descriptor.DeliverableId, optionalEnvironmentId, context);

                    var environments = ListUtil.ConcatLists(
                            iisSites.SelectMany(y => y.Environments),
                            iisApplications.SelectMany(y => y.Sites).SelectMany(y => y.Environments),
                            scheduledTasks.SelectMany(y => y.Environments),
                            windowsServices.SelectMany(y => y.Environments))
                        .Distinct().ToList();

                    var result = new InfrastructureDescriptor
                    {
                        RepositoryName = descriptor.RepositoryName,
                        SolutionName = descriptor.SolutionName,
                        ProjectName = descriptor.ProjectName,
                        IisApplicationPools = iisAppPools,
                        IisSites = iisSites,
                        IisApplications = iisApplications,
                        ScheduledTasks = scheduledTasks,
                        WindowsServices = windowsServices,
                        ApplicableEnvironments = environments
                    };

                    results.Add(result);
                }

                return new InfrastructureResult
                {
                    Infrastructure = results,
                    Environments = environmentDescriptorList
                };
            }
        }

        private IList<EnvironmentDescriptor> GetEnvironmentMap(Guid? environmentId)
        {
            using (var context = BuildSystemContextBuilder.GetNewContext())
            {
                var environments = context.Environments
                    .Where(x => environmentId == null
                                || x.EnvironmentId == environmentId)
                    .ToList();
                var result = new List<EnvironmentDescriptor>();

                foreach (var environment in environments)
                {
                    var query = context.Environments
                        .Include(a => a.ServerGroupEnvironmentAssociations)
                        .ThenInclude(a => a.ServerGroup)
                        .ThenInclude(a => a.ServerGroupAssociations)
                        .ThenInclude(a => a.Server)
                        .ThenInclude(a => a.ServerType)
                        .Where(x => x.EnvironmentId == environment.EnvironmentId)
                        .SelectMany(x => x.ServerGroupEnvironmentAssociations)
                        .Select(x => x.ServerGroup)
                        .SelectMany(x => x.ServerGroupAssociations)
                        .Select(x => x.Server)
                        .Select(x => new ServerNameType
                        {
                            ServerName = x.ServerName,
                            ServerType = x.ServerType.ServerTypeName
                        })
                        .ToList();

                    result.Add(new EnvironmentDescriptor
                    {
                        ServerNames = query,
                        Environment = environment.EnvironmentName
                    });
                }

                return result.OrderBy(x => x.Environment).ToList();
            }
        }

        private static IList<WindowsService> GetWindowsServices(
            Guid deliverableId,
            Guid? environmentId,
            BuildSystemContext parentContext)
        {
            var resourceAssociations = parentContext.ResourceAssociations
                .Where(x => x.DeliverableId == deliverableId);
            var services = resourceAssociations
                .Select(x => x.WindowsService)
                .Where(x => environmentId == null ||
                            x.WindowsServiceGroupAssociations
                                .Select(y => y.WindowsServiceGroup)
                                .Select(z => z.EnvironmentId)
                                .Contains(environmentId.Value));
            var results = services
                .Where(x => x != null)
                .Select(x => new WindowsService
                {
                    ServiceGuid = x.WindowsServiceId,
                    ServiceName = x.ServiceName,
                    BinaryPath = x.BinaryPath,
                    BinaryExecutableName = x.BinaryExecutableName,
                    BinaryExecutableArguments = x.BinaryExecutableArguments,
                    Environments = x.WindowsServiceGroupAssociations
                        .Select(y => y.WindowsServiceGroup)
                        .Select(z => z.Environment)
                        .Select(a => a.EnvironmentName)
                        .ToList()
                }).ToList();
            return results;
        }

        private static IQueryable<ResourceAssociations> GetResourceAssociations(
            Guid deliverableId,
            BuildSystemContext parentContext)
        {
            return parentContext.ResourceAssociations
                .Where(x => x.DeliverableId == deliverableId);
        }

        private static IList<ScheduledTask> GetScheduledTasks(
            Guid deliverableId,
            Guid? environmentId,
            BuildSystemContext parentContext)
        {
            var resourceAssociations = GetResourceAssociations(deliverableId, parentContext);
            var scheduledTasks = resourceAssociations
                .Where(x => x.ScheduledTaskId != null)
                .Select(x => x.ScheduledTask)
                .Where(x => environmentId == null ||
                            x.ScheduledTaskGroupAssociations
                                .Select(y => y.ScheduledTaskGroup)
                                .Select(z => z.EnvironmentId)
                                .Contains(environmentId.Value));

            var results = scheduledTasks
                .Where(x => x != null)
                .Select(x => new ScheduledTask
                {
                    TaskGuid = x.ScheduledTaskId,
                    TaskName = x.TaskName,
                    BinaryPath = x.BinaryPath,
                    BinaryExecutableName = x.BinaryExecutableName,
                    BinaryExecutableArguments = x.BinaryExecutableArguments,
                    ScheduleType = x.ScheduleType,
                    RepeatInterval = TimeSpan.FromTicks(x.RepeatInterval ?? 0),
                    RepetitionDuration = TimeSpan.FromTicks(x.RepetitionDuration ?? 0),
                    ExecutionTimeLimit = TimeSpan.FromTicks(x.ExecutionTimeLimit ?? 0),
                    Priority = x.Priority,
                    Environments = x.ScheduledTaskGroupAssociations
                        .Select(y => y.ScheduledTaskGroup)
                        .Select(z => z.Environment)
                        .Select(a => a.EnvironmentName)
                        .ToList()
                }).ToList();
            return results;
        }

        private static IList<IisApplication> GetIisApplications(
            Guid deliverableId,
            Guid? environmentId,
            BuildSystemContext parentContext)
        {
            // TODO: According to stupid ass EF Core, this node is wasn't reducible
            // probably due to nesting. Revise to combine query when EF Core gets more mature.
            var resourceAssociations = GetResourceAssociations(deliverableId, parentContext);
            var applications = resourceAssociations
                .Where(x => x.IisApplicationId != null)
                .Select(x => x.IisApplication)
                .Where(x => environmentId == null ||
                            x.IisApplicationGroupAssociations
                                .Select(y => y.IisApplicationGroup)
                                .SelectMany(z => z.IisSiteApplicationGroupAssociations)
                                .Select(a => a.IisSite)
                                .SelectMany(b => b.IisSiteGroupAssociations)
                                .Select(c => c.IisSiteGroup)
                                .Select(d => d.EnvironmentId)
                                .Contains(environmentId.Value));

            var results = applications
                .Where(x => x != null)
                .Select(x => new IisApplication
                {
                    ApplicationGuid = x.IisApplicationId,
                    ApplicationName = x.ApplicationName,
                    AppPool = new IisApplicationPool
                    {
                        AppPoolGuid = x.IisApplicationPoolId,
                        AppPoolName = x.IisApplicationPool.PoolName,
                        AppPoolFrameworkVersion = x.IisApplicationPool.PoolFrameworkVersion,
                        AppPoolType = x.IisApplicationPool.PoolType
                    },
                    PhysicalPath = x.PhysicalPath,
                    Sites = null
                }).ToList();

            foreach (var result in results)
            {
                // TODO: Why do I have to include all tables or get NullReferenceException?
                var appContext = parentContext.IisApplications
                    .Include(a => a.IisApplicationGroupAssociations)
                    .ThenInclude(a => a.IisApplicationGroup)
                    .ThenInclude(a => a.IisSiteApplicationGroupAssociations)
                    .ThenInclude(a => a.IisSite)
                    .ThenInclude(a => a.IisSiteGroupAssociations)
                    .ThenInclude(a => a.IisSiteGroup)
                    .ThenInclude(a => a.Environment)
                    .Include(a => a.IisApplicationGroupAssociations)
                    .ThenInclude(a => a.IisApplicationGroup)
                    .ThenInclude(a => a.IisSiteApplicationGroupAssociations)
                    .ThenInclude(a => a.IisSite).ThenInclude(a => a.IisSiteApplicationPool)
                    .Include(a => a.IisApplicationPool)
                    .First(x => x.IisApplicationId == result.ApplicationGuid);
                var sites = appContext
                    .IisApplicationGroupAssociations
                    .Select(y => y.IisApplicationGroup)
                    .SelectMany(z => z.IisSiteApplicationGroupAssociations)
                    .Select(a => a.IisSite)
                    .Where(y => environmentId == null ||
                                y.IisSiteGroupAssociations
                                    .Select(z => z.IisSiteGroup)
                                    .Select(d => d.EnvironmentId)
                                    .Contains(environmentId.Value))
                    .Select(b => new IisSite
                    {
                        SiteGuid = b.IisSiteId,
                        SiteName = b.SiteName,
                        AppPool = new IisApplicationPool
                        {
                            AppPoolGuid = b.IisSiteApplicationPoolId,
                            AppPoolName = b.IisSiteApplicationPool.PoolName,
                            AppPoolFrameworkVersion = b.IisSiteApplicationPool.PoolFrameworkVersion,
                            AppPoolType = b.IisSiteApplicationPool.PoolType
                        },
                        PhysicalPath = b.PhysicalPath,
                        Environments = b.IisSiteGroupAssociations
                            .Select(c => c.IisSiteGroup)
                            .Select(d => d.Environment)
                            .Select(e => e.EnvironmentName)
                            .ToList()
                    }).ToList();
                result.Sites = sites;
            }

            return results;
        }

        private static IList<IisSite> GetIisSites(
            Guid deliverableId,
            Guid? environmentId,
            BuildSystemContext parentContext)
        {
            var resourceAssociations = GetResourceAssociations(deliverableId, parentContext);
            var sites = resourceAssociations
                .Where(x => x.IisSiteId != null)
                .Select(x => x.IisSite)
                .Where(x => environmentId == null ||
                            x.IisSiteGroupAssociations
                                .Select(y => y.IisSiteGroup)
                                .Select(d => d.EnvironmentId)
                                .Contains(environmentId.Value));

            var results = sites
                .Where(x => x != null)
                .Select(x => new IisSite
                {
                    SiteGuid = x.IisSiteId,
                    SiteName = x.SiteName,
                    AppPool = new IisApplicationPool
                    {
                        AppPoolGuid = x.IisSiteApplicationPoolId,
                        AppPoolName = x.IisSiteApplicationPool.PoolName,
                        AppPoolFrameworkVersion = x.IisSiteApplicationPool.PoolFrameworkVersion,
                        AppPoolType = x.IisSiteApplicationPool.PoolType
                    },
                    PhysicalPath = x.PhysicalPath,
                    Environments = x.IisSiteGroupAssociations
                        .Select(c => c.IisSiteGroup)
                        .Select(d => d.Environment)
                        .Select(e => e.EnvironmentName)
                        .ToList()
                }).ToList();
            return results;
        }

        private static IList<IisApplicationPool> GetIisApplicationPools(
            Guid deliverableId,
            Guid? environmentId,
            BuildSystemContext parentContext)
        {
            var appPools = new List<IisApplicationPool>();
            // TODO: performance issues here, double the querying double the database fun
            var sites = GetIisSites(deliverableId, environmentId, parentContext);
            var applications = GetIisApplications(deliverableId, environmentId, parentContext);
            appPools = appPools.Concat(sites.Select(x => x.AppPool)).ToList();
            appPools = appPools.Concat(applications.Select(x => x.AppPool)).ToList();
            appPools = appPools.Concat(applications.SelectMany(x => x.Sites).Select(y => y.AppPool)).ToList();
            // TODO: Add distinct comparable class to ensure Distinct() returns the right data?
            return appPools.Distinct().ToList();
        }
    }
}